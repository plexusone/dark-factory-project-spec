# Constitution Template

This template defines the organization-wide standards and decisions that apply across all projects. Individual specs inherit from and may override (with justification) these defaults.

## Purpose

The constitution serves as the single source of truth for:

- **Technology choices** that don't need to be re-debated per project
- **Quality standards** that apply uniformly
- **Security requirements** that must be met
- **Process standards** for consistency

## When to Create a Constitution

Create a constitution when:

- Multiple projects share the same tech stack decisions
- You want to codify organizational standards
- New projects should inherit defaults without re-specifying everything
- You need governance over technology choices

## Sections

### 1. Metadata

```json
{
  "metadata": {
    "version": "1.0.0",
    "effective_date": "RFC3339 timestamp",
    "approved_by": ["Role/Name"],
    "last_reviewed": "RFC3339 timestamp",
    "next_review": "RFC3339 timestamp (optional)"
  }
}
```

**Guidance:**

- Version follows semantic versioning
- Review periodically (recommend quarterly)
- Track approvers for governance

### 2. Tech Stack

**Purpose:** Define approved technologies to prevent fragmentation and ensure expertise.

```json
{
  "tech_stack": {
    "languages": [],
    "frameworks": [],
    "databases": [],
    "infrastructure": [],
    "libraries": [],
    "prohibited": []
  }
}
```

**Technology entry structure:**
```json
{
  "name": "Technology Name",
  "category": "backend|frontend|data|etc",
  "versions": ">=1.0.0",
  "use_case": "When to use this",
  "rationale": "Why this is approved",
  "status": "preferred|approved|legacy|experimental|deprecated",
  "owner": "Team/Person responsible"
}
```

**Status definitions:**

| Status | Meaning |
|--------|---------|
| `preferred` | Default choice for new projects |
| `approved` | Allowed but not the default |
| `legacy` | Allowed for existing projects only |
| `experimental` | Requires explicit approval |
| `deprecated` | Being phased out |

**Prohibited technologies:**
```json
{
  "name": "Technology",
  "reason": "Why prohibited",
  "alternative": "What to use instead",
  "exception_process": "How to get an exception"
}
```

### 3. Architecture

**Purpose:** Define architectural patterns and principles.

```json
{
  "architecture": {
    "patterns": [],
    "principles": [],
    "api_standards": {},
    "data_standards": {}
  }
}
```

**Pattern structure:**
```json
{
  "name": "Pattern Name",
  "description": "What it is",
  "use_case": "When to apply",
  "reference": "Link to documentation"
}
```

**Principle structure:**
```json
{
  "name": "Principle Name",
  "statement": "The principle itself",
  "rationale": "Why it matters",
  "examples": ["Concrete examples"]
}
```

**API Standards:**
- Style (REST, GraphQL, gRPC)
- Versioning strategy
- Authentication method
- Documentation requirements
- Rate limiting policy

**Data Standards:**
- Naming conventions
- Retention policies
- Privacy requirements
- Encryption standards

### 4. File Organization

**Purpose:** Define where machine specs and human docs live in projects.

**Directory Structure:**

```
project/
├── specs/                              # Machine-readable specs
│   ├── roadmap.json                    # Feature prioritization
│   ├── constitution.json               # Org standards
│   ├── active/                         # Features being worked on
│   │   ├── query-caching/
│   │   │   ├── mrd.json
│   │   │   ├── prd.json
│   │   │   └── evaluation.json
│   │   └── user-authentication/
│   │       └── mrd.json
│   ├── backlog/                        # Planned features
│   │   └── data-export/
│   │       └── mrd.json
│   ├── completed/                      # Finished features
│   │   └── legacy-migration/
│   │       ├── mrd.json
│   │       ├── prd.json
│   │       └── evaluation.json
│   └── archived/                       # Abandoned features
│       └── graphql-api/
│           └── mrd.json
└── docs/                               # Human-readable docs
    └── design/
        ├── query-caching/              # Mirrors specs/active/query-caching/
        │   ├── index.md
        │   ├── mrd.md
        │   └── prd.md
        └── legacy-migration/           # Mirrors specs/completed/legacy-migration/
            ├── mrd.md
            └── prd.md
```

