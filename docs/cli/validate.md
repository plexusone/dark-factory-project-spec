# dfspec validate

Validate specs against schemas.

## Synopsis

```bash
dfspec validate [feature] [flags]
```

## Description

Validates spec files are valid JSON and conform to expected structure. Without arguments, validates all features.

## Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `feature` | No | Feature to validate (validates all if omitted) |

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--strict` | `false` | Enable strict validation |

## Examples

### Validate Single Feature

```bash
dfspec validate user-authentication
```

Output (success):
```
✓ All specs valid
```

Output (failure):
```
✗ specs/active/user-authentication/mrd.json: invalid JSON at line 15
```

### Validate All Features

```bash
dfspec validate
```

Output:
```
Validating active/user-authentication... ✓
Validating active/query-caching... ✓
Validating backlog/data-export... ✓

✓ All specs valid (3 features)
```

### Strict Mode

```bash
dfspec validate --strict
```

Strict mode checks:

- All required fields present
- Field types match schema
- Cross-references valid (MRD references exist, etc.)
- Roadmap entries match directories

## Validation Checks

### Basic Validation

- JSON syntax valid
- Required fields present (`metadata.id`, `metadata.title`, etc.)

### Strict Validation

Additional checks:

| Check | Description |
|-------|-------------|
| Schema conformance | All fields match schema types |
| ID format | IDs follow `{TYPE}-{YEAR}-{NUMBER}` format |
| Cross-references | `mrd_reference` points to existing MRD |
| Roadmap sync | Features in roadmap have directories |
| Status consistency | Directory location matches expected status |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | All validations passed |
| 1 | Validation errors found |
