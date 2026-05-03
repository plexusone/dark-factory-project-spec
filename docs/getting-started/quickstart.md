# Quick Start

This guide walks through creating your first feature spec.

## 1. Initialize a Feature

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

## 2. Edit the MRD

Open `specs/backlog/user-authentication/mrd.json` and fill in the template:

```json
{
  "metadata": {
    "id": "MRD-2025-001",
    "title": "User Authentication",
    "feature": "user-authentication",
    "version": "1.0.0",
    "status": "draft",
    "authors": ["Your Name"],
    "created_at": "2025-01-15T00:00:00Z",
    "updated_at": "2025-01-15T00:00:00Z"
  },
  "problem_statement": {
    "summary": "Users cannot securely access protected resources",
    "current_state": "No authentication system exists",
    "desired_state": "Users can register, login, and access protected resources",
    "impact": "Blocks all features requiring user identity"
  },
  "target_users": [
    {
      "name": "End User",
      "description": "Someone who wants to use protected features",
      "needs": ["Secure login", "Password recovery", "Session management"],
      "priority": "primary"
    }
  ],
  "success_metrics": [
    {
      "id": "SM-001",
      "name": "Login Success Rate",
      "description": "Percentage of login attempts that succeed",
      "target_value": ">99%",
      "measurement_method": "Analytics tracking"
    }
  ],
  "high_level_requirements": [
    {
      "id": "HLR-001",
      "description": "Users can register with email and password",
      "rationale": "Foundation for user identity",
      "priority": "must_have"
    },
    {
      "id": "HLR-002",
      "description": "Users can log in with credentials",
      "rationale": "Access to protected resources",
      "priority": "must_have"
    }
  ],
  "non_goals": [
    "Social login (OAuth) - future version",
    "Multi-factor authentication - future version"
  ]
}
```

## 3. Validate the Spec

```bash
dfspec validate user-authentication
```

Output:

```
✓ All specs valid
```

## 4. Move to Active

When ready to implement:

```bash
dfspec move user-authentication active
```

Output:

```
Moved 'user-authentication' from backlog/ to active/

Feature is now active. Next steps:
  1. Create/update PRD: specs/active/user-authentication/prd.json
  2. Implement the feature
  3. Run 'dfspec move user-authentication completed' when done
```

## 5. Create the PRD

Create `specs/active/user-authentication/prd.json` with detailed requirements:

```json
{
  "metadata": {
    "id": "PRD-2025-001",
    "title": "User Authentication - Technical Spec",
    "feature": "user-authentication",
    "version": "1.0.0",
    "status": "draft",
    "authors": ["Your Name"],
    "created_at": "2025-01-15T00:00:00Z",
    "updated_at": "2025-01-15T00:00:00Z"
  },
  "mrd_reference": "MRD-2025-001",
  "functional_requirements": [
    {
      "id": "FR-001",
      "mrd_requirement": "HLR-001",
      "description": "User registration endpoint",
      "acceptance_criteria": [
        "POST /api/auth/register accepts email and password",
        "Password must be at least 8 characters",
        "Email must be unique",
        "Returns user ID on success",
        "Returns 400 with validation errors on failure"
      ],
      "uncertainty": "low",
      "status": "pending"
    },
    {
      "id": "FR-002",
      "mrd_requirement": "HLR-002",
      "description": "User login endpoint",
      "acceptance_criteria": [
        "POST /api/auth/login accepts email and password",
        "Returns JWT token on success",
        "Returns 401 on invalid credentials",
        "Locks account after 5 failed attempts"
      ],
      "uncertainty": "low",
      "status": "pending",
      "checkpoint": "CP-001"
    }
  ],
  "checkpoints": [
    {
      "id": "CP-001",
      "name": "Core Auth Complete",
      "after_requirements": ["FR-001", "FR-002"],
      "validation_type": "human_review",
      "pause_on_discovery": true,
      "description": "Review auth implementation before proceeding"
    }
  ]
}
```

## 6. View the Roadmap

```bash
dfspec list
```

Output:

```
Active (1):
  • user-authentication
    Priority: 1

Backlog (0):
  (none)

Completed (0):
  (none)
```

## 7. Generate Documentation

```bash
dfspec generate roadmap
dfspec generate index user-authentication
```

This creates:

- `ROADMAP.md` - Human-readable roadmap
- `docs/design/user-authentication/index.md` - Feature documentation index

## Next Steps

- Learn about [Project Structure](structure.md)
- Understand [MRD concepts](../concepts/mrd.md)
- Explore [PRD concepts](../concepts/prd.md)
- Use the [dfspec-guide skill](../skill/dfspec-guide.md) for interactive spec creation
