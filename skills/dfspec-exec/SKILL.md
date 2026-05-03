# Execute Spec

Execute an approved PRD specification autonomously with checkpoint-based human oversight.

## Trigger

Use this skill when:
- User says "execute spec", "implement spec", "run the spec"
- User invokes `/dfspec-exec`
- User wants to implement an approved PRD

## Prerequisites

Before execution can begin:
1. Feature must exist in `specs/active/{feature}/`
2. PRD must exist: `specs/active/{feature}/prd.json`
3. PRD status must be "approved"
4. No blocking evaluation findings

## Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│  1. LOAD SPEC                                                   │
│     - Read PRD from specs/active/{feature}/prd.json             │
│     - Read or create execution state                            │
│     - Validate PRD is approved                                  │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│  2. PLAN EXECUTION                                              │
│     - Order requirements by dependencies                        │
│     - Identify checkpoints                                      │
│     - Check for high-uncertainty items                          │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│  3. EXECUTE REQUIREMENTS (loop)                                 │
│     For each requirement in order:                              │
│     a. Check if checkpoint reached → pause if yes               │
│     b. Check for blockers → mark blocked if any                 │
│     c. Implement the requirement                                │
│     d. Run tests/validation                                     │
│     e. Update execution state                                   │
│     f. Check for discoveries → record and assess severity       │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│  4. CHECKPOINT HANDLING (when reached)                          │
│     - Generate checkpoint report                                │
│     - If validation_type == "human_review": pause and wait      │
│     - If validation_type == "automated": run checks             │
│     - If pause_on_discovery && discoveries: pause               │
│     - Present findings to user                                  │
│     - Wait for approval/amendment/rejection                     │
└─────────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────────┐
│  5. COMPLETION                                                  │
│     - Generate final execution report                           │
│     - Update requirement statuses in PRD                        │
│     - Summarize what was implemented                            │
└─────────────────────────────────────────────────────────────────┘
```

## Usage

### Start Execution

```
/dfspec-exec user-authentication
```

### Resume Paused Execution

```
/dfspec-exec user-authentication --resume
```

### Check Status

```
/dfspec-exec user-authentication --status
```

## Execution State

State is persisted to `specs/active/{feature}/execution.json`:

```json
{
  "prd_reference": "PRD-2025-001",
  "feature": "user-authentication",
  "started_at": "2025-01-15T10:00:00Z",
  "status": "paused",
  "current_checkpoint": "CP-001",
  "requirement_states": [
    {"requirement_id": "FR-001", "status": "implemented"},
    {"requirement_id": "FR-002", "status": "in_progress"}
  ],
  "checkpoint_states": [
    {"checkpoint_id": "CP-001", "status": "reached"}
  ]
}
```

## Checkpoint Behavior

### human_review Checkpoints

When a `human_review` checkpoint is reached:

1. **Pause execution** - Stop implementing new requirements
2. **Generate report** - Show progress, discoveries, test results
3. **Present to user** - Display checkpoint summary
4. **Wait for response**:
   - **Approve**: Continue to next requirements
   - **Amend**: Apply spec changes, then continue
   - **Reject**: Stop execution, mark checkpoint failed

### automated Checkpoints

When an `automated` checkpoint is reached:

1. **Run validation criteria** - Execute specified checks
2. **If all pass**: Mark passed, continue
3. **If any fail**: Pause for human review

### Discovery Handling

When `pause_on_discovery` is true:

1. **Record discovery** - Log what was found
2. **Assess severity**:
   - Critical/High: Pause immediately
   - Medium: Continue but flag for checkpoint review
   - Low: Continue, note in report
3. **If pausing**: Present discovery to user for resolution

## Implementing Requirements

For each requirement:

### 1. Pre-Implementation

```
- Read requirement details
- Check acceptance criteria
- Review test hints
- Check uncertainty level
  - If high: pause and ask for clarification