**Configuration:**

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
      "archived_path": "specs/archived/",
      "feature_files": {
        "mrd_file": "mrd.json",
        "prd_file": "prd.json",
        "evaluation_file": "evaluation.json"
      }
    },
    "docs": {
      "root_path": "docs/",
      "design_path": "docs/design/",
      "versioned_design_pattern": "docs/design/{version}/",
      "feature_design_pattern": "docs/design/{feature}/",
      "adr_path": "docs/adr/",
      "runbooks_path": "docs/runbooks/"
    },
    "spec_to_doc_mapping": {
      "strategy": "feature",
      "feature_field": "metadata.feature",
      "doc_file_extension": ".md",
      "include_spec_type_prefix": false,
      "generate_index": true
    },
    "naming_conventions": {
      "version_format": "^v\\d+\\.\\d+\\.\\d+$",
      "version_example": "v1.2.3",
      "feature_format": "^[a-z][a-z0-9]*(-[a-z0-9]+)*$",
      "feature_example": "user-authentication",
      "spec_file_format": "{type}.json",
      "spec_file_example": "mrd.json"
    }
  }
}
```

**Specs Directory (`specs/`):**

| Path | Contents |
|------|----------|
| `specs/roadmap.json` | Feature prioritization and status |
| `specs/constitution.json` | Organization standards |
| `specs/active/{feature}/` | Features currently being worked on |
| `specs/backlog/{feature}/` | Planned features not yet started |
| `specs/completed/{feature}/` | Finished features |
| `specs/archived/{feature}/` | Abandoned or superseded features |

**Feature Directory Contents:**

Each feature directory contains:

| File | Purpose |
|------|---------|
| `mrd.json` | Market Requirements Document |
| `prd.json` | Product Requirements Document |
| `evaluation.json` | Spec evaluation results |

**Spec to Doc Mapping:**

Specs map to docs by feature name:

| Spec Location | Doc Location |
|---------------|--------------|
| `specs/active/query-caching/mrd.json` | `docs/design/query-caching/mrd.md` |
| `specs/active/query-caching/prd.json` | `docs/design/query-caching/prd.md` |
| `specs/completed/legacy-migration/mrd.json` | `docs/design/legacy-migration/mrd.md` |

**Naming Rules:**

| Type | Format | Example |
|------|--------|---------|
| Feature directory | `kebab-case` | `user-authentication` |
| Spec files | `{type}.json` | `mrd.json`, `prd.json` |
| Doc files | `{type}.md` | `mrd.md`, `prd.md` |
| Version directory | `vX.Y.Z` (semver) | `v1.2.3` |

**Why organize by feature?**

- Related specs (MRD + PRD) live together
- Mirrors how humans think about work
- Status is visible via parent directory (`active/`, `completed/`)
- Easy to see full context for any feature

**Why separate machine specs from human docs?**

- Machine specs are validated against schemas; human docs are freeform
- Specs are consumed by automation; docs are read by humans
- Different review/approval processes
- Clearer ownership and purpose

### 5. Quality

**Purpose:** Define quality and testing requirements.

```json
{
  "quality": {
    "testing": {
      "unit_test_coverage": "80%",
      "integration_test_required": true,
      "e2e_test_required": true,
      "performance_test_required": true,
      "security_test_required": true,
      "test_frameworks": []
    },
    "code_review": {
      "required_approvals": 2,
      "require_owner_approval": true,
      "automated_checks": [],
      "max_pr_size": 500
    },
    "metrics": []
  }
}
```

**Metric structure:**
```json
{
  "name": "Metric Name",
  "description": "What it measures",
  "target": "Target value",
  "measurement": "How to measure"
}
```

### 6. Security

**Purpose:** Define security requirements and compliance.

```json
{
  "security": {
    "authentication": {
      "methods": [],
      "mfa_required": true,
      "session_timeout": "8 hours",
      "password_policy": "..."
    },
    "authorization": {
      "model": "RBAC|ABAC|Cedar",
      "default_deny": true,
      "audit_required": true
    },
    "data_protection": {
      "encryption_at_rest": true,
      "encryption_in_transit": true,
      "pii_handling": "...",
      "secret_management": "..."
    },
    "compliance": ["SOC 2", "GDPR", "etc"],
    "vulnerability_scanning": {
      "dependency_scanning": true,
      "container_scanning": true,
      "sast_required": true,
      "dast_required": true,
      "max_critical_vulnerabilities": 0
    }
  }
}
```

### 7. Documentation

**Purpose:** Define documentation requirements.

```json
{
  "documentation": {
    "readme_required": true,
    "api_doc_required": true,
    "architecture_doc_required": true,
    "runbook_required": true,
    "changelog_required": true,
    "format": "Markdown|etc"
  }
}
```

### 8. Process

**Purpose:** Define development process standards.

```json
{
  "process": {
    "branching_strategy": "trunk-based|gitflow|etc",
    "commit_convention": "Conventional Commits v1.0.0",
    "release_process": "Semantic versioning with...",
    "incident_response": "...",
    "on_call_required": true
  }
}
```

## Overriding the Constitution

When a project needs to deviate from the constitution:

1. **Document the override explicitly**
2. **Justify the deviation**
3. **Get appropriate approval**
4. **Set an expiration if temporary**

Override structure in MRD:
```json
{
  "constitution_ref": "constitution-v1.0.0",
  "constitution_overrides": {
    "constitution_version": "1.0.0",
    "overrides": [
      {
        "path": "tech_stack.databases",
        "original_value": "PostgreSQL preferred",
        "override_value": "MongoDB",
        "justification": "Document-oriented model required",
        "approved_by": ["VP Engineering"],
        "approved_at": "2025-01-15T10:00:00Z",
        "expires_at": "2026-01-15T00:00:00Z"
      }
    ]
  }
}
```

## Versioning the Constitution

- **Patch** (1.0.X): Clarifications, typo fixes, no material changes
- **Minor** (1.X.0): New technologies added, non-breaking changes
- **Major** (X.0.0): Breaking changes, removed technologies, policy changes

When updating:

1. Create draft with proposed changes
2. Review with stakeholders
3. Announce deprecations (if any) with timeline
4. Publish new version
5. Update effective_date
6. Communicate changes to teams

## Best Practices

1. **Start small**: Don't try to cover everything initially
2. **Be specific**: Vague standards are hard to follow
3. **Explain rationale**: Help people understand "why"
4. **Allow exceptions**: But require justification
5. **Review regularly**: Technology evolves
6. **Communicate changes**: Don't surprise teams

## Anti-Patterns

- **Too restrictive**: Every decision requires override
- **Too loose**: Constitution adds no value
- **Stale**: Not reviewed in years
- **Secret**: Teams don't know it exists
- **Aspirational**: Standards nobody follows
