# Roadmap Schema

Feature tracking schema.

## Go Type

```go
type Roadmap struct {
    Version   string             `json:"version"`
    UpdatedAt string             `json:"updated_at"`
    Active    []ActiveFeature    `json:"active"`
    Backlog   []BacklogFeature   `json:"backlog"`
    Completed []CompletedFeature `json:"completed"`
    Archived  []ArchivedFeature  `json:"archived"`
}
```

## Fields

### Top-Level

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | string | Yes | Roadmap version |
| `updated_at` | string | Yes | RFC3339 timestamp |
| `active` | ActiveFeature[] | Yes | Active features |
| `backlog` | BacklogFeature[] | Yes | Backlog features |
| `completed` | CompletedFeature[] | Yes | Completed features |
| `archived` | ArchivedFeature[] | Yes | Archived features |

### ActiveFeature

Features currently being implemented.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `feature` | string | Yes | Feature name (kebab-case) |
| `priority` | integer | Yes | Priority (1 = highest) |
| `started_at` | string | No | RFC3339 timestamp |
| `owner` | string | No | Responsible team/person |
| `target_version` | string | No | Target release |
| `notes` | string | No | Additional context |

### BacklogFeature

Planned features not yet started.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `feature` | string | Yes | Feature name |
| `priority` | integer | Yes | Priority (1 = highest) |
| `blocked_by` | string[] | No | Blocking feature names |
| `estimated_version` | string | No | Estimated release |
| `notes` | string | No | Additional context |

### CompletedFeature

Features that have shipped.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `feature` | string | Yes | Feature name |
| `completed_at` | string | No | RFC3339 timestamp |
| `released_in` | string | No | Release version |
| `notes` | string | No | Additional context |

### ArchivedFeature

Cancelled or superseded features.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `feature` | string | Yes | Feature name |
| `archived_at` | string | No | RFC3339 timestamp |
| `reason` | string | Yes | Why archived |
| `superseded_by` | string | No | Replacement feature |

## Example

```json
{
  "version": "1.0.0",
  "updated_at": "2025-01-15T00:00:00Z",
  "active": [
    {
      "feature": "user-authentication",
      "priority": 1,
      "owner": "Platform Team",
      "target_version": "v2.5.0"
    }
  ],
  "backlog": [
    {
      "feature": "data-export",
      "priority": 1,
      "blocked_by": ["user-authentication"],
      "estimated_version": "v2.6.0"
    }
  ],
  "completed": [
    {
      "feature": "legacy-migration",
      "released_in": "v2.4.0"
    }
  ],
  "archived": [
    {
      "feature": "graphql-api",
      "reason": "Decided to focus on REST API",
      "superseded_by": "query-caching"
    }
  ]
}
```

See [example-roadmap.json](https://github.com/plexusone/dark-factory-project-spec/blob/main/examples/example-roadmap.json).
