package types

// ExecutionState persists the current state of PRD execution.
// This is saved to specs/{status}/{feature}/execution.json
type ExecutionState struct {
	// PRDReference links to the PRD being executed
	PRDReference string `json:"prd_reference"`

	// Feature is the feature name
	Feature string `json:"feature"`

	// StartedAt is when execution began (RFC3339)
	StartedAt string `json:"started_at"`

	// UpdatedAt is when state was last updated (RFC3339)
	UpdatedAt string `json:"updated_at"`

	// Status is the overall execution status
	Status ExecutionStatus `json:"status"`

	// CurrentRequirement is the requirement being worked on
	CurrentRequirement string `json:"current_requirement,omitempty"`

	// CurrentCheckpoint is the checkpoint we're at (if paused)
	CurrentCheckpoint string `json:"current_checkpoint,omitempty"`

	// RequirementStates tracks each requirement's execution state
	RequirementStates []RequirementState `json:"requirement_states"`

	// CheckpointStates tracks each checkpoint's state
	CheckpointStates []CheckpointState `json:"checkpoint_states"`

	// Discoveries lists all discoveries made during execution
	Discoveries []Discovery `json:"discoveries,omitempty"`

	// ExecutionLog captures key events
	ExecutionLog []ExecutionLogEntry `json:"execution_log,omitempty"`

	// PausedReason explains why execution is paused
	PausedReason string `json:"paused_reason,omitempty"`

	// ResumeInstructions tells what to do to resume
	ResumeInstructions string `json:"resume_instructions,omitempty"`
}

// RequirementState tracks execution state for a single requirement.
type RequirementState struct {
	// RequirementID is the requirement being tracked
	RequirementID string `json:"requirement_id"`

	// Status is the current state
	Status RequirementStatus `json:"status"`

	// StartedAt is when work began (RFC3339)
	StartedAt string `json:"started_at,omitempty"`

	// CompletedAt is when work finished (RFC3339)
	CompletedAt string `json:"completed_at,omitempty"`

	// BlockedReason explains why blocked
	BlockedReason string `json:"blocked_reason,omitempty"`

	// BlockedSince is when it became blocked (RFC3339)
	BlockedSince string `json:"blocked_since,omitempty"`

	// Implementation captures what was done
	Implementation *ImplementationRecord `json:"implementation,omitempty"`

	// ValidationResult captures test results
	ValidationResult *ValidationResult `json:"validation_result,omitempty"`
}

// ImplementationRecord captures what was implemented.
type ImplementationRecord struct {
	// FilesCreated lists new files
	FilesCreated []string `json:"files_created,omitempty"`

	// FilesModified lists modified files
	FilesModified []string `json:"files_modified,omitempty"`

	// TestsCreated lists test files created
	TestsCreated []string `json:"tests_created,omitempty"`

	// Summary describes what was implemented
	Summary string `json:"summary"`

	// Notes captures additional details
	Notes string `json:"notes,omitempty"`

	// CommitHash links to git commit (if applicable)
	CommitHash string `json:"commit_hash,omitempty"`
}

// ValidationResult captures test/validation results.
type ValidationResult struct {
	// Passed indicates all validations succeeded
	Passed bool `json:"passed"`

	// ValidatedAt is when validation occurred (RFC3339)
	ValidatedAt string `json:"validated_at"`

	// TestsPassed is count of passing tests
	TestsPassed int `json:"tests_passed"`

	// TestsFailed is count of failing tests
	TestsFailed int `json:"tests_failed"`

	// FailureDetails explains failures
	FailureDetails []string `json:"failure_details,omitempty"`

	// CriteriaResults maps criterion ID to pass/fail
	CriteriaResults map[string]bool `json:"criteria_results,omitempty"`
}

