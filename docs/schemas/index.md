# Schema Reference

Dark Factory uses JSON Schema for spec validation. Schemas are generated from Go types.

## Go-First Approach

Go structs are the source of truth:

```
types/*.go → JSON Schema → Validation
```

This ensures:

- Type safety in Go code
- Schema always matches Go types
- Single source of truth

## Available Schemas

| Schema | Purpose | Go Type |
|--------|---------|---------|
| [MRD](mrd.md) | Market Requirements Document | `types.MRD` |
| [PRD](prd.md) | Product Requirements Document | `types.PRD` |
| [Constitution](constitution.md) | Organization standards | `types.Constitution` |
| [Roadmap](roadmap.md) | Feature tracking | `types.Roadmap` |
| [Evaluation](evaluation.md) | Spec evaluation | `types.SpecEvaluation` |

## Schema Generation

Generate schemas from Go types:

```bash
go run cmd/generate-schema/main.go
```

Or use make:

```bash
make generate-schema
```

## Schema Location

Generated schemas are in `schemas/`:

```
schemas/
├── mrd.schema.json
├── prd.schema.json
├── constitution.schema.json
├── roadmap.schema.json
└── evaluation.schema.json
```

## Validation

### Using ajv-cli

```bash
# Install
npm install -g ajv-cli

# Validate
ajv validate -s schemas/mrd.schema.json -d examples/example-mrd.json
```

### Using Go

```go
import (
    "encoding/json"
    "github.com/plexusone/dark-factory-project-spec/types"
)

func ValidateMRD(data []byte) error {
    var mrd types.MRD
    return json.Unmarshal(data, &mrd)
}
```

## Schema Validation in CI

Add to your CI pipeline:

```yaml
- name: Validate schemas
  run: make validate-examples
```

## Common Patterns

### Required vs Optional Fields

Go struct tags control optionality:

```go
type Example struct {
    Required string `json:"required"`           // Required
    Optional string `json:"optional,omitempty"` // Optional
}
```

### Enums

Go string types with constants:

```go
type Status string

const (
    StatusDraft    Status = "draft"
    StatusApproved Status = "approved"
)
```

### Nested Objects

Go structs:

```go
type Parent struct {
    Child ChildType `json:"child"`
}

type ChildType struct {
    Field string `json:"field"`
}
```
