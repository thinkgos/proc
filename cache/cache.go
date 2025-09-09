package cache

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	// For use with functions that take an expiration time.
	NoExpiration time.Duration = -1
	// For use with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or
	// NewFrom() when the cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = 0
)

var (
	ErrValueNotFound       = errors.New("cache: value not found")
	ErrValueNotValidNumber = errors.New("cache: value not a valid number")
	ErrValueNotValidFloat  = errors.New("cache: value not a valid float32 or float64")
)

type Item struct {
	Value      any   // cache value
	Expiration int64 // unix nanosecond
}

// Returns true if the item has expired.
func (i *Item) Expired() bool {
	return i.Expiration > 0 && time.Now().UnixNano() > i.Expiration
}

type Cache struct {
	*cache
	// If this is confusing, see the comment at the bottom of New()
}

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	mu                sync.RWMutex
	onEvicted         func(string, any)
	janitor           *janitor
}

// Set an item to the cache, replacing any existing item.
// If the duration is 0(DefaultExpiration), the cache's default expiration time is used.
// If it is -1(NoExpiration), the item never expires.
func (c *cache) Set(k string, x any, d time.Duration) {
	e := c.calcExpiration(d)
	c.mu.Lock()
	c.items[k] = Item{
		Value:      x,
		Expiration: e,
	}
	c.mu.Unlock()
}

// SetDefault set an item to the cache, replacing any existing item, using the default
// expiration.
func (c *cache) SetDefault(k string, x any) {
	c.Set(k, x, DefaultExpiration)
}

// SetNX set an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an true if set success.
func (c *cache) SetNX(k string, x any, d time.Duration) bool {
	c.mu.Lock()
	_, found := c.getValue(k)
	if found {
		c.mu.Unlock()
		return false
	}
	c.items[k] = Item{
		Value:      x,
		Expiration: c.calcExpiration(d),
	}
	c.mu.Unlock()
	return true
}

// SetXX set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise. Returns an true if set success.
func (c *cache) SetXX(k string, x any, d time.Duration) (any, bool) {
	c.mu.Lock()
	val, found := c.getValue(k)
	if !found {
		c.mu.Unlock()
		return nil, false
	}
	c.items[k] = Item{
		Value:      x,
		Expiration: c.calcExpiration(d),
	}
	c.mu.Unlock()
	return val, true
}

type UpsertCb func(exist bool, valueInMap any) any

func (c *cache) Upsert(k string, cb UpsertCb, d time.Duration) (val any) {
	var e int64

	c.mu.Lock()
	defer c.mu.Unlock()
	v, found := c.items[k]
	if !found || v.Expired() {
		val = cb(false, nil)
		e = c.calcExpiration(d)
	} else {
		val = cb(true, v.Value)
		e = v.Expiration
	}
	c.items[k] = Item{
		Value:      val,
		Expiration: e,
	}
	return val
}

func (c *cache) calcExpiration(d time.Duration) int64 {
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		return time.Now().Add(d).UnixNano()
	} else {
		return 0
	}
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (c *cache) Get(k string) (any, bool) {
	c.mu.RLock()
	val, found := c.getValue(k)
	c.mu.RUnlock()
	return val, found
}

// GetEx get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found. if key found, update with new expires.
func (c *cache) GetEx(k string, d time.Duration) (any, bool) {
	c.mu.Lock()
	item, found := c.items[k]
	if !found {
		c.mu.Unlock()
		return nil, false
	}
	if item.Expired() {
		c.mu.Unlock()
		return nil, false
	}
	item.Expiration = c.calcExpiration(d)
	c.items[k] = item
	c.mu.Unlock()
	return item.Value, true
}

// GetDel get an item from the cache the delete it from the cache. Returns the item or nil, and a bool indicating
// whether the key was found. if key found, delete it.
func (c *cache) GetDel(k string) (any, bool) {
	c.mu.Lock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.Unlock()
		return nil, false
	}
	if item.Expired() {
		c.mu.Unlock()
		return nil, false
	}
	delete(c.items, k)
	onEvicted := c.onEvicted
	c.mu.Unlock()
	if onEvicted != nil {
		onEvicted(k, item.Value)
	}
	return item.Value, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (c *cache) GetWithExpiration(k string) (any, time.Time, bool) {
	c.mu.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, time.Time{}, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, time.Time{}, false
		}
		// Return the item and the expiration time
		c.mu.RUnlock()
		return item.Value, time.Unix(0, item.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	c.mu.RUnlock()
	return item.Value, time.Time{}, true
}

