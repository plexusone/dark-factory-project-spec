# CLI Reference

The `dfspec` CLI manages specs through their lifecycle.

## Installation

```bash
go install github.com/plexusone/dark-factory-project-spec/cmd/dfspec@latest
```

## Global Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--specs-dir` | `specs` | Path to specs directory |
| `--docs-dir` | `docs` | Path to docs directory |
| `-h, --help` | | Show help |
| `-v, --version` | | Show version |

## Commands

| Command | Description |
|---------|-------------|
| [`list`](list.md) | List features by status |
| [`init`](init.md) | Initialize a new feature |
| [`move`](move.md) | Move feature between statuses |
| [`validate`](validate.md) | Validate specs |
| [`generate`](generate.md) | Generate documentation |

## Quick Examples

```bash
# List all features
dfspec list

# List only active features
dfspec list active

# Create a new feature
dfspec init my-feature

# Move to active
dfspec move my-feature active

# Validate a feature
dfspec validate my-feature

# Generate roadmap
dfspec generate roadmap
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error (invalid args, validation failure, etc.) |
