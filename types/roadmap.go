package types

// Roadmap tracks feature prioritization and status.
// Lives at specs/roadmap.json and is the source of truth for
// which features are active, in backlog, completed, or archived.
type Roadmap struct {
	// Version of the roadmap schema
	Version string `json:"version"`

	// UpdatedAt is when this roadmap was last updated (RFC3339)
	UpdatedAt string `json:"updated_at"`

	// Active features currently being worked on (ordered by priority)
	Active []ActiveFeature `json:"active"`

	// Backlog features planned but not started (ordered by priority)
	Backlog []BacklogFeature `json:"backlog"`

	// Completed features that are done
	Completed []CompletedFeature `json:"completed"`

	// Archived features that are abandoned or superseded
	Archived []ArchivedFeature `json:"archived"`
}

// ActiveFeature is a feature currently being worked on.
type ActiveFeature struct {
	// Feature is the kebab-case feature directory name
	Feature string `json:"feature"`

	// Priority within active features (1 = highest)
	Priority int `json:"priority"`

	// StartedAt is when work began (RFC3339)
	StartedAt string `json:"started_at,omitempty"`

	// Owner is the team/person responsible
	Owner string `json:"owner,omitempty"`

	// TargetVersion is the planned release version
	TargetVersion string `json:"target_version,omitempty"`

	// Notes for context
	Notes string `json:"notes,omitempty"`
}

// BacklogFeature is a planned feature not yet started.
type BacklogFeature struct {
	// Feature is the kebab-case feature directory name
	Feature string `json:"feature"`

	// Priority within backlog (1 = highest)
	Priority int `json:"priority"`

	// BlockedBy lists features that must complete first
	BlockedBy []string `json:"blocked_by,omitempty"`

	// EstimatedVersion is the tentative release version
	EstimatedVersion string `json:"estimated_version,omitempty"`

	// Notes for context
	Notes string `json:"notes,omitempty"`
}

// CompletedFeature is a feature that has been finished.
type CompletedFeature struct {
	// Feature is the kebab-case feature directory name
	Feature string `json:"feature"`

	// CompletedAt is when the feature was completed (RFC3339)
	CompletedAt string `json:"completed_at"`

	// ReleasedIn is the version that included this feature
	ReleasedIn string `json:"released_in,omitempty"`

	// Notes for context
	Notes string `json:"notes,omitempty"`
}

// ArchivedFeature is a feature that is abandoned or superseded.
type ArchivedFeature struct {
	// Feature is the kebab-case feature directory name
	Feature string `json:"feature"`

	// ArchivedAt is when the feature was archived (RFC3339)
	ArchivedAt string `json:"archived_at"`

	// Reason explains why it was archived
	Reason string `json:"reason"`

	// SupersededBy is the feature that replaced this one (if any)
	SupersededBy string `json:"superseded_by,omitempty"`
}
