# 枚举定义规范 (Enum Specification)

> 规范版本: 3.0.0 · 表达形态: 参考 OpenAPI 3.1 Schema Object (JSON Schema 2020-12)
> 配套文件: [enums.schema.json](./enums.schema.json) (校验用 JSON Schema) · [enums.example.json](./enums.example.json) (完整示例)

## 1. 背景与目标

枚举值散落在服务端常量、proto、OpenAPI 文档、前端字典、系统字典等多个地方, 各写各的会导致: 取值不一致、中文标签不统一、成员缺文档、废弃值被误删。

本规范定义**枚举注册表 (Enum Registry)** —— 一种 JSON 文档格式, 作为枚举定义的唯一事实来源:

- **表达形态** = 参考 OpenAPI 3.1 Schema Object (`oneOf` + `const`): 枚举级字段均为 OpenAPI 原生关键字, 成员语义字段 `name`/`label`/`description`, 不含任何 `x-` 扩展; 映射到 `components.schemas.<枚举名>` 时按 §8 做字段投影;
- **约束边界** = 与 OpenAPI v3 规范一致, 不额外限制枚举名/成员名/取值的书写格式 (见 §4);
- **产出对象** = 技术写手、AI 生成器、代码生成器;
- **产物验收** = 必须通过 [enums.schema.json](./enums.schema.json) 校验, 且满足 §9 语义规则。

**消费方** (由注册表派生, 不各自手写):

| 消费方 | 派生产物 |
|---|---|
| 服务端代码 | 枚举类型/常量/标签映射/取值合法性校验 |
| OpenAPI 文档 | `components.schemas.<枚举名>` 枚举 Schema |
| 前端 | 字典 (value/label/deprecated 列表) |
| 系统字典 | 字典类型与字典项 |
| 文档/客服 | 中英对照取值表 |

## 2. 快速开始

三步产出一个合规的枚举注册表:

1. **命名**: 枚举名将作为 `components.schemas` 的 key, 只需满足 §4 的字符集; 成员名、取值的写法不设约束;
2. **填字段**: 每个枚举写 `type/title/oneOf` (integer 枚举可选 `format` 定宽), 每个成员写 `name/const/label` (完整骨架见 §3);
3. **校验**: 用 §10 的命令跑 [enums.schema.json](./enums.schema.json), 再过 §12 检查清单。

最小示例:

```json
{
  "$schema": "./enums.schema.json",
  "version": "1.0.0",
  "enums": {
    "Core_Gender": {
      "type": "integer",
      "title": "性别",
      "oneOf": [
        { "name": "Male", "const": 1, "label": "男" },
        { "name": "Female", "const": 2, "label": "女" },
        { "name": "Unknown", "const": 3, "label": "未知" }
      ]
    }
  }
}
```

## 3. 文档结构

注册表由三层构成, 每层只允许 schema 声明的字段 (校验器 `additionalProperties: false`, 拼错字段名会被拒绝):

```
注册表 (root)
├── $schema      可选, 指向 enums.schema.json
├── version      必填, 文档版本 semver
└── enums        必填, 枚举名 → 枚举定义
    └── 枚举定义 (OpenAPI 3.1 Schema Object 受限子集)
        ├── type / title / oneOf    必填
        ├── format                  integer 枚举可选, string 枚举禁用
        ├── description / default / deprecated   可选
        └── oneOf                   必填, 成员列表
            └── 成员 (含 const 取值)
                ├── name / const / label             必填
                └── description / deprecated         可选
```

**文件组织**: 按服务拆分注册表文件——共享枚举单独一份, 供多服务合并; 各服务一份 (如 `enums.bos.json` / `enums.bss.json`)。一个文件一个 `version`。

## 4. 命名与取值 (与 OpenAPI v3 一致)

枚举名、成员名、取值**均不设书写格式约束**——不要求 PascalCase、kebab-case 或任何命名风格, 与 OpenAPI v3 规范保持一致, 写法由团队自行统一。唯一的硬约束来自 OpenAPI 3.x Components Object 对 key 的要求:

