# Constitution

The Constitution is the organization-wide source of truth for standards and decisions that apply across all projects.

## Purpose

The Constitution:

- Establishes approved technologies and patterns
- Defines quality and security requirements
- Standardizes file organization
- Documents development processes
- Provides defaults that specs inherit

## Structure

```json
{
  "metadata": {
    "version": "1.0.0",
    "effective_date": "2025-01-01T00:00:00Z",
    "approved_by": ["CTO", "Engineering Lead"],
    "last_reviewed": "2025-01-01T00:00:00Z"
  },
  "tech_stack": { ... },
  "architecture": { ... },
  "file_organization": { ... },
  "quality": { ... },
  "security": { ... },
  "documentation": { ... },
  "process": { ... }
}
```

## Sections

### Tech Stack

Defines approved technologies:

```json
{
  "tech_stack": {
    "languages": [
      {
        "name": "Go",
        "versions": ">=1.21",
        "use_case": "Backend services",
        "status": "preferred"
      },
      {
        "name": "TypeScript",
        "versions": ">=5.0",
        "use_case": "Frontend applications",
        "status": "preferred"
      }
    ],
    "frameworks": [
      {
        "name": "React",
        "category": "frontend",
        "status": "preferred"
      }
    ],
    "databases": [
      {
        "name": "PostgreSQL",
        "status": "preferred"
      }
    ],
    "prohibited": [
      {
        "name": "MongoDB",
        "reason": "Consistency requirements",
        "alternative": "PostgreSQL with JSONB"
      }
    ]
  }
}
```

Technology status levels:

| Status | Meaning |
|--------|---------|
| `preferred` | Default choice for new projects |
| `approved` | Allowed but not default |
| `legacy` | Allowed for existing projects only |
| `experimental` | Requires explicit approval |
| `deprecated` | Being phased out |

### Architecture

Defines architectural standards:

```json
{
  "architecture": {
    "patterns": [
      {
        "name": "Hexagonal Architecture",
        "use_case": "Backend services",
        "description": "Ports and adapters pattern"
      }
    ],
    "principles": [
      {
        "name": "Single Responsibility",
        "statement": "Each module should have one reason to change"
      }
    ],
    "api_standards": {
      "style": "REST",
      "versioning": "URL path (/v1/, /v2/)",
      "authentication": "JWT Bearer tokens"
    }
  }
}
```

### File Organization

Defines project structure:

```json
{
  "file_organization": {
    "specs": {
      "root_path": "specs/",
      "roadmap_path": "specs/roadmap.json",
      "constitution_path": "specs/constitution.json",
      "active_path": "specs/active/",
      "backlog_path": "specs/backlog/",
      "completed_path": "specs/completed/",
      "archived_path": "specs/archived/"
    },
    "docs": {
      "root_path": "docs/",
      "design_path": "docs/design/"
    },
    "naming_conventions": {
      "version_format": "^v\\d+\\.\\d+\\.\\d+$",
      "feature_format": "^[a-z][a-z0-9]*(-[a-z0-9]+)*$"
    }
  }
}
```

### Quality Standards

Defines testing and review requirements:

```json
{
  "quality": {
    "testing": {
      "unit_test_coverage": "80%",
      "integration_test_required": true,
      "e2e_test_required": true
    },
    "code_review": {
      "required_approvals": 2,
      "require_owner_approval": true,
      "automated_checks": ["lint", "test", "build"]
    }
  }
}
```

### Security Standards

Defines security requirements:

```json
{
  "security": {
    "authentication": {
      "methods": ["JWT", "API Key"],
      "mfa_required": true
    },
    "authorization": {
      "model": "RBAC",
      "default_deny": true
    },
    "data_protection": {
      "encryption_at_rest": true,
      "encryption_in_transit": true
    }
  }
}
```

## Constitution Overrides

Individual specs can override constitution defaults with justification:

```json
{
  "constitution_ref": "specs/constitution.json",
  "constitution_overrides": {
    "constitution_version": "1.0.0",
    "overrides": [
      {
        "path": "tech_stack.languages",
        "original_value": "Go",
        "override_value": "Rust",
        "justification": "Performance-critical component",
        "approved_by": ["CTO"],
        "approved_at": "2025-01-15T00:00:00Z",
        "expires_at": "2026-01-15T00:00:00Z"
      }
    ]
  }
}
```

Override fields:

| Field | Description |
|-------|-------------|
| `path` | JSON path being overridden |
| `original_value` | What constitution specifies |
| `override_value` | What this spec uses instead |
| `justification` | Why the override is needed |
| `approved_by` | Who approved the override |
| `expires_at` | When override should be reviewed |

## Best Practices

1. **Review regularly** - Set `next_review` date and honor it
2. **Document rationale** - Explain why each standard exists
3. **Allow exceptions** - Define clear exception processes
4. **Version carefully** - Breaking changes need migration plans
5. **Keep focused** - Don't over-specify; leave room for judgment
