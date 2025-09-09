package cache

import (
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Incr increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an number, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrInt64.
func (c *cache) Incr(k string, n int64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return ErrValueNotFound
	}
	switch vv := v.Value.(type) {
	case int:
		v.Value = vv + int(n)
	case int8:
		v.Value = vv + int8(n)
	case int16:
		v.Value = vv + int16(n)
	case int32:
		v.Value = vv + int32(n)
	case int64:
		v.Value = vv + n
	case uint:
		v.Value = vv + uint(n)
	case uintptr:
		v.Value = vv + uintptr(n)
	case uint8:
		v.Value = vv + uint8(n)
	case uint16:
		v.Value = vv + uint16(n)
	case uint32:
		v.Value = vv + uint32(n)
	case uint64:
		v.Value = vv + uint64(n)
	case float32:
		v.Value = vv + float32(n)
	case float64:
		v.Value = vv + float64(n)
	default:
		c.mu.Unlock()
		return ErrValueNotValidNumber
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// IncrFloat increment an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to increment it by n. Pass a negative number to decrement the
// value. To retrieve the incremented value, use one of the specialized methods.
func (c *cache) IncrFloat(k string, n float64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return ErrValueNotFound
	}
	switch vv := v.Value.(type) {
	case float32:
		v.Value = vv + float32(n)
	case float64:
		v.Value = vv + n
	default:
		c.mu.Unlock()
		return ErrValueNotValidFloat
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// IncrInt increment an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrInt(k string, n int) (int, error) {
	return incr(c, k, n)
}

// IncrInt8 increment an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrInt8(k string, n int8) (int8, error) {
	return incr(c, k, n)
}

// IncrInt16 increment an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrInt16(k string, n int16) (int16, error) {
	return incr(c, k, n)
}

// IncrInt32 increment an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrInt32(k string, n int32) (int32, error) {
	return incr(c, k, n)
}

// IncrInt64 increment an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrInt64(k string, n int64) (int64, error) {
	return incr(c, k, n)
}

// IncrUint increment an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the incremented
// value is returned.
func (c *cache) IncrUint(k string, n uint) (uint, error) {
	return incr(c, k, n)
}

// IncrUintptr increment an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrUintptr(k string, n uintptr) (uintptr, error) {
	return incr(c, k, n)
}

// IncrUint8 increment an item of type uint8 by n. Returns an error if the item's value
// is not an uint8, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrUint8(k string, n uint8) (uint8, error) {
	return incr(c, k, n)
}

// IncrUint16 increment an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrUint16(k string, n uint16) (uint16, error) {
	return incr(c, k, n)
}

// IncrUint32 increment an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrUint32(k string, n uint32) (uint32, error) {
	return incr(c, k, n)
}

// IncrUint64 increment an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrUint64(k string, n uint64) (uint64, error) {
	return incr(c, k, n)
}

// IncrFloat32 increment an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrFloat32(k string, n float32) (float32, error) {
	return incr(c, k, n)
}

// IncrFloat64 increment an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// incremented value is returned.
func (c *cache) IncrFloat64(k string, n float64) (float64, error) {
	return incr(c, k, n)
}

