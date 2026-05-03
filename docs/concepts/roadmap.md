# Roadmap

The Roadmap tracks all features organized by status, providing a single source of truth for what's being worked on.

## Purpose

The Roadmap:

- Lists all features across the organization
- Tracks status (active, backlog, completed, archived)
- Defines priorities within each status
- Captures dependencies between features
- Records ownership and target versions

## Structure

```json
{
  "version": "1.0.0",
  "updated_at": "2025-01-15T00:00:00Z",
  "active": [ ... ],
  "backlog": [ ... ],
  "completed": [ ... ],
  "archived": [ ... ]
}
```

## Status Categories

### Active

Features currently being implemented:

```json
{
  "active": [
    {
      "feature": "user-authentication",
      "priority": 1,
      "started_at": "2025-01-15T00:00:00Z",
      "owner": "Platform Team",
      "target_version": "v2.5.0",
      "notes": "Critical path for all user features"
    },
    {
      "feature": "query-caching",
      "priority": 2,
      "started_at": "2025-01-20T00:00:00Z",
      "owner": "Performance Team",
      "target_version": "v2.5.0"
    }
  ]
}
```

Active features have:

| Field | Required | Description |
|-------|----------|-------------|
| `feature` | Yes | Feature directory name (kebab-case) |
| `priority` | Yes | Order within active (1 = highest) |
| `started_at` | No | When work began |
| `owner` | No | Team or person responsible |
| `target_version` | No | Planned release version |
| `notes` | No | Additional context |

### Backlog

Planned features not yet started:

```json
{
  "backlog": [
    {
      "feature": "data-export",
      "priority": 1,
      "estimated_version": "v2.6.0",
      "notes": "Requested by enterprise customers"
    },
    {
      "feature": "real-time-sync",
      "priority": 2,
      "blocked_by": ["query-caching"],
      "estimated_version": "v3.0.0",
      "notes": "Depends on caching infrastructure"
    }
  ]
}
```

Backlog features have:

| Field | Required | Description |
|-------|----------|-------------|
| `feature` | Yes | Feature directory name |
| `priority` | Yes | Order within backlog |
| `blocked_by` | No | Features that must complete first |
| `estimated_version` | No | Estimated release version |
| `notes` | No | Additional context |

### Completed

Features that have shipped:

```json
{
  "completed": [
    {
      "feature": "legacy-migration",
      "completed_at": "2025-01-10T00:00:00Z",
      "released_in": "v2.4.0",
      "notes": "Migrated from v1 schema"
    }
  ]
}
```

Completed features have:

| Field | Required | Description |
|-------|----------|-------------|
| `feature` | Yes | Feature directory name |
| `completed_at` | No | When implementation finished |
| `released_in` | No | Version that shipped it |
| `notes` | No | Additional context |

### Archived

Features that were cancelled or superseded:

```json
{
  "archived": [
    {
      "feature": "graphql-api",
      "archived_at": "2025-01-05T00:00:00Z",
      "reason": "Decided to focus on REST API improvements instead",
      "superseded_by": "query-caching"
    }
  ]
}
```

Archived features have:

| Field | Required | Description |
|-------|----------|-------------|
| `feature` | Yes | Feature directory name |
| `archived_at` | No | When archived |
| `reason` | Yes | Why it was archived |
| `superseded_by` | No | Feature that replaced it |

## Directory Mapping

Roadmap status maps to spec directories:

| Roadmap Status | Spec Directory |
|----------------|----------------|
| `active` | `specs/active/{feature}/` |
| `backlog` | `specs/backlog/{feature}/` |
| `completed` | `specs/completed/{feature}/` |
| `archived` | `specs/archived/{feature}/` |

When a feature moves status, use the CLI:

```bash
dfspec move user-authentication active
```

This moves both the directory and updates the roadmap.

## CLI Commands

### List Features

```bash
# All features
dfspec list

# By status
dfspec list active
dfspec list backlog

# JSON output
dfspec list --json
```

### Generate Markdown

```bash
dfspec generate roadmap
```

Creates `ROADMAP.md`:

```markdown
# Roadmap

## Active

| Priority | Feature | Owner | Target |
|----------|---------|-------|--------|
| 1 | [user-authentication](specs/active/user-authentication/) | Platform Team | v2.5.0 |

## Backlog

| Priority | Feature | Blocked By | Est. Version |
|----------|---------|------------|---------------|
| 1 | [data-export](specs/backlog/data-export/) | - | v2.6.0 |
```

## Best Practices

1. **Keep it updated** - Stale roadmaps lose trust
2. **Use priorities** - Makes ordering explicit
3. **Track blockers** - Surfaces dependencies
4. **Archive, don't delete** - History is valuable
5. **Include owners** - Accountability matters
6. **Add notes** - Context helps future readers
