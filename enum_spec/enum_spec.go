package enum_spec

import (
	"encoding/json"
	"iter"
	"slices"
	"strconv"
	"strings"
)

const Version = "1.0.0"

const (
	TypeInteger = "integer"
	TypeString  = "string"
)

type T struct {
	Version string `json:"version,omitempty" yaml:"version,omitempty"` // Required
	Info    *Info  `json:"info,omitempty" yaml:"info,omitempty"`
	Enums   *Enums `json:"enums,omitempty" yaml:"enums,omitempty"`
}

type Info struct {
	Version     string   `json:"version,omitempty" yaml:"version,omitempty"`
	Title       string   `json:"title,omitempty" yaml:"title,omitempty"`
	Summary     string   `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
	Contact     *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Url   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type Enums struct {
	m map[string]*Enumerate
}

func NewEnums() *Enums {
	return &Enums{m: make(map[string]*Enumerate)}
}

func (e *Enums) Keys() []string {
	out := make([]string, 0, len(e.m))
	for k := range e.m {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

func (e *Enums) Value(key string) *Enumerate {
	if e.Len() == 0 {
		return nil
	}
	return e.m[key]
}

func (e *Enums) Set(key string, value *Enumerate) *Enums {
	if e.m == nil {
		e.m = make(map[string]*Enumerate)
	}
	e.m[key] = value
	return e
}

func (e *Enums) Len() int {
	if e == nil || e.m == nil {
		return 0
	}
	return len(e.m)
}

func (e *Enums) Maps() map[string]*Enumerate {
	if e == nil {
		return nil
	}
	return e.m
}

func (e *Enums) All() iter.Seq2[string, *Enumerate] {
	return func(yield func(string, *Enumerate) bool) {
		if e == nil {
			return
		}
		for k, v := range e.m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (e *Enums) Values() iter.Seq[*Enumerate] {
	return func(yield func(*Enumerate) bool) {
		if e == nil {
			return
		}
		for _, v := range e.m {
			if !yield(v) {
				return
			}
		}
	}
}

func (e *Enums) MarshalJSON() ([]byte, error) {
	if e == nil || e.m == nil {
		return []byte("null"), nil
	}
	return json.Marshal(e.m)
}
func (e *Enums) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.m)
}

type Enumerate struct {
	Type        string            `json:"type"`        // 枚举数据类型, string, integer
	Format      string            `json:"format"`      // 枚举类型真实类型, string: 就是string, integer: 真实的整数类型,
	Description string            `json:"description"` // 枚举注释
	Explain     string            `json:"explain"`     // 枚举详细解释
	Oneof       []*EnumerateValue `json:"oneof"`       // 枚举项
}

type EnumerateValue struct {
	GoName   string `json:"goName"` // 枚举项定义名称
	Name     string `json:"name"`   // 枚举项名称, 已去掉枚举项定义名称的前缀
	Const    string `json:"const"`  // 枚举项值, string: unquote string.
	Label    string `json:"label"`  // 枚举项的标签
	RawValue string `json:"value"`  // 枚举项的原始值, string: quote string.
}

type EnumerateValueSlices []*EnumerateValue

// Explain convert to explain string format [0:aaa,1:bbb,3:ccc]
func (vs EnumerateValueSlices) Explain() string {
	if len(vs) == 0 {
		return "[]"
	}
	b := strings.Builder{}
	b.WriteString("[")
	for i, k := range vs {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(k.RawValue)
		b.WriteString(":")
		b.WriteString(strconv.Quote(k.Label))
	}
	b.WriteString("]")
	return b.String()
}
