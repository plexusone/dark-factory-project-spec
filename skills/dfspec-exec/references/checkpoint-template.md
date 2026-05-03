# Checkpoint Report Template

Use this template when generating checkpoint reports.

## Format

```markdown
# Checkpoint Report: {checkpoint_id} ({checkpoint_name})

**Generated:** {timestamp}
**Type:** {validation_type}
**Status:** {Awaiting Review | Passed | Failed}

---

## Progress Summary

| Metric | Value |
|--------|-------|
| Requirements in checkpoint | {count} |
| Requirements completed | {completed}/{total} |
| Tests passing | {pass_count} |
| Tests failing | {fail_count} |
| Test coverage | {coverage}% |
| Discoveries | {discovery_count} |

---

## Requirements Completed

{For each completed requirement:}

### {requirement_id}: {title} ✓

- **Status:** {Implemented | Validated}
- **Files:** {list of files}
- **Tests:** {passed}/{total} passing
- **Acceptance Criteria:**
  {For each criterion:}
  - [x] {criterion description}

{If any issues:}
- **Notes:** {notes}

---

## Requirements Pending

{If any requirements not yet complete:}

| ID | Title | Status | Blocker |
|----|-------|--------|---------|
| {id} | {title} | {status} | {blocker if any} |

---

## Discoveries

{If discoveries exist:}

### {discovery_id}: {title} ({severity})

- **Type:** {discovery_type}
- **Description:** {description}
- **Impact:** {impact}
- **Related Requirements:** {requirement_ids}
- **Suggested Action:** {suggested_action}
- **Status:** {Pending | Resolved}

{If resolved:}
- **Resolution:** {resolution_action} - {notes}

---

## Validation Results

{For automated checkpoints:}

| Criterion | Result |
|-----------|--------|
| {criterion 1} | ✓ Passed |
| {criterion 2} | ✗ Failed: {reason} |

---

## Test Summary

```
{test output summary}

Total: {total} | Passed: {passed} | Failed: {failed} | Skipped: {skipped}
```

{If failures:}
### Failed Tests

1. `{test_name}`: {failure reason}

---

## Next Steps

{Based on checkpoint status:}

### If Awaiting Review:

1. Review completed requirements
2. Address any discoveries
3. Choose: Approve | Amend | Reject

### If Passed:

Continuing to next requirements:
- {next_requirement_1}
- {next_requirement_2}

### If Failed:

Issues to resolve:
1. {issue_1}
2. {issue_2}

---

## Actions

{For human_review checkpoints:}

- **[Approve]** - Continue execution
- **[Amend]** - Modify spec, then continue
- **[Reject]** - Stop execution

{For automated checkpoints that failed:}

- **[Retry]** - Run validation again
- **[Override]** - Approve despite failures (requires justification)
- **[Reject]** - Stop execution
```

## Example: Human Review Checkpoint

```markdown
# Checkpoint Report: CP-001 (Core Auth)

**Generated:** 2025-01-15T14:30:00Z
**Type:** human_review
**Status:** Awaiting Review

---

## Progress Summary

| Metric | Value |
|--------|-------|
| Requirements in checkpoint | 2 |
| Requirements completed | 2/2 |
| Tests passing | 13 |
| Tests failing | 0 |
| Test coverage | 85% |
| Discoveries | 1 |

---

## Requirements Completed

### FR-001: User registration endpoint ✓

- **Status:** Validated
- **Files:** `src/auth/register.go`, `src/auth/register_test.go`
- **Tests:** 5/5 passing
- **Acceptance Criteria:**
  - [x] AC-001: Valid email and password returns 201 with user_id
  - [x] AC-002: Invalid email returns 400 with error
  - [x] AC-003: Weak password returns 400 with error

### FR-002: User login endpoint ✓

- **Status:** Validated
- **Files:** `src/auth/login.go`, `src/auth/login_test.go`
- **Tests:** 8/8 passing
- **Acceptance Criteria:**
  - [x] AC-001: Valid credentials return JWT token
  - [x] AC-002: Invalid credentials return 401
  - [x] AC-003: Account lockout after 5 failures

---

## Discoveries

### D-001: Email case sensitivity (Medium)

- **Type:** edge_case
- **Description:** Email comparison is case-sensitive, allowing User@Example.com and user@example.com as different accounts
- **Impact:** Potential duplicate accounts, login confusion
- **Related Requirements:** FR-001, FR-002
- **Suggested Action:** Normalize all emails to lowercase before storing and comparing
- **Status:** Pending

---

## Test Summary

```
=== RUN   TestRegister_ValidCredentials
--- PASS: TestRegister_ValidCredentials (0.02s)
=== RUN   TestRegister_InvalidEmail
--- PASS: TestRegister_InvalidEmail (0.01s)
...

Total: 13 | Passed: 13 | Failed: 0 | Skipped: 0
```

---

## Next Steps

1. Review FR-001 and FR-002 implementations
2. Decide on D-001 (email case sensitivity)
3. Choose action below

---

## Actions

- **[Approve]** - Continue to FR-003, FR-004, FR-005
- **[Amend]** - Add email normalization to spec, then continue
- **[Reject]** - Stop execution
```

## Example: Automated Checkpoint

```markdown
# Checkpoint Report: CP-002 (Full Auth)

**Generated:** 2025-01-15T16:00:00Z
**Type:** automated
**Status:** Passed

---

## Validation Results

| Criterion | Result |
|-----------|--------|
| All unit tests pass | ✓ Passed |
| Test coverage > 80% | ✓ Passed (87%) |
| No lint errors | ✓ Passed |
| No security vulnerabilities | ✓ Passed |

---

## Summary

All validation criteria met. Continuing execution automatically.
```