func (c *cache) getValue(k string) (any, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	// "Inlining" of Expired
	if item.Expired() {
		return nil, false
	}
	return item.Value, true
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (c *cache) Delete(k string) {
	var val Item
	var found bool

	c.mu.Lock()
	onEvicted := c.onEvicted
	if onEvicted != nil {
		val, found = c.items[k]
	}
	delete(c.items, k)
	c.mu.Unlock()
	if found {
		onEvicted(k, val.Value)
	}
}

type keyValPair struct {
	key   string
	value any
}

// Delete all expired items from the cache.
func (c *cache) DeleteExpired() {
	var evictedItems []keyValPair

	now := time.Now().UnixNano()
	c.mu.Lock()
	onEvicted := c.onEvicted
	for k, v := range c.items {
		// "Inlining" of expired
		if v.Expiration > 0 && now > v.Expiration {
			delete(c.items, k)
			if onEvicted != nil {
				evictedItems = append(evictedItems, keyValPair{k, v.Value})
			}
		}
	}
	c.mu.Unlock()
	for _, v := range evictedItems {
		onEvicted(v.key, v.value)
	}
}

// Clear all items from the cache.
func (c *cache) Clear() {
	c.mu.Lock()
	onEvicted := c.onEvicted
	old := c.items
	c.items = make(map[string]Item)
	c.mu.Unlock()
	if onEvicted != nil {
		for k, v := range old {
			onEvicted(k, v.Value)
		}
	}
}

// Expire  update with new expires if key found, and return a bool indicating
// whether update expires successfully
func (c *cache) Expire(k string, d time.Duration) bool {
	c.mu.Lock()
	item, found := c.items[k]
	if !found || item.Expired() {
		c.mu.Unlock()
		return false
	}
	item.Expiration = c.calcExpiration(d)
	c.items[k] = item
	c.mu.Unlock()
	return true
}

func (c *cache) setOnEvicted(f func(string, any)) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

// Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache) Save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = errors.New("registering item types with Gob library")
		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		gob.Register(v.Value)
	}
	err = enc.Encode(&c.items)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache) SaveFile(fname string) error {
	fp, err := os.Create(fname)
	if err != nil {
		return err
	}
	err = c.Save(fp)
	if err != nil {
		_ = fp.Close()
		return err
	}
	return fp.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache) Load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := map[string]Item{}
	err := dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range items {
			ov, found := c.items[k]
			if !found || ov.Expired() {
				c.items[k] = v
			}
		}
	}
	return err
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache) LoadFile(fname string) error {
	fp, err := os.Open(fname)
	if err != nil {
		return err
	}
	err = c.Load(fp)
	if err != nil {
		_ = fp.Close()
		return err
	}
	return fp.Close()
}

// Items Copies all unexpired items in the cache into a new map and returns it.
func (c *cache) Items() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]Item, len(c.items))
	now := time.Now().UnixNano()
	for k, v := range c.items {
		// "Inlining" of Expired
		if v.Expiration > 0 && now > v.Expiration {
			continue
		}
		m[k] = v
	}
	return m
}

// Count returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (c *cache) Count() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}

type janitor struct {
	interval time.Duration
	stop     chan struct{}
}

func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- struct{}{}
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		interval: ci,
		stop:     make(chan struct{}),
	}
	c.janitor = j
	go c.janitor.Run(c)
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return NewFrom(defaultExpiration, cleanupInterval, make(map[string]Item))
}

// Return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
//
// NewFrom() also accepts an items map which will serve as the underlying map
// for the cache. This is useful for starting from a deserialized cache
// (serialized using e.g. gob.Encode() on c.Items()), or passing in e.g.
// make(map[string]Item, 500) to improve startup performance when the cache
// is expected to reach a certain minimum size.
//
// Only the cache's methods synchronize access to this map, so it is not
// recommended to keep any references to the map around after creating a cache.
// If need be, the map can be accessed at a later point using c.Items() (subject
// to the same caveat.)
//
// Note regarding serialization: When using e.g. gob, make sure to
// gob.Register() the individual types stored in the cache before encoding a
// map retrieved with c.Items(), and to register those same types before
// decoding a blob containing an items map.
func NewFrom(defaultExpiration, cleanupInterval time.Duration, items map[string]Item) *Cache {
	if defaultExpiration == 0 {
		defaultExpiration = -1
	}
	c := &cache{
		defaultExpiration: defaultExpiration,
		items:             items,
	}
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.
	C := &Cache{c}
	if cleanupInterval > 0 {
		runJanitor(c, cleanupInterval)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}

// Sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but
// not when it is overwritten.) Set to nil to disable.
func (c *Cache) OnEvicted(f func(string, any)) *Cache {
	c.setOnEvicted(f)
	return c
}
