package types

// Checkpoint defines a validation point during execution.
// The factory pauses at checkpoints to validate progress and handle discoveries.
type Checkpoint struct {
	// ID uniquely identifies this checkpoint
	ID string `json:"id"`

	// Name is a human-readable name
	Name string `json:"name"`

	// Description explains what is validated at this checkpoint
	Description string `json:"description"`

	// AfterRequirements lists requirement IDs to complete before this checkpoint
	AfterRequirements []string `json:"after_requirements"`

	// Validation specifies how to validate
	Validation ValidationType `json:"validation"`

	// PauseOnDiscovery stops if new requirements emerge
	PauseOnDiscovery bool `json:"pause_on_discovery"`

	// ValidationCriteria describes what to check
	ValidationCriteria []string `json:"validation_criteria,omitempty"`

	// Status tracks checkpoint completion
	Status CheckpointStatus `json:"status"`

	// Result captures the validation outcome
	Result *CheckpointResult `json:"result,omitempty"`
}

// ValidationType specifies how to validate at a checkpoint.
type ValidationType string

const (
	// ValidationHumanReview requires human approval
	ValidationHumanReview ValidationType = "human_review"

	// ValidationAutomated uses automated checks
	ValidationAutomated ValidationType = "automated"

	// ValidationSkip skips validation
	ValidationSkip ValidationType = "skip"
)

// CheckpointStatus tracks checkpoint completion.
type CheckpointStatus string

const (
	// CheckpointPending means not reached yet
	CheckpointPending CheckpointStatus = "pending"

	// CheckpointReached means checkpoint reached, awaiting validation
	CheckpointReached CheckpointStatus = "reached"

	// CheckpointPassed means validation succeeded
	CheckpointPassed CheckpointStatus = "passed"

	// CheckpointFailed means validation failed
	CheckpointFailed CheckpointStatus = "failed"

	// CheckpointBlocked means blocked by discoveries
	CheckpointBlocked CheckpointStatus = "blocked"
)

// CheckpointResult captures the validation outcome.
type CheckpointResult struct {
	// ValidatedAt is when validation occurred (RFC3339)
	ValidatedAt string `json:"validated_at"`

	// ValidatedBy is who or what performed validation
	ValidatedBy string `json:"validated_by"`

	// Passed indicates whether validation succeeded
	Passed bool `json:"passed"`

	// Notes capture validation observations
	Notes string `json:"notes,omitempty"`

	// Discoveries lists new requirements found during validation
	Discoveries []Discovery `json:"discoveries,omitempty"`

	// BlockingIssues lists issues that block progress
	BlockingIssues []string `json:"blocking_issues,omitempty"`
}

// Discovery represents a new requirement or finding discovered during execution.
type Discovery struct {
	// ID uniquely identifies this discovery
	ID string `json:"id"`

	// Type categorizes the discovery
	Type DiscoveryType `json:"type"`

	// Description explains what was discovered
	Description string `json:"description"`

	// Context provides background
	Context string `json:"context,omitempty"`

	// Impact describes how this affects the project
	Impact string `json:"impact,omitempty"`

	// Severity indicates how critical this is
	Severity DiscoverySeverity `json:"severity"`

	// RelatedRequirements lists affected requirement IDs
	RelatedRequirements []string `json:"related_requirements,omitempty"`

	// SuggestedAction describes what to do
	SuggestedAction string `json:"suggested_action,omitempty"`

	// Resolution tracks how this was handled
	Resolution *DiscoveryResolution `json:"resolution,omitempty"`
}

// DiscoveryType categorizes discoveries.
type DiscoveryType string

const (
	// DiscoveryNewRequirement is a new requirement
	DiscoveryNewRequirement DiscoveryType = "new_requirement"

	// DiscoveryAmbiguity is unclear specification
	DiscoveryAmbiguity DiscoveryType = "ambiguity"

	// DiscoveryConflict is conflicting requirements
	DiscoveryConflict DiscoveryType = "conflict"

	// DiscoveryMissingInfo is missing information
	DiscoveryMissingInfo DiscoveryType = "missing_info"

	// DiscoveryTechnicalIssue is a technical problem
	DiscoveryTechnicalIssue DiscoveryType = "technical_issue"

	// DiscoveryEdgeCase is an unhandled edge case
	DiscoveryEdgeCase DiscoveryType = "edge_case"
)

// DiscoverySeverity indicates how critical a discovery is.
type DiscoverySeverity string