// Decr decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an number, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrInt64.
func (c *cache) Decr(k string, n int64) error {
	// (Cannot do Incr(k, n*-1) for uints.)
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return ErrValueNotFound
	}
	switch vv := v.Value.(type) {
	case int:
		v.Value = vv - int(n)
	case int8:
		v.Value = vv - int8(n)
	case int16:
		v.Value = vv - int16(n)
	case int32:
		v.Value = vv - int32(n)
	case int64:
		v.Value = vv - n
	case uint:
		v.Value = vv - uint(n)
	case uintptr:
		v.Value = vv - uintptr(n)
	case uint8:
		v.Value = vv - uint8(n)
	case uint16:
		v.Value = vv - uint16(n)
	case uint32:
		v.Value = vv - uint32(n)
	case uint64:
		v.Value = vv - uint64(n)
	case float32:
		v.Value = vv - float32(n)
	case float64:
		v.Value = vv - float64(n)
	default:
		c.mu.Unlock()
		return ErrValueNotValidNumber
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// DecrFloat decrement an item of type float32 or float64 by n. Returns an error if the
// item's value is not floating point, if it was not found, or if it is not
// possible to decrement it by n. Pass a negative number to decrement the
// value. To retrieve the decremented value, use one of the specialized methods,
func (c *cache) DecrFloat(k string, n float64) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return ErrValueNotFound
	}
	switch vv := v.Value.(type) {
	case float32:
		v.Value = vv - float32(n)
	case float64:
		v.Value = vv - n
	default:
		c.mu.Unlock()
		return ErrValueNotValidFloat
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// DecrInt decrement an item of type int by n. Returns an error if the item's value is
// not an int, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrInt(k string, n int) (int, error) {
	return decr(c, k, n)
}

// DecrInt8 decrement an item of type int8 by n. Returns an error if the item's value is
// not an int8, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrInt8(k string, n int8) (int8, error) {
	return decr(c, k, n)
}

// DecrInt16 decrement an item of type int16 by n. Returns an error if the item's value is
// not an int16, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrInt16(k string, n int16) (int16, error) {
	return decr(c, k, n)
}

// DecrInt32 an item of type int32 by n. Returns an error if the item's value is
// not an int32, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrInt32(k string, n int32) (int32, error) {
	return decr(c, k, n)
}

// DecrInt64 an item of type int64 by n. Returns an error if the item's value is
// not an int64, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrInt64(k string, n int64) (int64, error) {
	return decr(c, k, n)
}

// decrement an item of type uint by n. Returns an error if the item's value is
// not an uint, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrUint(k string, n uint) (uint, error) {
	return decr(c, k, n)
}

// decrement an item of type uintptr by n. Returns an error if the item's value
// is not an uintptr, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrUintptr(k string, n uintptr) (uintptr, error) {
	return decr(c, k, n)
}

// decrement an item of type uint8 by n. Returns an error if the item's value is
// not an uint8, or if it was not found. If there is no error, the decremented
// value is returned.
func (c *cache) DecrUint8(k string, n uint8) (uint8, error) {
	return decr(c, k, n)
}

// decrement an item of type uint16 by n. Returns an error if the item's value
// is not an uint16, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrUint16(k string, n uint16) (uint16, error) {
	return decr(c, k, n)
}

// decrement an item of type uint32 by n. Returns an error if the item's value
// is not an uint32, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrUint32(k string, n uint32) (uint32, error) {
	return decr(c, k, n)
}

// decrement an item of type uint64 by n. Returns an error if the item's value
// is not an uint64, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrUint64(k string, n uint64) (uint64, error) {
	return decr(c, k, n)
}

// decrement an item of type float32 by n. Returns an error if the item's value
// is not an float32, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrFloat32(k string, n float32) (float32, error) {
	return decr(c, k, n)
}

// decrement an item of type float64 by n. Returns an error if the item's value
// is not an float64, or if it was not found. If there is no error, the
// decremented value is returned.
func (c *cache) DecrFloat64(k string, n float64) (float64, error) {
	return decr(c, k, n)
}

func incr[T Number](c *cache, k string, n T) (T, error) {
	var nv T

	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return nv, ErrValueNotFound
	}
	rv, ok := v.Value.(T)
	if !ok {
		c.mu.Unlock()
		return nv, fmt.Errorf("cache: the value is not an %T", nv)
	}
	nv = rv + n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}

func decr[T Number](c *cache, k string, n T) (T, error) {
	var nv T

	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return nv, ErrValueNotFound
	}
	rv, ok := v.Value.(T)
	if !ok {
		c.mu.Unlock()
		return nv, fmt.Errorf("cache: the value is not an %T", nv)
	}
	nv = rv - n
	v.Value = nv
	c.items[k] = v
	c.mu.Unlock()
	return nv, nil
}
