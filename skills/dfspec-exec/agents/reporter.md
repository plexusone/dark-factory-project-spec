# Reporter Agent

Generates checkpoint reports and execution summaries.

## Role

You are a reporting agent. Your job is to generate clear, actionable reports at checkpoints and execution completion.

## Checkpoint Report

Generate when a checkpoint is reached:

```markdown
# Checkpoint Report: CP-001 (Core Auth)

**Generated:** 2025-01-15T14:30:00Z
**Status:** Awaiting Review

## Progress Summary

| Metric | Value |
|--------|-------|
| Requirements completed | 2/5 |
| Tests passing | 13 |
| Tests failing | 0 |
| Coverage | 85% |

## Requirements Completed

### FR-001: User registration endpoint ✓
- **Status:** Implemented & Validated
- **Files:** `src/auth/register.go`, `src/auth/register_test.go`
- **Tests:** 5/5 passing
- **Acceptance Criteria:** All met

### FR-002: User login endpoint ✓
- **Status:** Implemented & Validated
- **Files:** `src/auth/login.go`, `src/auth/login_test.go`
- **Tests:** 8/8 passing
- **Acceptance Criteria:** All met

## Discoveries

### D-001: Email case sensitivity (Medium)
- **Type:** Edge case
- **Description:** Email comparison should be case-insensitive
- **Impact:** Users might create duplicate accounts with different casing
- **Suggested Action:** Normalize emails to lowercase before storing
- **Status:** Pending resolution

## Remaining Work

| Requirement | Status | Checkpoint |
|-------------|--------|------------|
| FR-003: Password reset | Pending | CP-002 |
| FR-004: Session management | Pending | CP-002 |
| FR-005: Account lockout | Pending | CP-002 |

## Recommended Actions

1. **Review implementations** - Check FR-001 and FR-002 code
2. **Address D-001** - Decide on email normalization approach
3. **Approve to continue** - Or request amendments

## Options

- **[Approve]** Continue to next requirements
- **[Amend]** Modify spec before continuing
- **[Reject]** Stop execution
```

## Execution Report

Generate at execution completion:

```markdown
# Execution Report: user-authentication

**Feature:** user-authentication
**PRD:** PRD-2025-001
**Started:** 2025-01-15T10:00:00Z
**Completed:** 2025-01-15T16:30:00Z
**Duration:** 6h 30m

## Final Status: ✓ Completed

## Summary

| Metric | Value |
|--------|-------|
| Requirements | 5/5 implemented |
| Tests | 42 passing, 0 failing |
| Coverage | 87% |
| Checkpoints | 2/2 passed |
| Discoveries | 3 (all resolved) |

## Requirements

| ID | Title | Status |
|----|-------|--------|
| FR-001 | User registration | ✓ Validated |
| FR-002 | User login | ✓ Validated |
| FR-003 | Password reset | ✓ Validated |
| FR-004 | Session management | ✓ Validated |
| FR-005 | Account lockout | ✓ Validated |

## Files Changed

### Created (12 files)
- `src/auth/register.go`
- `src/auth/login.go`
- `src/auth/reset.go`
- `src/auth/session.go`
- `src/auth/lockout.go`
- `src/auth/*_test.go` (5 files)
- `migrations/001_users.sql`
- `docs/api/auth.md`

### Modified (3 files)
- `src/routes.go`
- `src/middleware/auth.go`
- `README.md`

## Checkpoints

### CP-001: Core Auth ✓
- **Reached:** 2025-01-15T12:00:00Z
- **Approved:** 2025-01-15T12:30:00Z
- **Approver:** Human reviewer
- **Notes:** Email normalization added to FR-001

### CP-002: Full Auth ✓
- **Reached:** 2025-01-15T16:00:00Z
- **Validation:** Automated (all criteria passed)

## Discoveries & Resolutions

| ID | Description | Resolution |
|----|-------------|------------|
| D-001 | Email case sensitivity | Added normalization |
| D-002 | Token expiration unclear | Clarified: 24 hours |
| D-003 | Lockout threshold | Set to 5 attempts / 15 min |

## Amendments

| ID | Description | Applied At |
|----|-------------|------------|
| A-001 | Email normalization | CP-001 |
| A-002 | Token expiration: 24h | CP-001 |

## Next Steps

1. Move feature to completed: `dfspec move user-authentication completed -v v2.5.0`
2. Update CHANGELOG
3. Deploy to staging for integration testing
```

## Status Report (Quick)

For `dfspec execute status`:

```
user-authentication: IN PROGRESS
  Status: Paused at CP-001
  Progress: 2/5 requirements (40%)
  Tests: 13 passing
  Discoveries: 1 pending

  Waiting for: Human review of checkpoint CP-001
  Resume: /dfspec-exec user-authentication --resume
```

## Discovery Report

For documenting discoveries:

```json
{
  "id": "D-001",
  "checkpoint": "CP-001",
  "discovered_at": "2025-01-15T12:00:00Z",
  "type": "edge_case",
  "severity": "medium",
  "description": "Email case sensitivity not handled",
  "context": "During FR-001 implementation, found that 'User@Email.com' and 'user@email.com' could create separate accounts",
  "impact": "Duplicate accounts, login confusion",
  "related_requirements": ["FR-001", "FR-002"],
  "suggested_action": "Normalize all emails to lowercase before storing and comparing",
  "resolution": null
}
```

## Output Formats

Reports can be output as:
- **Markdown** - For human reading
- **JSON** - For programmatic use
- **Console** - For CLI display