| 对象 | 约束 | 说明 |
|---|---|---|
| 枚举名 (enums 键) | `^[a-zA-Z0-9\.\-_]+$` | 将作为 `components.schemas` 的 key, 字符集遵循 OpenAPI 规范; `Gender`、`Core_Gender`、`core.gender` 均合法 |
| 成员名 (成员 `name`) | 非空字符串 | 无格式约束 |
| 取值 (`const`) | 与 `type` 同型 | string 取任意字符串, integer 取任意整数, 无格式约束 |

## 5. 字段字典

枚举级字段均为 OpenAPI 3.1 原生关键字, 语义分配如下 (这是本规范唯一"约定俗成"的部分): 枚举级 `title` = 枚举中文名称; 成员级 `name` = 成员名, `label` = 中文标签 (UI 文案), `description` = 描述。

### 5.1 注册表根

| 字段 | 必填 | 类型 | 约束 | 说明 |
|---|---|---|---|---|
| `$schema` | 否 | string | — | 指向 `enums.schema.json`, 便于编辑器补全 |
| `version` | 是 | string | `^\d+\.\d+\.\d+$` | 注册表文档版本, 内容变更时递增 |
| `enums` | 是 | object | 至少 1 项; 键符合 §4 枚举名字符集 | 枚举名 → 枚举定义 |

### 5.2 枚举定义

| 字段 | 必填 | 类型 | 约束 | 说明 |
|---|---|---|---|---|
| `type` | 是 | string | `string` \| `integer` | 取值类型 |
| `format` | integer 枚举可选 | string | `int8`/`int16`/`int32`/`int64`/`uint8`/`uint16`/`uint32`/`uint64` | 取值宽度, 覆盖所有整型宽度; 不填 = 对应语言默认整数类型 (如 Go `int`); string 枚举**不写** `format` |
| `title` | 是 | string | 非空 | 中文名称, 如 `性别` |
| `description` | 否 | string | 非空 | 用途/适用范围补充说明 |
| `default` | 否 | 与 `type` 同型 | 必须等于某**非废弃**成员的 `const` | 默认值, 通常取"未知/未指定" |
| `deprecated` | 否 | boolean | — | 整个枚举废弃; 须在 `description` 注明替代 |
| `oneOf` | 是 | array | ≥1 项 | 成员列表, 顺序 = 展示/生成顺序 |

### 5.3 成员

| 字段 | 必填 | 类型 | 约束 | 说明 |
|---|---|---|---|---|
| `name` | 是 | string | 非空, 同枚举内唯一 | 成员名, 如 `Male`; 书写格式不约束 |
| `const` | 是 | 与 `type` 同型 | 全枚举唯一 | 线上取值 (wire value), 发布后不可改、不可复用 |
| `label` | 是 | string | 非空, 同枚举内互不相同 | 中文标签 (UI 文案), 如 `男` |
| `description` | 否 | string | 非空 | 描述/补充说明, 不入 UI |
| `deprecated` | 否 | boolean | — | 成员废弃; 取值保留不回收, 须在 `description` 注明去向 |

### 5.4 字段边界

- schema 采用**封闭字段集** (`additionalProperties: false`): 只收上表字段, 拼错/多写一律拒绝;
- **不使用 `x-` 扩展字段**: `x-go-name`/`x-aliases`/`x-enum-varnames` 之类都不入注册表; 需要新字段时先改 [enums.schema.json](./enums.schema.json) 再用。

## 6. 取值规范

### 6.1 类型选择

| 选 `integer` | 选 `string` |
|---|---|
| 状态/类型/级别等内部语义 | 跨端可见、被 URL/配置/日志直接使用的码值 |
| 存储紧凑、可比较排序 | 自描述、可读、跨系统稳定 |
| 值本身无需被人类直接读到 | 值会被人工排查/配置引用 |

- 同一业务概念**全链路只用一种类型**, 发布后不得改 `type`/`format`;
- 通用概念**复用**共享枚举 (如 `Core_Boolean`/`Core_Visible`/`Core_ClientPlatform`), 不另造近义枚举;
- 是否用枚举: 固定、可穷举、会出现在文档/UI 的离散取值才建枚举; 布尔语义直接用共享布尔枚举。

