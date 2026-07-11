package matcher

import (
	"regexp"
	"testing"
)

// TestMatcher_Matches 测试 Matcher.Matches 方法, 覆盖 HTTP 方法和各种路由匹配
func TestMatcher_Matches(t *testing.T) {
	// 构建一个常用的 Matcher, 包含多种匹配方式
	m := NewMatcherHttp()
	m.Exact("GET", "/api/users", "/api/users/detail").
		Exact("POST", "/api/users").
		Exact("PUT", "/api/users").
		Exact("DELETE", "/api/users/1").
		Exact("PATCH", "/api/users/1").
		Exact("HEAD", "/api/users").
		Exact("OPTIONS", "/api/users").
		Prefix("GET", "/api/orders/").
		Prefix("POST", "/api/files/upload/").
		Regex("GET", `^/api/items/\d+$`).
		Regex("GET", `^/api/v[12]/search(\?.*)?$`).
		// 数字 id 参数路由: /api/user/{id}/menu, /api/user/{id}/menu/{menuId}
		Regex("GET", `^/api/user/\d+/menu$`).
		Regex("GET", `^/api/user/\d+/menu/\d+$`).
		Regex("POST", `^/api/user/\d+/menu$`).
		Regex("DELETE", `^/api/user/\d+/menu/\d+$`).
		Regex("PUT", `^/api/user/\d+/menu/\d+$`).
		// 字符串 id 参数路由: /api/user/{id} id 可以是任意非空非斜杠字符串
		Regex("GET", `^/api/user/[^/]+$`).
		Regex("GET", `^/api/user/[^/]+/profile$`).
		Regex("POST", `^/api/user/[^/]+$`).
		Regex("PUT", `^/api/user/[^/]+$`).
		Regex("DELETE", `^/api/user/[^/]+$`).
		// 通配路由 /api/user/{id}/menu/**  匹配 menu 下任意子路径
		Regex("GET", `^/api/user/\d+/menu/.+$`).
		Regex("POST", `^/api/user/\d+/menu/.+$`).
		Prefix("GET", "/api/static/").
		Prefix("GET", "/api/user/123/files/")

	tests := []struct {
		name   string
		method string
		path   string
		want   bool
	}{
		// ========== GET ==========
		{name: "GET 精确匹配 /api/users", method: "GET", path: "/api/users", want: true},
		{name: "GET 精确匹配 /api/users/detail", method: "GET", path: "/api/users/detail", want: true},
		{name: "GET 前缀匹配 /api/orders/123", method: "GET", path: "/api/orders/123", want: true},
		{name: "GET 前缀匹配 /api/orders/", method: "GET", path: "/api/orders/", want: true},
		{name: "GET 正则匹配 /api/items/42", method: "GET", path: "/api/items/42", want: true},
		{name: "GET 正则匹配 /api/items/0", method: "GET", path: "/api/items/0", want: true},
		{name: "GET 正则匹配 /api/v1/search", method: "GET", path: "/api/v1/search", want: true},
		{name: "GET 正则匹配 /api/v2/search?q=abc", method: "GET", path: "/api/v2/search?q=abc", want: true},
		{name: "GET 不匹配 /api/items/abc (非数字)", method: "GET", path: "/api/items/abc", want: false},
		{name: "GET 不匹配 /api/items/ (无数字)", method: "GET", path: "/api/items/", want: false},
		{name: "GET 不匹配 /api/orders (无尾斜杠)", method: "GET", path: "/api/orders", want: false},
		{name: "GET 不匹配 /api/v3/search (版本不存在)", method: "GET", path: "/api/v3/search", want: false},
		// 多层参数路由 /api/user/{id}/menu
		{name: "GET /api/user/1/menu (单层参数)", method: "GET", path: "/api/user/1/menu", want: true},
		{name: "GET /api/user/123/menu (多位id)", method: "GET", path: "/api/user/123/menu", want: true},
		{name: "GET /api/user/1/menu/99 (双层参数)", method: "GET", path: "/api/user/1/menu/99", want: true},
		{name: "GET /api/user/0/menu/0 (零值id)", method: "GET", path: "/api/user/0/menu/0", want: true},
		{name: "GET 不匹配 /api/user//menu (空id)", method: "GET", path: "/api/user//menu", want: false},
		{name: "GET 不匹配 /api/user/abc/menu (非数字id)", method: "GET", path: "/api/user/abc/menu", want: false},
		{name: "GET 不匹配 /api/user/1/menu/ (尾斜杠)", method: "GET", path: "/api/user/1/menu/", want: false},
		{name: "GET /api/user/1/menu/abc (通配命中)", method: "GET", path: "/api/user/1/menu/abc", want: true},
		{name: "GET 不匹配 /api/user/1/menus (拼写不同)", method: "GET", path: "/api/user/1/menus", want: false},
		{name: "GET /api/user/1/menu/1/extra (通配命中)", method: "GET", path: "/api/user/1/menu/1/extra", want: true},

		// 字符串 id 参数路由
		{name: "GET /api/user/abc123 (字符串id)", method: "GET", path: "/api/user/abc123", want: true},
		{name: "GET /api/user/550e8400-e29b (uuid风格id)", method: "GET", path: "/api/user/550e8400-e29b", want: true},
		{name: "GET /api/user/user_name (下划线id)", method: "GET", path: "/api/user/user_name", want: true},
		{name: "GET /api/user/abc (短字符串id)", method: "GET", path: "/api/user/abc", want: true},
		{name: "GET /api/user/123 (数字id也命中字符串规则)", method: "GET", path: "/api/user/123", want: true},
		{name: "GET /api/user/abc123/profile (字符串id+子路径)", method: "GET", path: "/api/user/abc123/profile", want: true},
		{name: "GET 不匹配 /api/user// (空id)", method: "GET", path: "/api/user/", want: false},
		{name: "GET 不匹配 /api/user/abc/def (额外斜杠)", method: "GET", path: "/api/user/abc/def", want: false},

		// 通配路由 /api/user/{id}/menu/**
		{name: "GET /api/user/1/menu/1/extra (通配)", method: "GET", path: "/api/user/1/menu/1/extra", want: true},
		{name: "GET /api/user/1/menu/a/b/c (深层通配)", method: "GET", path: "/api/user/1/menu/a/b/c", want: true},
		{name: "GET /api/user/123/menu/abc (字符串子路径)", method: "GET", path: "/api/user/123/menu/abc", want: true},
		{name: "GET /api/user/1/menu/some-id/nested (通配多层)", method: "GET", path: "/api/user/1/menu/some-id/nested", want: true},
		// /api/user/1/menu 命中精确正则 ^/api/user/\d+/menu$, 非通配规则
		{name: "GET /api/user/1/menu (精确正则命中, 非通配)", method: "GET", path: "/api/user/1/menu", want: true},

		// 前缀匹配补充
		{name: "GET 前缀 /api/static/css/main.css", method: "GET", path: "/api/static/css/main.css", want: true},
		{name: "GET 前缀 /api/user/123/files/avatar.png", method: "GET", path: "/api/user/123/files/avatar.png", want: true},
		{name: "GET 前缀 /api/user/123/files/a/b/c.txt (深层前缀)", method: "GET", path: "/api/user/123/files/a/b/c.txt", want: true},
		{name: "GET 不匹配 /api/user/456/files/x (不同id)", method: "GET", path: "/api/user/456/files/x", want: false},

		{name: "GET 不匹配 /other/path", method: "GET", path: "/other/path", want: false},
		{name: "GET 不匹配空路径", method: "GET", path: "", want: false},

		// ========== POST ==========
		{name: "POST 精确匹配 /api/users", method: "POST", path: "/api/users", want: true},
		{name: "POST 前缀匹配 /api/files/upload/image.png", method: "POST", path: "/api/files/upload/image.png", want: true},
		{name: "POST 前缀匹配 /api/files/upload/", method: "POST", path: "/api/files/upload/", want: true},
		{name: "POST 不匹配 /api/files/other", method: "POST", path: "/api/files/other", want: false},
		{name: "POST 不匹配 /api/users/1", method: "POST", path: "/api/users/1", want: false},
		// 多层参数路由 POST /api/user/{id}/menu
		{name: "POST /api/user/1/menu 创建菜单", method: "POST", path: "/api/user/1/menu", want: true},
		{name: "POST /api/user/123/menu 创建菜单(多位id)", method: "POST", path: "/api/user/123/menu", want: true},
		{name: "POST 不匹配 /api/user/abc/menu (非数字id)", method: "POST", path: "/api/user/abc/menu", want: false},
		{name: "POST /api/user/1/menu/99 (通配命中)", method: "POST", path: "/api/user/1/menu/99", want: true},
		// 字符串 id
		{name: "POST /api/user/new-user (字符串id创建)", method: "POST", path: "/api/user/new-user", want: true},
		{name: "POST /api/user/abc_123 (带下划线)", method: "POST", path: "/api/user/abc_123", want: true},
		// 通配路由
		{name: "POST /api/user/1/menu/a/b (通配)", method: "POST", path: "/api/user/1/menu/a/b", want: true},
		{name: "POST /api/user/1/menu/extra/path/deep (深层通配)", method: "POST", path: "/api/user/1/menu/extra/path/deep", want: true},

		// ========== PUT ==========
		{name: "PUT 精确匹配 /api/users", method: "PUT", path: "/api/users", want: true},
		{name: "PUT 不匹配 /api/users/1", method: "PUT", path: "/api/users/1", want: false},
		// 多层参数路由 PUT /api/user/{id}/menu/{menuId}
		{name: "PUT /api/user/1/menu/99 更新菜单", method: "PUT", path: "/api/user/1/menu/99", want: true},
		{name: "PUT 不匹配 /api/user/1/menu (缺menuId)", method: "PUT", path: "/api/user/1/menu", want: false},
		{name: "PUT 不匹配 /api/user/abc/menu/99 (非数字id)", method: "PUT", path: "/api/user/abc/menu/99", want: false},
		// 字符串 id
		{name: "PUT /api/user/uuid-123 (字符串id更新)", method: "PUT", path: "/api/user/uuid-123", want: true},
		{name: "PUT /api/user/test_user (下划线id)", method: "PUT", path: "/api/user/test_user", want: true},

		// ========== DELETE ==========
		{name: "DELETE 精确匹配 /api/users/1", method: "DELETE", path: "/api/users/1", want: true},
		{name: "DELETE 不匹配 /api/users", method: "DELETE", path: "/api/users", want: false},
		{name: "DELETE 不匹配 /api/users/2", method: "DELETE", path: "/api/users/2", want: false},
		// 多层参数路由 DELETE /api/user/{id}/menu/{menuId}
		{name: "DELETE /api/user/1/menu/99 删除菜单", method: "DELETE", path: "/api/user/1/menu/99", want: true},
		{name: "DELETE /api/user/123/menu/456 删除菜单(多位id)", method: "DELETE", path: "/api/user/123/menu/456", want: true},
		{name: "DELETE 不匹配 /api/user/1/menu (缺menuId)", method: "DELETE", path: "/api/user/1/menu", want: false},
		{name: "DELETE 不匹配 /api/user/1/menu/abc (非数字menuId)", method: "DELETE", path: "/api/user/1/menu/abc", want: false},
		// 字符串 id
		{name: "DELETE /api/user/user-abc (字符串id删除)", method: "DELETE", path: "/api/user/user-abc", want: true},
		{name: "DELETE /api/user/very-long-uuid-here (长字符串id)", method: "DELETE", path: "/api/user/very-long-uuid-here", want: true},

		// ========== PATCH ==========
		{name: "PATCH 精确匹配 /api/users/1", method: "PATCH", path: "/api/users/1", want: true},
		{name: "PATCH 不匹配 /api/users", method: "PATCH", path: "/api/users", want: false},

		// ========== HEAD ==========
		{name: "HEAD 精确匹配 /api/users", method: "HEAD", path: "/api/users", want: true},
		{name: "HEAD 不匹配 /api/users/1", method: "HEAD", path: "/api/users/1", want: false},

		// ========== OPTIONS ==========
		{name: "OPTIONS 精确匹配 /api/users", method: "OPTIONS", path: "/api/users", want: true},
		{name: "OPTIONS 不匹配 /api/other", method: "OPTIONS", path: "/api/other", want: false},

		// ========== 未注册方法 ==========
		{name: "TRACE 未注册返回 false", method: "TRACE", path: "/api/users", want: false},
		{name: "CONNECT 未注册返回 false", method: "CONNECT", path: "/api/users", want: false},
		{name: "小写 get 不匹配 GET", method: "get", path: "/api/users", want: false},

		// ========== 边界场景 ==========
		{name: "空 method 不匹配", method: "", path: "/api/users", want: false},
		{name: "空 path 不匹配", method: "GET", path: "", want: false},
		{name: "空 method 和空 path", method: "", path: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := m.Matches(tt.method, tt.path)
			if got != tt.want {
				t.Errorf("Matches(%q, %q) = %v, want %v", tt.method, tt.path, got, tt.want)
			}
		})
	}
}

