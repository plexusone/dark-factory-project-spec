# Design Documentation

This section contains design documentation for features.

## Structure

Each feature has its own directory:

```
docs/design/
├── {feature}/
│   ├── index.md        # Feature overview
│   ├── mrd.md          # Rendered MRD
│   └── prd.md          # Rendered PRD
```

## Generating Index Files

Use the CLI to generate index files:

```bash
# All features
dfspec generate index

# Single feature
dfspec generate index user-authentication
```

## Current Features

Feature documentation is generated from specs. Run:

```bash
dfspec generate index
```

Then feature directories will appear here with links to:

- MRD (Market Requirements Document)
- PRD (Product Requirements Document)
- Evaluation results

## Spec-to-Doc Mapping

Specs in `specs/` map to docs in `docs/design/`:

| Spec | Doc |
|------|-----|
| `specs/active/user-auth/mrd.json` | `docs/design/user-auth/mrd.md` |
| `specs/active/user-auth/prd.json` | `docs/design/user-auth/prd.md` |

The mapping strategy is configured in `constitution.json`:

```json
{
  "file_organization": {
    "spec_to_doc_mapping": {
      "strategy": "feature",
      "doc_file_extension": ".md",
      "generate_index": true
    }
  }
}
```