### 6.2 值稳定性 (最重要的一组规则)

1. `const` 一经发布: **不可修改、不可复用、不可删除**; 不再使用只标 `deprecated`;
2. 新增值**追加在列表末尾**, 不重排、不插空;
3. 同一枚举内 `const` 全体唯一。

### 6.3 default

- 可选; 若写, 必须等于某个非废弃成员的 `const`;
- 惯例: 没有"合理默认"就不写; 需要默认时优先取语义为"未知/未指定"的成员。

### 6.4 废弃

废弃成员: `deprecated: true` + `description` 注明去向/替代:

```json
{
  "name": "ResetPasswd",
  "const": "reset-passwd",
  "label": "重置密码",
  "description": "旧场景, 由 ResetLoginPasswd 取代",
  "deprecated": true
}
```

禁止: 删除废弃成员、把废弃成员的 `const` 复用给新成员。

## 7. 中文名称与标签

| 位置 | 规则 | 正例 | 反例 |
|---|---|---|---|
| 枚举 `title` | 中文名词, 不带"枚举/类型"冗余词 | `性别`、`公告状态` | `性别枚举`、`gender` |
| 成员 `label` | UI 短文案, 无句末标点; 专有名词保留原文 | `男`、`已发布`、`小程序` | `男性用户` (过度展开)、`Male` (未翻译)、`已发布。` |
| 成员 `description` | 补充说明, 自由句式, 仅描述不入 UI | `旧场景, 由 ResetLoginPasswd 取代` | 把 UI 文案写在这里 |

- 同一枚举内标签互不相同; 标签面向 UI, 详细语义放枚举级 `description` 或文档;
- 成员名不强制英文, 一旦确定不随文案调整, 保持稳定。

## 8. OpenAPI 3.1 映射

注册表结构与 OpenAPI 3.1 Schema Object 对齐 (`oneOf` + `const`), 但成员的 `name`/`label` 不是 OpenAPI 原生关键字, 放入 `components.schemas.<枚举名>` 前需做一次字段投影:

| 注册表字段 | OpenAPI 3.1 Schema Object | 说明 |
|---|---|---|
| 枚举 `type`/`format`/`title`/`description`/`default`/`deprecated` | 同名关键字 | 原生, 语义不变 |
| `oneOf` / 成员 `const` / 成员 `deprecated` | 同名关键字 | 原生, 语义不变 |
| 成员 `name` | 子 Schema `title` | 成员名 |
| 成员 `label` | 子 Schema `description` | 中文标签 (UI 文案) |
| 成员 `description` | 不投射 | 内部补充说明, 需要时并入 `label` 括号说明 |

```yaml
components:
  schemas:
    Core_Gender:
      type: integer
      format: uint32
      title: 性别          # ← 枚举 title
      oneOf:
        - title: Male      # ← 成员 name
          const: 1         # ← 成员 const
          description: 男  # ← 成员 label
        - title: Female
          const: 2
          description: 女
        - title: Unknown
          const: 3
          description: 未知
```

本规范只认 3.1 形态; 与 OpenAPI 3.0 及更早工具链的互通不在范围内。

## 9. 语义校验规则 (schema 之外)

JSON Schema 无法表达跨字段关系, 生成器与代码评审必须额外执行:

| # | 规则 |
|---|---|
| R1 | 同一枚举内 `const` 唯一 |
| R2 | 同一枚举内 `name` 唯一, `label` 互不相同 |
| R3 | `default` 必须是某个非废弃成员的 `const` |
| R4 | 成员只增不删: 废弃标 `deprecated`, 值不回收、不复用 |
| R5 | 已发布枚举的 `type`/`format` 不可变更; 需要换类型 = 新建枚举 |
| R6 | 同一业务概念只定义一次; 跨服务共享一份定义, 不复制粘贴 |

## 10. 校验方法

以下命令在仓库根目录 (proc/) 执行:

