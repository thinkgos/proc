package matcher

import (
	"maps"
	"regexp"
	"slices"
	"strings"
)

type MatcherNode struct {
	name     string
	exact    map[string]struct{}
	prefixes []string
	regexes  []string
	rxs      []*regexp.Regexp
}

func NewMatcherNode(name string) *MatcherNode {
	return &MatcherNode{name: name}
}

func (mn *MatcherNode) Name() string {
	if mn == nil {
		return "<nil>"
	}
	return mn.name
}

func (mn *MatcherNode) Clone() *MatcherNode {
	if mn == nil {
		return nil
	}
	return &MatcherNode{
		name:     mn.name,
		exact:    maps.Clone(mn.exact),
		prefixes: slices.Clone(mn.prefixes),
		regexes:  slices.Clone(mn.regexes),
		rxs:      slices.Clone(mn.rxs),
	}
}

func (mn *MatcherNode) Matches(v string) bool {
	return mn.MatchExact(v) || mn.MatchPrefix(v) || mn.MatchRegex(v)
}

func (mn *MatcherNode) MatchExact(v string) bool {
	if mn == nil || mn.exact == nil {
		return false
	}
	_, ok := mn.exact[v]
	return ok
}

func (mn *MatcherNode) MatchPrefix(v string) bool {
	return mn != nil && slices.ContainsFunc(mn.prefixes, func(p string) bool { return strings.HasPrefix(v, p) })
}

func (mn *MatcherNode) MatchRegex(v string) bool {
	return mn != nil && slices.ContainsFunc(mn.rxs, func(r *regexp.Regexp) bool { return r.MatchString(v) })
}

func (mn *MatcherNode) AddExacts(vs ...string) *MatcherNode {
	if mn.exact == nil {
		mn.exact = make(map[string]struct{})
	}
	for _, v := range vs {
		mn.exact[v] = struct{}{}
	}
	return mn
}

func (mn *MatcherNode) AddPrefixes(vs ...string) *MatcherNode {
	mn.prefixes = append(mn.prefixes, vs...)
	return mn
}

func (mn *MatcherNode) MustAddRegex(vs ...string) *MatcherNode {
	mn.regexes = append(mn.regexes, vs...)
	for _, r := range vs {
		mn.rxs = append(mn.rxs, regexp.MustCompile(r))
	}
	return mn
}

func (mn *MatcherNode) AddRegex(v string) error {
	r, err := regexp.Compile(v)
	if err != nil {
		return err
	}
	mn.regexes = append(mn.regexes, v)
	mn.rxs = append(mn.rxs, r)
	return nil
}