const (
	// SeverityCritical blocks all progress
	SeverityCritical DiscoverySeverity = "critical"

	// SeverityHigh blocks the current requirement
	SeverityHigh DiscoverySeverity = "high"

	// SeverityMedium requires attention but doesn't block
	SeverityMedium DiscoverySeverity = "medium"

	// SeverityLow can be addressed later
	SeverityLow DiscoverySeverity = "low"
)

// DiscoveryResolution tracks how a discovery was handled.
type DiscoveryResolution struct {
	// ResolvedAt is when resolution occurred (RFC3339)
	ResolvedAt string `json:"resolved_at"`

	// ResolvedBy is who resolved it
	ResolvedBy string `json:"resolved_by"`

	// Action describes what was done
	Action ResolutionAction `json:"action"`

	// Notes provide additional context
	Notes string `json:"notes,omitempty"`

	// AmendmentID links to any spec amendment
	AmendmentID string `json:"amendment_id,omitempty"`
}

// ResolutionAction describes how a discovery was resolved.
type ResolutionAction string

const (
	// ResolutionAddedToSpec means added as new requirement
	ResolutionAddedToSpec ResolutionAction = "added_to_spec"

	// ResolutionDeferred means postponed to future version
	ResolutionDeferred ResolutionAction = "deferred"

	// ResolutionRejected means intentionally not addressed
	ResolutionRejected ResolutionAction = "rejected"

	// ResolutionClarified means clarified existing requirement
	ResolutionClarified ResolutionAction = "clarified"

	// ResolutionWorkaround means implemented a workaround
	ResolutionWorkaround ResolutionAction = "workaround"
)

// ExecutionReport captures the current state of execution.
type ExecutionReport struct {
	// PRDReference links to the PRD being executed
	PRDReference string `json:"prd_reference"`

	// GeneratedAt is when this report was generated (RFC3339)
	GeneratedAt string `json:"generated_at"`

	// OverallStatus summarizes execution state
	OverallStatus ExecutionStatus `json:"overall_status"`

	// Completion breaks down progress
	Completion CompletionSummary `json:"completion"`

	// CurrentCheckpoint is the active checkpoint (if any)
	CurrentCheckpoint string `json:"current_checkpoint,omitempty"`

	// BlockedItems lists blocked requirements
	BlockedItems []BlockedItem `json:"blocked_items,omitempty"`

	// Discoveries lists all discoveries during execution
	Discoveries []Discovery `json:"discoveries,omitempty"`

	// NextActions suggests what to do next
	NextActions []string `json:"next_actions,omitempty"`
}

// ExecutionStatus summarizes the overall execution state.
type ExecutionStatus string

const (
	// ExecutionNotStarted means execution hasn't begun
	ExecutionNotStarted ExecutionStatus = "not_started"

	// ExecutionInProgress means actively executing
	ExecutionInProgress ExecutionStatus = "in_progress"

	// ExecutionPaused means waiting at checkpoint
	ExecutionPaused ExecutionStatus = "paused"

	// ExecutionBlocked means cannot continue
	ExecutionBlocked ExecutionStatus = "blocked"

	// ExecutionCompleted means all requirements done
	ExecutionCompleted ExecutionStatus = "completed"

	// ExecutionFailed means execution failed
	ExecutionFailed ExecutionStatus = "failed"
)

// CompletionSummary breaks down progress by status.
type CompletionSummary struct {
	// Pending is count of pending requirements
	Pending int `json:"pending"`

	// InProgress is count of in-progress requirements
	InProgress int `json:"in_progress"`

	// Blocked is count of blocked requirements
	Blocked int `json:"blocked"`

	// Implemented is count of implemented requirements
	Implemented int `json:"implemented"`

	// Validated is count of validated requirements
	Validated int `json:"validated"`

	// Total is total requirement count
	Total int `json:"total"`

	// PercentComplete is percentage done
	PercentComplete float64 `json:"percent_complete"`
}

// BlockedItem describes a blocked requirement.
type BlockedItem struct {
	// RequirementID is the blocked requirement
	RequirementID string `json:"requirement_id"`

	// Reason explains why it's blocked
	Reason string `json:"reason"`

	// BlockedSince is when it became blocked (RFC3339)
	BlockedSince string `json:"blocked_since"`

	// DependsOn lists blocking dependencies
	DependsOn []string `json:"depends_on,omitempty"`
}