- Identify dependencies
```

### 2. Implementation

```
- Create/modify files as needed
- Follow coding standards from constitution
- Add inline comments for complex logic
- Create tests based on test hints
```

### 3. Post-Implementation

```
- Run tests
- Validate acceptance criteria
- Update requirement status
- Record implementation details (files changed, etc.)
- Check for discoveries
```

## Example Session

```
User: /dfspec-exec user-authentication

Claude: Loading PRD for user-authentication...

PRD: User Authentication (PRD-2025-001)
Status: approved
Requirements: 5 functional, 2 non-functional
Checkpoints: 2

Execution Plan:
  1. FR-001: User registration endpoint
  2. FR-002: User login endpoint
  → CP-001: Core Auth (human_review)
  3. FR-003: Password reset
  4. FR-004: Session management
  5. FR-005: Account lockout
  → CP-002: Full Auth (automated)

Starting execution...

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Implementing FR-001: User registration endpoint
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

[Implementation details...]

✓ FR-001 implemented
  Files: src/auth/register.go, src/auth/register_test.go
  Tests: 5 passed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Implementing FR-002: User login endpoint
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

[Implementation details...]

✓ FR-002 implemented
  Files: src/auth/login.go, src/auth/login_test.go
  Tests: 8 passed

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
CHECKPOINT CP-001: Core Auth
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Validation: human_review
Requirements completed: FR-001, FR-002

Summary:
  ✓ Registration endpoint working
  ✓ Login endpoint working
  ✓ JWT token generation
  ✓ 13 tests passing

Discoveries: 1
  ⚠ D-001 (medium): Email case sensitivity
    Found that email comparison should be case-insensitive.
    Suggested: Normalize to lowercase before storing.

Awaiting your review. Options:
  1. [Approve] - Continue to next requirements
  2. [Amend] - Modify spec and continue
  3. [Reject] - Stop execution

User: Approve, and add the email normalization to FR-001

Claude: ✓ Checkpoint CP-001 approved

Amending FR-001 to include email normalization...
Amendment A-001 added to PRD.

Continuing execution...
```

## Handling Edge Cases

### High Uncertainty Requirements

When encountering a requirement with `uncertainty: high`:

```
Claude: ⚠ FR-003 has high uncertainty

Uncertainty: The password reset flow is underspecified.
Reason: "Token expiration and delivery method not specified"
Discovery prompt: "How should reset tokens be delivered? Email? SMS?
                   What should the expiration be?"

I need clarification before implementing. Options:
  1. Provide details now
  2. Skip and mark blocked
  3. Make reasonable assumptions (will flag for review)
```

### Blocked Requirements

When a requirement has unmet dependencies:

```
Claude: ⚡ FR-004 is blocked

Depends on: FR-003 (Password reset)
FR-003 status: blocked

Skipping FR-004. Will attempt after dependencies resolve.
```

### Critical Discovery

When a critical discovery occurs:

```
Claude: 🛑 CRITICAL DISCOVERY

D-002: Database schema conflict
  The user table already exists with incompatible columns.
  This affects: FR-001, FR-002

Execution paused. This must be resolved before continuing.

Options:
  1. Provide resolution (migration strategy, etc.)
  2. Amend spec to work with existing schema
  3. Abort execution
```

## Files

```
skills/dfspec-exec/
├── SKILL.md                    # This file
├── agents/
│   ├── implementer.md          # Implements individual requirements
│   ├── validator.md            # Runs tests and validations
│   └── reporter.md             # Generates checkpoint reports
└── references/
    ├── implementation-guide.md # How to implement requirements
    └── checkpoint-template.md  # Checkpoint report format
```

## Integration with CLI

The `/dfspec-exec` skill works with the `dfspec` CLI:

```bash
# Check execution status
dfspec execute status user-authentication

# List pending checkpoints
dfspec checkpoint list user-authentication

# Approve a checkpoint (can also be done in skill)
dfspec checkpoint approve user-authentication CP-001

# View execution log
dfspec execute log user-authentication
```