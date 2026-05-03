# dfspec list

List features by status.

## Synopsis

```bash
dfspec list [status] [flags]
```

## Description

Lists all features, optionally filtered by status. Features are read from `roadmap.json` if it exists, otherwise from directory structure.

## Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `status` | No | Filter by status: `active`, `backlog`, `completed`, `archived` |

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--json` | `false` | Output as JSON |

## Examples

### List All Features

```bash
dfspec list
```

Output:

```
Active (2):
  • user-authentication
    Priority: 1
    Owner: Platform Team
    Version: v2.5.0
  • query-caching
    Priority: 2
    Owner: Performance Team

Backlog (1):
  • data-export
    Priority: 1
    Version: v2.6.0

Completed (1):
  • legacy-migration
    Version: v2.4.0

Archived (1):
  • graphql-api
    Note: Decided to focus on REST API instead
```

### List Active Only

```bash
dfspec list active
```

Output:

```
Active (2):
  • user-authentication
    Priority: 1
    Owner: Platform Team
    Version: v2.5.0
  • query-caching
    Priority: 2
    Owner: Performance Team
```

### JSON Output

```bash
dfspec list --json
```

Output:

```json
{
  "active": [
    {
      "name": "user-authentication",
      "status": "active",
      "priority": 1,
      "owner": "Platform Team",
      "version": "v2.5.0"
    }
  ],
  "backlog": [],
  "completed": [],
  "archived": []
}
```

### JSON with Filter

```bash
dfspec list active --json
```

Output:

```json
{
  "active": [
    {
      "name": "user-authentication",
      "status": "active",
      "priority": 1,
      "owner": "Platform Team",
      "version": "v2.5.0"
    }
  ]
}
```