// TestMethodPathMatch_Matches 测试 MethodPathMatch.Matches 的三种匹配路径
func TestMethodPathMatch_Matches(t *testing.T) {
	tests := []struct {
		name string
		mp   *MatcherNode
		path string
		want bool
	}{
		// ========== nil 接收者 ==========
		{name: "nil 接收者返回 false", mp: nil, path: "/any", want: false},

		// ========== 精确路径匹配 ==========
		{
			name: "精确路径命中",
			mp:   &MatcherNode{name: "GET", exact: map[string]struct{}{"/api/users": {}, "/api/orders": {}}},
			path: "/api/users",
			want: true,
		},
		{
			name: "精确路径不命中",
			mp:   &MatcherNode{name: "GET", exact: map[string]struct{}{"/api/users": {}}},
			path: "/api/users/1",
			want: false,
		},
		{
			name: "空路径列表不匹配",
			mp:   &MatcherNode{name: "GET", exact: map[string]struct{}{}},
			path: "/api/users",
			want: false,
		},

		// ========== 前缀匹配 ==========
		{
			name: "前缀匹配命中",
			mp:   &MatcherNode{name: "GET", prefixes: []string{"/api/"}},
			path: "/api/users",
			want: true,
		},
		{
			name: "前缀匹配 - 完全相同也命中",
			mp:   &MatcherNode{name: "GET", prefixes: []string{"/api/users"}},
			path: "/api/users",
			want: true,
		},
		{
			name: "前缀匹配不命中 - 无公共前缀",
			mp:   &MatcherNode{name: "GET", prefixes: []string{"/api/"}},
			path: "/other/path",
			want: false,
		},
		{
			name: "前缀匹配 - 多个前缀取第一个命中",
			mp:   &MatcherNode{name: "GET", prefixes: []string{"/v1/", "/v2/"}},
			path: "/v2/users",
			want: true,
		},

		// ========== 正则匹配 ==========
		{
			name: "正则匹配命中",
			mp: func() *MatcherNode {
				r := `^/api/items/\d+$`
				return &MatcherNode{
					name:    "GET",
					regexes: []string{r},
					rxs:     compileRegexps([]string{r}),
				}
			}(),
			path: "/api/items/123",
			want: true,
		},
		{
			name: "正则匹配不命中",
			mp: func() *MatcherNode {
				r := `^/api/items/\d+$`
				return &MatcherNode{
					name:    "GET",
					regexes: []string{r},
					rxs:     compileRegexps([]string{r}),
				}
			}(),
			path: "/api/items/abc",
			want: false,
		},
		{
			name: "多个正则 - 任一命中即可",
			mp: func() *MatcherNode {
				rs := []string{`^/a$`, `^/b$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/b",
			want: true,
		},

		// ========== 混合匹配 ==========
		{
			name: "精确路径和前缀都配置 - 精确命中",
			mp:   &MatcherNode{name: "GET", exact: map[string]struct{}{"/api/users": {}}, prefixes: []string{"/api/"}},
			path: "/api/users",
			want: true,
		},
		{
			name: "精确路径和前缀都配置 - 前缀命中",
			mp:   &MatcherNode{name: "GET", exact: map[string]struct{}{"/api/users": {}}, prefixes: []string{"/api/"}},
			path: "/api/orders",
			want: true,
		},
		{
			name: "三种方式都配置 - 仅正则命中",
			mp: func() *MatcherNode {
				rs := []string{`^/items/\d+$`}
				return &MatcherNode{
					name:     "GET",
					exact:    map[string]struct{}{"/home": {}},
					prefixes: []string{"/api/"},
					regexes:  rs,
					rxs:      compileRegexps(rs),
				}
			}(),
			path: "/items/99",
			want: true,
		},
		{
			name: "三种方式都配置 - 无一命中",
			mp: func() *MatcherNode {
				rs := []string{`^/items/\d+$`}
				return &MatcherNode{
					name:     "GET",
					exact:    map[string]struct{}{"/home": {}},
					prefixes: []string{"/api/"},
					regexes:  rs,
					rxs:      compileRegexps(rs),
				}
			}(),
			path: "/other/abc",
			want: false,
		},

		// ========== 多层参数路由正则 ==========
		{
			name: "正则 /api/user/{id}/menu 命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/42/menu",
			want: true,
		},
		{
			name: "正则 /api/user/{id}/menu 不命中 - 空id",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user//menu",
			want: false,
		},
		{
			name: "正则 /api/user/{id}/menu/{menuId} 命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/\d+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu/99",
			want: true,
		},
		{
			name: "正则 /api/user/{id}/menu/{menuId} 不命中 - 尾斜杠",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/\d+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu/99/",
			want: false,
		},
		{
			name: "混合 - 精确+前缀+多层参数正则, 仅正则命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu$`}
				return &MatcherNode{
					name:     "GET",
					exact:    map[string]struct{}{"/api/users": {}},
					prefixes: []string{"/api/orders/"},
					regexes:  rs,
					rxs:      compileRegexps(rs),
				}
			}(),
			path: "/api/user/5/menu",
			want: true,
		},

		// ========== 字符串 id 参数路由 ==========
		{
			name: "正则 /api/user/{id} 字符串id 命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/[^/]+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/abc123",
			want: true,
		},
		{
			name: "正则 /api/user/{id} uuid风格 命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/[^/]+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/550e8400-e29b-41d4-a716-446655440000",
			want: true,
		},
		{
			name: "正则 /api/user/{id} 空id不命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/[^/]+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/",
			want: false,
		},
		{
			name: "正则 /api/user/{id}/profile 字符串id+子路径 命中",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/[^/]+/profile$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/abc-uuid/profile",
			want: true,
		},

		// ========== 通配路由 /api/user/{id}/menu/** ==========
		{
			name: "正则 /api/user/{id}/menu/** 命中单层子路径",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/.+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu/extra",
			want: true,
		},
		{
			name: "正则 /api/user/{id}/menu/** 命中深层子路径",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/.+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu/a/b/c/d",
			want: true,
		},
		{
			name: "正则 /api/user/{id}/menu/** 不命中 - 无子路径",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/.+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu",
			want: false,
		},
		{
			name: "正则 /api/user/{id}/menu/** 不命中 - 尾斜杠不算子路径内容",
			mp: func() *MatcherNode {
				rs := []string{`^/api/user/\d+/menu/.+$`}
				return &MatcherNode{
					name:    "GET",
					regexes: rs,
					rxs:     compileRegexps(rs),
				}
			}(),
			path: "/api/user/1/menu/",
			want: false,
		},
		{
			name: "通配 + 字符串id 前缀匹配 /api/user/{id}/files/**",
			mp: &MatcherNode{
				name:     "GET",
				prefixes: []string{"/api/user/123/files/"},
			},
			path: "/api/user/123/files/a/b/c.txt",
			want: true,
		},
		{
			name: "前缀匹配不命中 - 不同id",
			mp: &MatcherNode{
				name:     "GET",
				prefixes: []string{"/api/user/123/files/"},
			},
			path: "/api/user/456/files/a.txt",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Matches(tt.path)
			if got != tt.want {
				t.Errorf("MethodPathMatch.Matches(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

// TestMatcher_Chain 测试链式调用
func TestMatcher_Chain(t *testing.T) {
	m := NewMatcherHttp().
		Exact("GET", "/a").
		Prefix("GET", "/b/").
		Regex("GET", `^/c/\d+$`).
		Exact("POST", "/d")

	if !m.Matches("GET", "/a") {
		t.Error("链式调用 Path 未生效")
	}
	if !m.Matches("GET", "/b/1") {
		t.Error("链式调用 Prefix 未生效")
	}
	if !m.Matches("GET", "/c/1") {
		t.Error("链式调用 Regex 未生效")
	}
	if !m.Matches("POST", "/d") {
		t.Error("链式调用跨 method Path 未生效")
	}
}

// TestMatcher_MultipleMethods 同一 method 下多次注册都生效
func TestMatcher_MultipleMethods(t *testing.T) {
	m := NewMatcherHttp()
	m.Exact("GET", "/first")
	m.Exact("GET", "/second")
	m.Prefix("GET", "/api/")
	m.Regex("GET", `^/v\d+/ok$`)

	tests := []struct {
		name string
		path string
		want bool
	}{
		{name: "第一次 Path 注册", path: "/first", want: true},
		{name: "第二次 Path 注册", path: "/second", want: true},
		{name: "Prefix 注册", path: "/api/anything", want: true},
		{name: "Regex 注册", path: "/v2/ok", want: true},
		{name: "不匹配", path: "/nope", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.Matches("GET", tt.path); got != tt.want {
				t.Errorf("Matches(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

// TestMatcher_RegexPanic 测试无效正则会 panic
func TestMatcher_RegexPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("无效正则未触发 panic")
		}
	}()
	NewMatcherHttp().Regex("GET", `[invalid`)
}

// TestMatcher_Clone 测试 Clone 方法, 克隆后独立且匹配一致
func TestMatcher_Clone(t *testing.T) {
	orig := NewMatcherHttp().
		Exact("GET", "/api/users", "/api/orders").
		Prefix("GET", "/api/files/").
		Regex("GET", `^/api/items/\d+$`).
		Exact("POST", "/api/users").
		Regex("POST", `^/api/user/\d+/menu$`)

	clone := orig.Clone()

	// 克隆能匹配原 Matcher 所有规则
	tests := []struct {
		method, path string
		want         bool
	}{
		{"GET", "/api/users", true},
		{"GET", "/api/orders", true},
		{"GET", "/api/files/a.txt", true},
		{"GET", "/api/items/42", true},
		{"GET", "/api/items/abc", false},
		{"POST", "/api/users", true},
		{"POST", "/api/user/1/menu", true},
		{"POST", "/api/user/abc/menu", false},
		{"DELETE", "/api/users", false},
	}
	for _, tt := range tests {
		got := clone.Matches(tt.method, tt.path)
		if got != tt.want {
			t.Errorf("Clone.Matches(%q, %q) = %v, want %v", tt.method, tt.path, got, tt.want)
		}
	}

	// 克隆后修改不影响原对象
	clone.Exact("GET", "/clone-only")
	clone.Prefix("DELETE", "/clone/")
	clone.Regex("PUT", `^/clone/\d+$`)

	if orig.Matches("GET", "/clone-only") {
		t.Error("修改克隆后, 原 Matcher 不应匹配 /clone-only")
	}
	if orig.Matches("DELETE", "/clone/x") {
		t.Error("修改克隆后, 原 Matcher 不应匹配 DELETE /clone/x")
	}
	if orig.Matches("PUT", "/clone/1") {
		t.Error("修改克隆后, 原 Matcher 不应匹配 PUT /clone/1")
	}

	// 克隆自身应能匹配新增规则
	if !clone.Matches("GET", "/clone-only") {
		t.Error("克隆应匹配自己新增的 /clone-only")
	}
	if !clone.Matches("DELETE", "/clone/x") {
		t.Error("克隆应匹配自己新增的 DELETE /clone/x")
	}
	if !clone.Matches("PUT", "/clone/1") {
		t.Error("克隆应匹配自己新增的 PUT /clone/1")
	}
}

// TestMatcher_CloneNilMethod 测试克隆空 Matcher
func TestMatcher_CloneEmpty(t *testing.T) {
	orig := NewMatcherHttp()
	clone := orig.Clone()

	if clone.Matches("GET", "/any") {
		t.Error("空 Matcher 克隆后不应匹配任何路径")
	}

	// 克隆后添加规则不影响原对象
	clone.Exact("GET", "/added")
	if orig.Matches("GET", "/added") {
		t.Error("修改克隆后, 原空 Matcher 不应受影响")
	}
	if !clone.Matches("GET", "/added") {
		t.Error("克隆应匹配自己新增的 /added")
	}
}

// TestMatcher_AnyMethod 测试 AnyMethod 通配方法匹配
func TestMatcher_AnyMethod(t *testing.T) {
	t.Run("PathAnyMethod 所有方法命中", func(t *testing.T) {
		m := NewMatcherHttp().ExactWildcard("/health", "/ping")
		methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE", "CONNECT"}
		for _, method := range methods {
			if !m.Matches(method, "/health") {
				t.Errorf("PathAnyMethod: Matches(%q, %q) = false, want true", method, "/health")
			}
			if !m.Matches(method, "/ping") {
				t.Errorf("PathAnyMethod: Matches(%q, %q) = false, want true", method, "/ping")
			}
		}
	})

	t.Run("PrefixAnyMethod 所有方法命中", func(t *testing.T) {
		m := NewMatcherHttp().PrefixWildcard("/public/")
		methods := []string{"GET", "POST", "PUT", "DELETE"}
		for _, method := range methods {
			if !m.Matches(method, "/public/css/style.css") {
				t.Errorf("PrefixAnyMethod: Matches(%q, %q) = false, want true", method, "/public/css/style.css")
			}
		}
	})

	t.Run("RegexAnyMethod 所有方法命中", func(t *testing.T) {
		m := NewMatcherHttp().RegexWildcard(`^/api/v\d+/health$`)
		methods := []string{"GET", "POST", "PUT", "DELETE"}
		for _, method := range methods {
			if !m.Matches(method, "/api/v1/health") {
				t.Errorf("RegexAnyMethod: Matches(%q, %q) = false, want true", method, "/api/v1/health")
			}
			if !m.Matches(method, "/api/v2/health") {
				t.Errorf("RegexAnyMethod: Matches(%q, %q) = false, want true", method, "/api/v2/health")
			}
		}
	})

	t.Run("AnyMethod 不影响不匹配的路径", func(t *testing.T) {
		m := NewMatcherHttp().ExactWildcard("/health")
		if m.Matches("GET", "/other") {
			t.Error("AnyMethod 不应匹配未注册的路径")
		}
	})

	t.Run("AnyMethod + 具体方法共存", func(t *testing.T) {
		m := NewMatcherHttp().
			ExactWildcard("/health").
			Exact("GET", "/api/users")

		// AnyMethod 路径对所有方法命中
		if !m.Matches("POST", "/health") {
			t.Error("AnyMethod /health 应对 POST 命中")
		}
		if !m.Matches("GET", "/health") {
			t.Error("AnyMethod /health 应对 GET 命中")
		}

		// 具体方法路径只对该方法命中
		if !m.Matches("GET", "/api/users") {
			t.Error("GET /api/users 应命中")
		}
		if m.Matches("POST", "/api/users") {
			t.Error("POST /api/users 不应命中(仅注册了 GET)")
		}
	})
}

// TestMatcher_AnyMethodWithSpecificRules 测试 AnyMethod 和具体方法规则的优先级
func TestMatcher_AnyMethodWithSpecificRules(t *testing.T) {
	m := NewMatcherHttp().
		ExactWildcard("/health").
		PrefixWildcard("/public/").
		RegexWildcard(`^/metrics/\d+$`).
		Exact("GET", "/api/users").
		Prefix("POST", "/api/files/").
		Regex("DELETE", `^/api/items/\d+$`)

	tests := []struct {
		name   string
		method string
		path   string
		want   bool
	}{
		// AnyMethod Path
		{name: "AnyMethod Path GET", method: "GET", path: "/health", want: true},
		{name: "AnyMethod Path POST", method: "POST", path: "/health", want: true},
		{name: "AnyMethod Path PUT", method: "PUT", path: "/health", want: true},
		{name: "AnyMethod Path 未注册方法", method: "CUSTOM", path: "/health", want: true},

		// AnyMethod Prefix
		{name: "AnyMethod Prefix GET", method: "GET", path: "/public/a.js", want: true},
		{name: "AnyMethod Prefix POST", method: "POST", path: "/public/b.css", want: true},

		// AnyMethod Regex
		{name: "AnyMethod Regex GET", method: "GET", path: "/metrics/123", want: true},
		{name: "AnyMethod Regex PUT", method: "PUT", path: "/metrics/456", want: true},
		{name: "AnyMethod Regex 不命中", method: "GET", path: "/metrics/abc", want: false},

		// 具体方法规则
		{name: "GET /api/users 命中", method: "GET", path: "/api/users", want: true},
		{name: "POST /api/users 不命中", method: "POST", path: "/api/users", want: false},
		{name: "POST /api/files/upload 命中", method: "POST", path: "/api/files/upload", want: true},
		{name: "GET /api/files/upload 不命中", method: "GET", path: "/api/files/upload", want: false},
		{name: "DELETE /api/items/1 命中", method: "DELETE", path: "/api/items/1", want: true},
		{name: "GET /api/items/1 不命中", method: "GET", path: "/api/items/1", want: false},

		// 不匹配
		{name: "完全不匹配", method: "GET", path: "/unknown", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := m.Matches(tt.method, tt.path)
			if got != tt.want {
				t.Errorf("Matches(%q, %q) = %v, want %v", tt.method, tt.path, got, tt.want)
			}
		})
	}
}

// TestMatcher_CloneWithAnyMethod 测试克隆包含 AnyMethod 规则的 Matcher
func TestMatcher_CloneWithAnyMethod(t *testing.T) {
	orig := NewMatcherHttp().
		ExactWildcard("/health").
		PrefixWildcard("/public/").
		RegexWildcard(`^/metrics/\d+$`).
		Exact("GET", "/api/users")

	clone := orig.Clone()

	// 克隆保留 AnyMethod 规则
	if !clone.Matches("POST", "/health") {
		t.Error("克隆应保留 AnyMethod Path 规则")
	}
	if !clone.Matches("PUT", "/public/a.js") {
		t.Error("克隆应保留 AnyMethod Prefix 规则")
	}
	if !clone.Matches("DELETE", "/metrics/1") {
		t.Error("克隆应保留 AnyMethod Regex 规则")
	}
	if !clone.Matches("GET", "/api/users") {
		t.Error("克隆应保留具体方法规则")
	}

	// 克隆后修改不影响原对象
	clone.ExactWildcard("/clone-health")
	if orig.Matches("GET", "/clone-health") {
		t.Error("修改克隆 AnyMethod 后, 原 Matcher 不应受影响")
	}
	if !clone.Matches("GET", "/clone-health") {
		t.Error("克隆应匹配自己新增的 AnyMethod 路径")
	}
}

// TestMatcher_MultiMethod 测试 MultiMethod 系列方法
func TestMatcher_MultiMethod(t *testing.T) {
	t.Run("PathMultiMethod", func(t *testing.T) {
		m := NewMatcherHttp().ExactMultiMethod("/api/users", "GET", "POST", "PUT")
		if !m.Matches("GET", "/api/users") {
			t.Error("PathMultiMethod GET 未生效")
		}
		if !m.Matches("POST", "/api/users") {
			t.Error("PathMultiMethod POST 未生效")
		}
		if !m.Matches("PUT", "/api/users") {
			t.Error("PathMultiMethod PUT 未生效")
		}
		if m.Matches("DELETE", "/api/users") {
			t.Error("PathMultiMethod DELETE 不应命中")
		}
	})

	t.Run("PrefixMultiMethod", func(t *testing.T) {
		m := NewMatcherHttp().PrefixMethods("/api/files/", "GET", "POST")
		if !m.Matches("GET", "/api/files/a.txt") {
			t.Error("PrefixMultiMethod GET 未生效")
		}
		if !m.Matches("POST", "/api/files/b.txt") {
			t.Error("PrefixMultiMethod POST 未生效")
		}
		if m.Matches("PUT", "/api/files/c.txt") {
			t.Error("PrefixMultiMethod PUT 不应命中")
		}
	})

	t.Run("RegexMultiMethod", func(t *testing.T) {
		m := NewMatcherHttp().RegexMethods(`^/api/items/\d+$`, "GET", "DELETE")
		if !m.Matches("GET", "/api/items/1") {
			t.Error("RegexMultiMethod GET 未生效")
		}
		if !m.Matches("DELETE", "/api/items/2") {
			t.Error("RegexMultiMethod DELETE 未生效")
		}
		if m.Matches("POST", "/api/items/3") {
			t.Error("RegexMultiMethod POST 不应命中")
		}
	})
}

// compileRegexps 辅助函数, 编译正则切片
func compileRegexps(patterns []string) []*regexp.Regexp {
	rs := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		rs = append(rs, regexp.MustCompile(p))
	}
	return rs
}