```sh
# 方式一: ajv-cli (Node.js)
npx --yes ajv-cli@5 validate --spec=draft2020 \
  -s enum_spec/enums.schema.json \
  -d enum_spec/enums.example.json

# 方式二: Python jsonschema (>= 4.17)
python3 - <<'PY'
import json
from jsonschema import Draft202012Validator

schema = json.load(open("enum_spec/enums.schema.json"))
Draft202012Validator.check_schema(schema)   # schema 自检
Draft202012Validator(schema).validate(      # 校验文档
    json.load(open("enum_spec/enums.example.json")))
print("OK")
PY
```

- 产出流程: 先跑 schema 校验, 再人工/工具过 §9 语义规则与 §12 清单;
- 方式三 (Go/CI): 用 `github.com/santhosh-tekuri/jsonschema/v6` 在 CI 中编译 schema (即自检) 后 `Validate` 注册表文档;
- CI 建议: 对注册表文件全量校验 + schema 自检; `version` 未递增可作 warning。

## 11. 常见错误 (反例)

| 错误 | 为何错 | 正确做法 |
|---|---|---|
| integer 枚举写 `"const": "1"` | 类型与 `type` 不符, 校验失败 | `const: 1` |
| `format` 写 `int`/`uint` 等非词表值 / string 枚举写 `format` | 默认整数类型用"不填"表达, 词表只收固定宽度; string 枚举禁用 `format` | 固定宽度按词表写, 默认整数类型不填 |
| 枚举名含空格、`/` 等字符 | 无法作为 `components.schemas` 的 key | 只用 `字母/数字/./-/_` |
| 成员缺 `label` 或标签写成英文长句 | UI 无法直接展示 | `label: 男` |
| 直接删除废弃成员 / 复用旧值 | 线上存量数据变脏 | `deprecated: true` + 标签注明去向 |
| `default: 5` 但成员里没有 5 | 默认值悬空 | 取某成员 `const`, 或不写 |

## 12. 写手 / AI 生成检查清单

- [ ] 枚举名只含 `字母/数字/./-/_` (可直接作 `components.schemas` key)
- [ ] 每个枚举: `type`/`title` 齐全; integer 枚举按需带 `format`, string 枚举不带
- [ ] 每个成员: `name` + `const` (与 `type` 同型) + `label` (中文标签)
- [ ] `const`/`name`/标签在枚举内唯一; 新增值追加在列表末尾
- [ ] 未出现任何 `x-` 扩展字段
- [ ] `default` (如有) 是非废弃成员的 `const`; 废弃成员 (如有) 标 `deprecated` 且 `description` 注明去向
- [ ] 通用概念复用共享枚举, 不新建近义枚举
- [ ] 通过 `enums.schema.json` 校验
- [ ] `version` 已按变更递增

**AI 生成提示词模板** (复制后填空):

```text
# 角色
你是枚举定义生成器, 输出必须通过 enum_spec/enums.schema.json 校验 (规则见 enum_spec/enum-spec.md)。

# 任务
为以下业务概念生成枚举注册表 JSON:
- <枚举名>: <中文名称 / 取值清单(成员名=取值=标签)>

# 硬性规则
1. 枚举名仅含 字母/数字/./-/_ ; 成员名与取值的写法不设约束
2. integer 枚举 format 可选 (int8/int16/int32/int64/uint8/uint16/uint32/uint64, 不填 = 语言默认整数类型); string 枚举不写 format
3. 每个成员必须有简短中文标签 (label)
4. 只输出 schema 定义的字段 (type/format/title/description/default/deprecated/oneOf/name/const/label), 任何 x- 扩展字段禁止出现
5. 仅输出一个 JSON 代码块, 不输出其它解释
```

## 附录. 完整示例

见 [enums.example.json](./enums.example.json) —— 覆盖 integer (`core.gender`)、string 词表 (`core.client_platform`)、废弃成员 (`sys.temp_grant_scene`)、命名空间取值 (`sys.sse_channel`) 四种典型形态, 可直接用作 §10 校验命令的输入。
