package matcher

const HttpMethodNum = 12

// MatcherHttp is a selector builder
type MatcherHttp struct {
	mns []*MatcherNode
}

func NewMatcherHttp() *MatcherHttp {
	m := &MatcherHttp{
		mns: make([]*MatcherNode, 0, HttpMethodNum),
	}
	m.mns = append(m.mns, NewMatcherNode(WildcardName))
	return m
}

func (m *MatcherHttp) Clone() *MatcherHttp {
	mr := &MatcherHttp{
		mns: make([]*MatcherNode, 0, len(m.mns)),
	}
	for _, mp := range m.mns {
		mr.mns = append(mr.mns, mp.Clone())
	}
	return mr
}

func (m *MatcherHttp) getOrNew(method string) *MatcherNode {
	for _, mn := range m.mns {
		if mn.name == method {
			return mn
		}
	}
	mn := NewMatcherNode(method)
	m.mns = append(m.mns, mn)
	return mn
}

func (m *MatcherHttp) get(method string) *MatcherNode {
	for _, mm := range m.mns {
		if mm.name == method {
			return mm
		}
	}
	return nil
}

// Exact is with Matcher's method, path
func (m *MatcherHttp) Exact(method string, paths ...string) *MatcherHttp {
	if len(paths) == 0 {
		return m
	}
	mm := m.getOrNew(method)
	mm.AddExacts(paths...)
	return m
}

// Prefix is with Matcher's method, prefix
func (m *MatcherHttp) Prefix(method string, prefixes ...string) *MatcherHttp {
	if len(prefixes) == 0 {
		return m
	}
	mm := m.getOrNew(method)
	mm.AddPrefixes(prefixes...)
	return m
}

// MustRegex is with Matcher's method, regex
func (m *MatcherHttp) MustRegex(method string, regexes ...string) *MatcherHttp {
	if len(regexes) == 0 {
		return m
	}
	mm := m.getOrNew(method)
	mm.MustAddRegex(regexes...)
	return m
}

// ExactWildcard is with Matcher's path
func (m *MatcherHttp) ExactWildcard(paths ...string) *MatcherHttp {
	return m.Exact(WildcardName, paths...)
}

// PrefixWildcard is with Matcher's prefix
func (m *MatcherHttp) PrefixWildcard(prefixes ...string) *MatcherHttp {
	return m.Prefix(WildcardName, prefixes...)
}

// RegexWildcard is with Matcher's regex on
func (m *MatcherHttp) RegexWildcard(regexes ...string) *MatcherHttp {
	return m.MustRegex(WildcardName, regexes...)
}

// ExactMultiMethod is with Matcher's method, path
func (m *MatcherHttp) ExactMultiMethod(path string, methods ...string) *MatcherHttp {
	for _, method := range methods {
		m.Exact(method, path)
	}
	return m
}

// PrefixMethods is with Matcher's method, prefix
func (m *MatcherHttp) PrefixMethods(prefix string, methods ...string) *MatcherHttp {
	for _, method := range methods {
		m.Prefix(method, prefix)
	}
	return m
}

// RegexMethods is with Matcher's method, regex
func (m *MatcherHttp) RegexMethods(regex string, methods ...string) *MatcherHttp {
	for _, method := range methods {
		m.MustRegex(method, regex)
	}
	return m
}

// Matches is match method, path
// GET ^/api/user/[^/]+$				--> GET /api/user/{id}
// GET ^/api/user/[^/]+/menu$ 			--> GET /api/user/{id}/menu
// GET ^/api/user/.+$ 					--> GET /api/user/a, /api/user/a/b
// GET ^/api/user/[^/]+/menu/.+$ 		--> GET /api/user/{id}/menu/a, /api/user/{id}/menu/a/b
func (m *MatcherHttp) Matches(method, path string) bool {
	return m.matches(WildcardName, path) || m.matches(method, path)
}

func (m *MatcherHttp) matches(method, path string) bool {
	if mn := m.get(method); mn != nil {
		return mn.Matches(path)
	}
	return false
}
