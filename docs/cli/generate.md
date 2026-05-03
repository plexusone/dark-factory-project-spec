# dfspec generate

Generate documentation from specs.

## Synopsis

```bash
dfspec generate <type> [flags]
```

## Subcommands

| Command | Description |
|---------|-------------|
| `roadmap` | Generate ROADMAP.md from roadmap.json |
| `index` | Generate index.md for feature docs |

## dfspec generate roadmap

Generates a human-readable `ROADMAP.md` from `specs/roadmap.json`.

### Synopsis

```bash
dfspec generate roadmap [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-o, --output` | `ROADMAP.md` | Output file path |

### Example

```bash
dfspec generate roadmap
```

Output:
```
Generated ROADMAP.md
```

Generated file:

```markdown
# Roadmap

_Generated from specs/roadmap.json_

## Active

| Priority | Feature | Owner | Target |
|----------|---------|-------|--------|
| 1 | [user-authentication](specs/active/user-authentication/) | Platform Team | v2.5.0 |

## Backlog

| Priority | Feature | Blocked By | Est. Version |
|----------|---------|------------|---------------|
| 1 | [data-export](specs/backlog/data-export/) | - | v2.6.0 |

## Completed

| Feature | Released In | Completed |
|---------|-------------|------------|
| [legacy-migration](specs/completed/legacy-migration/) | v2.4.0 | 2025-01-10 |

## Archived

| Feature | Reason | Superseded By |
|---------|--------|---------------|
| graphql-api | Decided to focus on REST API | query-caching |
```

### Custom Output

```bash
dfspec generate roadmap -o docs/ROADMAP.md
```

## dfspec generate index

Generates `index.md` files for feature documentation directories.

### Synopsis

```bash
dfspec generate index [feature] [flags]
```

### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `feature` | No | Feature to generate index for (all if omitted) |

### Example: All Features

```bash
dfspec generate index
```

Output:
```
Generated docs/design/user-authentication/index.md
Generated docs/design/query-caching/index.md
Generated docs/design/data-export/index.md
```

### Example: Single Feature

```bash
dfspec generate index user-authentication
```

Output:
```
Generated docs/design/user-authentication/index.md
```

### Generated Index

```markdown
# User Authentication

**Status:** active

## Specs

- [Market Requirements Document (MRD)](mrd.md) | [spec](../../specs/active/user-authentication/mrd.json)
- [Product Requirements Document (PRD)](prd.md) | [spec](../../specs/active/user-authentication/prd.json)
- [Spec Evaluation](evaluation.md) | [spec](../../specs/active/user-authentication/evaluation.json)
```
