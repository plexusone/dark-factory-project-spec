# Constitution Schema

Organization-wide standards schema.

## Go Type

```go
type Constitution struct {
    Metadata         ConstitutionMetadata       `json:"metadata"`
    TechStack        TechStack                  `json:"tech_stack"`
    Architecture     ArchitectureStandards      `json:"architecture"`
    FileOrganization FileOrganizationStandards  `json:"file_organization"`
    Quality          QualityStandards           `json:"quality"`
    Security         SecurityStandards          `json:"security"`
    Documentation    DocumentationStandards     `json:"documentation"`
    Process          ProcessStandards           `json:"process"`
}
```

## Fields

### metadata

Constitution versioning.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | string | Yes | Semantic version |
| `effective_date` | string | Yes | RFC3339 timestamp |
| `approved_by` | string[] | Yes | Who approved |
| `last_reviewed` | string | Yes | RFC3339 timestamp |
| `next_review` | string | No | RFC3339 timestamp |

### tech_stack

Approved technologies.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `languages` | ApprovedTechnology[] | Yes | Programming languages |
| `frameworks` | ApprovedTechnology[] | Yes | Frameworks |
| `databases` | ApprovedTechnology[] | Yes | Databases |
| `infrastructure` | ApprovedTechnology[] | Yes | Infrastructure/cloud |
| `libraries` | ApprovedTechnology[] | No | Required libraries |
| `prohibited` | ProhibitedTechnology[] | No | Banned technologies |

#### ApprovedTechnology

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Technology name |
| `category` | string | No | Category grouping |
| `versions` | string | No | Allowed versions |
| `use_case` | string | No | When to use |
| `rationale` | string | No | Why approved |
| `status` | enum | Yes | `preferred`, `approved`, `legacy`, `experimental`, `deprecated` |
| `owner` | string | No | Responsible team |

### architecture

Architectural standards.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `patterns` | ArchitecturePattern[] | Yes | Approved patterns |
| `principles` | Principle[] | Yes | Guiding principles |
| `api_standards` | APIStandards | No | API design |
| `data_standards` | DataStandards | No | Data handling |

### file_organization

Project structure standards.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `specs` | SpecsDirectoryStandards | Yes | Spec file locations |
| `docs` | DocsDirectoryStandards | Yes | Doc file locations |
| `spec_to_doc_mapping` | SpecToDocMapping | Yes | How specs map to docs |
| `naming_conventions` | NamingConventions | Yes | Naming rules |

### quality

Quality requirements.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `testing` | TestingStandards | Yes | Test requirements |
| `code_review` | CodeReviewStandards | Yes | Review requirements |
| `metrics` | QualityMetric[] | No | Quality metrics |

### security

Security requirements.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `authentication` | AuthenticationStandards | Yes | Auth requirements |
| `authorization` | AuthorizationStandards | Yes | Authz requirements |
| `data_protection` | DataProtectionStandards | Yes | Data protection |
| `compliance` | string[] | No | Compliance requirements |
| `vulnerability_scanning` | VulnerabilityScanningStandards | Yes | Scanning requirements |

### process

Development process.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `branching_strategy` | string | Yes | Git branching |
| `commit_convention` | string | Yes | Commit format |
| `release_process` | string | Yes | Release process |
| `incident_response` | string | No | Incident handling |
| `on_call_required` | boolean | Yes | On-call needed? |

## Example

See [example-constitution.json](https://github.com/plexusone/dark-factory-project-spec/blob/main/examples/example-constitution.json).
