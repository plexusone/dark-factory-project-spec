# dfspec init

Initialize a new feature.

## Synopsis

```bash
dfspec init <feature> [flags]
```

## Description

Creates a new feature directory in `specs/backlog/` with an MRD template. Also creates the corresponding docs directory.

## Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `feature` | Yes | Feature name (kebab-case) |

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--status` | `backlog` | Initial status: `backlog`, `active` |

## Examples

### Basic Usage

```bash
dfspec init user-authentication
```

Output:

```
Created feature 'user-authentication' in backlog/

Files created:
  specs/backlog/user-authentication/mrd.json
  docs/design/user-authentication/

Next steps:
  1. Edit specs/backlog/user-authentication/mrd.json to define requirements
  2. Run 'dfspec validate' to check validity
  3. When ready, run 'dfspec move user-authentication active' to start work
```

### Start Active

```bash
dfspec init urgent-fix --status active
```

Output:

```
Created feature 'urgent-fix' in active/

Files created:
  specs/active/urgent-fix/mrd.json
  docs/design/urgent-fix/

Next steps:
  1. Edit specs/active/urgent-fix/mrd.json to define requirements
  2. Create PRD: specs/active/urgent-fix/prd.json
  3. Run 'dfspec validate urgent-fix' to check validity
```

## Generated MRD Template

The created `mrd.json`:

```json
{
  "metadata": {
    "id": "MRD-user-authentication",
    "title": "user-authentication",
    "feature": "user-authentication",
    "version": "0.1.0",
    "status": "draft",
    "authors": [],
    "created_at": "2025-01-15T10:00:00Z",
    "updated_at": "2025-01-15T10:00:00Z"
  },
  "problem_statement": {
    "summary": "",
    "current_state": "",
    "desired_state": "",
    "impact": ""
  },
  "target_users": [],
  "success_metrics": [],
  "high_level_requirements": [],
  "constraints": [],
  "non_goals": [],
  "assumptions": [],
  "open_questions": []
}
```

## Naming Conventions

Feature names must be kebab-case:

- ✅ `user-authentication`
- ✅ `data-export-v2`
- ❌ `userAuthentication`
- ❌ `user_authentication`
- ❌ `User Authentication`
