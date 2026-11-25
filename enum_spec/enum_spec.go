package enum_spec

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 枚举取值类型, 对应 OpenAPI type。
const (
	TypeInteger = "integer"
	TypeString  = "string"
)

// T 枚举注册表文档。
type T struct {
	Schema  string                `json:"$schema,omitempty"` // 可选, 指向 enums.schema.json
	Version string                `json:"version"`           // 文档版本 (semver)
	Enums   map[string]*Enumerate `json:"enums"`             // 枚举名 → 枚举定义
}

// Enumerate 枚举定义
type Enumerate struct {
	Type        string            `json:"type"`                  // 取值类型: string | integer
	Format      string            `json:"format,omitempty"`      // 取值宽度: int8/int16/int32/int64/uint8/uint16/uint32/uint64, 仅 integer 枚举, 可选; 不填 = 对应语言默认整数类型
	Title       string            `json:"title"`                 // 枚举中文名称
	Description string            `json:"description,omitempty"` // 用途/适用范围补充说明
	Default     *Valuer           `json:"default,omitempty"`     // 默认值, 须等于某非废弃成员的 const
	Deprecated  bool              `json:"deprecated,omitempty"`  // 整个枚举废弃
	OneOf       []*EnumerateValue `json:"oneOf"`                 // 成员列表, 顺序 = 展示/生成顺序
}

// Explain 派生展示串 [<const>:<label>,...], 仅用于注释/备注, 不入注册表。
func (e *Enumerate) Explain() string {
	if len(e.OneOf) == 0 {
		return "[]"
	}
	b := strings.Builder{}
	b.WriteString("[")
	for i, v := range e.OneOf {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(v.Const.Literal())
		b.WriteString(":")
		b.WriteString(v.Label)
	}
	b.WriteString("]")
	return b.String()
}

// EnumerateValue 枚举成员。
type EnumerateValue struct {
	Name        string `json:"name"`                  // 成员名
	Const       Valuer `json:"const"`                 // 线上取值, 发布后不可修改、不可复用
	Label       string `json:"label"`                 // 中文标签 (UI 文案)
	Description string `json:"description,omitempty"` // 描述
	Deprecated  bool   `json:"deprecated,omitempty"`  // 成员废弃
}

type Valuer interface {
	fmt.Stringer
	Literal() string
	IsString() bool
}

var _ Valuer = (*String)(nil)
var _ Valuer = (*Integer)(nil)

type String struct {
	value string
}

// StringValue string 取值。
func StringValue(s string) Valuer { return &String{value: s} }

// String 返回取值的字符串形式: string 取值本身, integer 取值为十进制。
func (s String) String() string {
	return strconv.Quote(s.value)
}

// Literal implements [Valuer].
func (s String) Literal() string {
	return s.value
}

// IsString implements [Valuer].
func (s String) IsString() bool {
	return true
}

// MarshalJSON string 序列化为 JSON 字符串, integer 序列化为 JSON 数字。
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

// UnmarshalJSON 仅接受 JSON 字符串或整数, 其它类型报错。
func (s *String) UnmarshalJSON(data []byte) error {
	var ss string

	if err := json.Unmarshal(data, &ss); err != nil {
		return fmt.Errorf("enum const want string, %w", err)
	}
	s.value = ss
	return nil
}

type Integer struct {
	value int64
}

// IntegerValue integer 取值。
func IntegerValue(i int64) Valuer { return &Integer{value: i} }

func (i Integer) String() string {
	return strconv.FormatInt(i.value, 10)
}

// Literal implements [Valuer].
func (i Integer) Literal() string {
	return strconv.FormatInt(i.value, 10)
}

// IsString implements [Valuer].
func (i *Integer) IsString() bool {
	return false
}

func (i Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	var ii int64

	if err := json.Unmarshal(data, &ii); err != nil {
		return fmt.Errorf("enum const want integer, %w", err)
	}
	i.value = ii
	return nil
}
