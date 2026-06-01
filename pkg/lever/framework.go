package lever

// FrameworkID identifies a learning/mastery framework.
type FrameworkID string

const (
	// FrameworkBloom is Bloom's Taxonomy (cognitive mastery).
	FrameworkBloom FrameworkID = "bloom"

	// FrameworkDreyfus is the Dreyfus Model of Skill Acquisition.
	FrameworkDreyfus FrameworkID = "dreyfus"

	// FrameworkShuHaRi is the Japanese mastery model.
	FrameworkShuHaRi FrameworkID = "shu-ha-ri"

	// FrameworkApprenticeship is the classical apprenticeship model.
	FrameworkApprenticeship FrameworkID = "apprenticeship"

	// FrameworkConsciousCompetence is the four stages of competence model.
	FrameworkConsciousCompetence FrameworkID = "conscious-competence"
)

// FrameworkStage represents a stage within a classical framework.
type FrameworkStage struct {
	// ID is the unique identifier within the framework.
	ID string `json:"id"`

	// Name is the human-readable stage name.
	Name string `json:"name"`

	// Order is the position in the progression.
	Order int `json:"order"`

	// Description explains the stage.
	Description string `json:"description"`
}

// Framework represents a classical learning/mastery framework.
type Framework struct {
	// ID is the unique identifier for this framework.
	ID FrameworkID `json:"id"`

	// Name is the human-readable framework name.
	Name string `json:"name"`

	// Description provides context about the framework.
	Description string `json:"description"`

	// Origin describes where the framework comes from.
	Origin string `json:"origin"`

	// PrimaryFocus describes what the framework measures.
	PrimaryFocus string `json:"primary_focus"`

	// Stages are the progression stages in order.
	Stages []FrameworkStage `json:"stages"`
}

// BloomTaxonomy returns Bloom's Taxonomy framework definition.
func BloomTaxonomy() Framework {
	return Framework{
		ID:           FrameworkBloom,
		Name:         "Bloom's Taxonomy",
		Description:  "A hierarchical model of cognitive learning objectives widely used in education.",
		Origin:       "Benjamin Bloom, 1956 (revised 2001)",
		PrimaryFocus: "Cognitive mastery levels",
		Stages: []FrameworkStage{
			{ID: "remember", Name: "Remember", Order: 1, Description: "Recall facts and basic concepts"},
			{ID: "understand", Name: "Understand", Order: 2, Description: "Explain ideas or concepts"},
			{ID: "apply", Name: "Apply", Order: 3, Description: "Use information in new situations"},
			{ID: "analyze", Name: "Analyze", Order: 4, Description: "Draw connections among ideas"},
			{ID: "evaluate", Name: "Evaluate", Order: 5, Description: "Justify a decision or course of action"},
			{ID: "create", Name: "Create", Order: 6, Description: "Produce new or original work"},
		},
	}
}

// DreyfusModel returns the Dreyfus Model of Skill Acquisition.
func DreyfusModel() Framework {
	return Framework{
		ID:           FrameworkDreyfus,
		Name:         "Dreyfus Model of Skill Acquisition",
		Description:  "A model describing how students acquire skills through formal instruction and practice.",
		Origin:       "Stuart and Hubert Dreyfus, 1980",
		PrimaryFocus: "Skill acquisition and independence",
		Stages: []FrameworkStage{
			{ID: "novice", Name: "Novice", Order: 1, Description: "Follows rules exactly, requires explicit instructions"},
			{ID: "advanced-beginner", Name: "Advanced Beginner", Order: 2, Description: "Recognizes patterns from experience"},
			{ID: "competent", Name: "Competent", Order: 3, Description: "Plans and prioritizes, manages complexity"},
			{ID: "proficient", Name: "Proficient", Order: 4, Description: "Sees situations holistically, knows what matters"},
			{ID: "expert", Name: "Expert", Order: 5, Description: "Operates intuitively, transcends rules"},
		},
	}
}

// ShuHaRi returns the Shu-Ha-Ri mastery model.
func ShuHaRi() Framework {
	return Framework{
		ID:           FrameworkShuHaRi,
		Name:         "Shu-Ha-Ri",
		Description:  "A Japanese martial arts concept describing stages of learning to mastery.",
		Origin:       "Japanese martial arts tradition, popularized by Alistair Cockburn in software",
		PrimaryFocus: "Relationship to rules and tradition",
		Stages: []FrameworkStage{
			{ID: "shu", Name: "Shu (守)", Order: 1, Description: "Obey - Follow the rules exactly, learn the forms"},
			{ID: "ha", Name: "Ha (破)", Order: 2, Description: "Detach - Break from tradition, understand when rules bend"},
			{ID: "ri", Name: "Ri (離)", Order: 3, Description: "Transcend - Create your own approach, rules become instinct"},
		},
	}
}

// ApprenticeshipModel returns the classical apprenticeship model.
func ApprenticeshipModel() Framework {
	return Framework{
		ID:           FrameworkApprenticeship,
		Name:         "Apprenticeship Model",
		Description:  "The traditional craft-based progression from apprentice to master.",
		Origin:       "Medieval guild system, formalized craft traditions",
		PrimaryFocus: "Responsibility and teaching progression",
		Stages: []FrameworkStage{
			{ID: "observe", Name: "Observe", Order: 1, Description: "Watch experts perform the craft"},
			{ID: "assist", Name: "Assist", Order: 2, Description: "Help experts with supervised tasks"},
			{ID: "perform", Name: "Perform", Order: 3, Description: "Work independently on full tasks"},
			{ID: "lead", Name: "Lead", Order: 4, Description: "Guide others, take responsibility for outcomes"},
			{ID: "teach", Name: "Teach", Order: 5, Description: "Train apprentices, preserve and advance the craft"},
		},
	}
}

// ConsciousCompetenceModel returns the four stages of competence model.
func ConsciousCompetenceModel() Framework {
	return Framework{
		ID:           FrameworkConsciousCompetence,
		Name:         "Conscious Competence Model",
		Description:  "A model describing the psychological states in learning a new skill.",
		Origin:       "Noel Burch, Gordon Training International, 1970s",
		PrimaryFocus: "Awareness of own competence",
		Stages: []FrameworkStage{
			{ID: "unconscious-incompetence", Name: "Unconscious Incompetence", Order: 1, Description: "Don't know what you don't know"},
			{ID: "conscious-incompetence", Name: "Conscious Incompetence", Order: 2, Description: "Aware of skill gap, learning begins"},
			{ID: "conscious-competence", Name: "Conscious Competence", Order: 3, Description: "Can perform with deliberate effort"},
			{ID: "unconscious-competence", Name: "Unconscious Competence", Order: 4, Description: "Skill becomes automatic"},
		},
	}
}

// AllFrameworks returns all classical frameworks.
func AllFrameworks() []Framework {
	return []Framework{
		BloomTaxonomy(),
		DreyfusModel(),
		ShuHaRi(),
		ApprenticeshipModel(),
		ConsciousCompetenceModel(),
	}
}

// FrameworkByID returns the Framework for the given ID, or nil if not found.
func FrameworkByID(id FrameworkID) *Framework {
	for _, f := range AllFrameworks() {
		if f.ID == id {
			return &f
		}
	}
	return nil
}