// CheckpointState tracks execution state for a checkpoint.
type CheckpointState struct {
	// CheckpointID is the checkpoint being tracked
	CheckpointID string `json:"checkpoint_id"`

	// Status is the current state
	Status CheckpointStatus `json:"status"`

	// ReachedAt is when checkpoint was reached (RFC3339)
	ReachedAt string `json:"reached_at,omitempty"`

	// ResolvedAt is when checkpoint was resolved (RFC3339)
	ResolvedAt string `json:"resolved_at,omitempty"`

	// Result captures validation outcome
	Result *CheckpointResult `json:"result,omitempty"`

	// WaitingFor describes what we're waiting for
	WaitingFor string `json:"waiting_for,omitempty"`
}

// ExecutionLogEntry captures a key event during execution.
type ExecutionLogEntry struct {
	// Timestamp is when this occurred (RFC3339)
	Timestamp string `json:"timestamp"`

	// Event is the event type
	Event ExecutionEvent `json:"event"`

	// Details provides context
	Details string `json:"details"`

	// RequirementID links to affected requirement
	RequirementID string `json:"requirement_id,omitempty"`

	// CheckpointID links to affected checkpoint
	CheckpointID string `json:"checkpoint_id,omitempty"`

	// DiscoveryID links to related discovery
	DiscoveryID string `json:"discovery_id,omitempty"`
}

// ExecutionEvent categorizes log entries.
type ExecutionEvent string

const (
	// EventStarted means execution began
	EventStarted ExecutionEvent = "started"

	// EventRequirementStarted means work on requirement began
	EventRequirementStarted ExecutionEvent = "requirement_started"

	// EventRequirementCompleted means requirement implemented
	EventRequirementCompleted ExecutionEvent = "requirement_completed"

	// EventRequirementBlocked means requirement blocked
	EventRequirementBlocked ExecutionEvent = "requirement_blocked"

	// EventCheckpointReached means checkpoint reached
	EventCheckpointReached ExecutionEvent = "checkpoint_reached"

	// EventCheckpointPassed means checkpoint passed
	EventCheckpointPassed ExecutionEvent = "checkpoint_passed"

	// EventCheckpointFailed means checkpoint failed
	EventCheckpointFailed ExecutionEvent = "checkpoint_failed"

	// EventDiscovery means discovery made
	EventDiscovery ExecutionEvent = "discovery"

	// EventPaused means execution paused
	EventPaused ExecutionEvent = "paused"

	// EventResumed means execution resumed
	EventResumed ExecutionEvent = "resumed"

	// EventCompleted means all requirements done
	EventCompleted ExecutionEvent = "completed"

	// EventFailed means execution failed
	EventFailed ExecutionEvent = "failed"

	// EventAmendment means spec was amended
	EventAmendment ExecutionEvent = "amendment"
)

// ExecutionCommand represents a command to the executor.
type ExecutionCommand struct {
	// Command is the action to take
	Command CommandType `json:"command"`

	// CheckpointID is for checkpoint-related commands
	CheckpointID string `json:"checkpoint_id,omitempty"`

	// DiscoveryID is for discovery-related commands
	DiscoveryID string `json:"discovery_id,omitempty"`

	// Notes provides context for the command
	Notes string `json:"notes,omitempty"`

	// Amendments lists spec amendments to apply
	Amendments []Amendment `json:"amendments,omitempty"`
}

// CommandType is the type of command.
type CommandType string

const (
	// CommandStart begins execution
	CommandStart CommandType = "start"

	// CommandResume resumes paused execution
	CommandResume CommandType = "resume"

	// CommandPause pauses execution
	CommandPause CommandType = "pause"

	// CommandApproveCheckpoint approves a checkpoint
	CommandApproveCheckpoint CommandType = "approve_checkpoint"

	// CommandRejectCheckpoint rejects a checkpoint
	CommandRejectCheckpoint CommandType = "reject_checkpoint"

	// CommandResolveDiscovery resolves a discovery
	CommandResolveDiscovery CommandType = "resolve_discovery"

	// CommandAmend amends the spec
	CommandAmend CommandType = "amend"

	// CommandAbort aborts execution
	CommandAbort CommandType = "abort"
)
