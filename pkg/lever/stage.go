// Package lever provides types for the LEVER Capability Framework.
//
// LEVER (Learn, Execute, Value, Enable, Replicate) is a capability
// multiplication framework that synthesizes classical learning models
// (Bloom, Dreyfus, Shu-Ha-Ri, Apprenticeship) for the AI era.
package lever

// StageID identifies a LEVER stage.
type StageID string

const (
	// StageLearn represents knowledge acquisition.
	// Question: "What do I know?"
	// Output: Knowledge
	StageLearn StageID = "learn"

	// StageExecute represents skill application.
	// Question: "What can I do?"
	// Output: Results
	StageExecute StageID = "execute"

	// StageValue represents independent outcome creation.
	// Question: "Can I independently create outcomes?"
	// Output: Measurable outcomes
	StageValue StageID = "value"

	// StageEnable represents developing others.
	// Question: "Can I make others successful?"
	// Output: Capable people
	StageEnable StageID = "enable"

	// StageReplicate represents creating scalable systems.
	// Question: "Can I make success repeatable at scale?"
	// Output: Systems (platforms, standards, agents, institutions)
	StageReplicate StageID = "replicate"
)

// Stage represents a LEVER capability stage.
type Stage struct {
	// ID is the unique identifier for this stage.
	ID StageID `json:"id"`

	// Name is the human-readable stage name.
	Name string `json:"name"`

	// Order is the position in the LEVER progression (1-5).
	Order int `json:"order"`

	// Question is the core question this stage answers.
	Question string `json:"question"`

	// Output describes what this stage produces.
	Output string `json:"output"`

	// Description provides a detailed explanation of the stage.
	Description string `json:"description"`

	// Characteristics are observable traits of someone at this stage.
	Characteristics []string `json:"characteristics,omitempty"`
}

// Stages returns the canonical LEVER stages in order.
func Stages() []Stage {
	return []Stage{
		{
			ID:       StageLearn,
			Name:     "Learn",
			Order:    1,
			Question: "What do I know?",
			Output:   "Knowledge",
			Description: "Acquire knowledge through study, observation, and instruction. " +
				"Build vocabulary, concepts, and mental models. This is the foundation " +
				"of all capability development.",
			Characteristics: []string{
				"Studies concepts and terminology",
				"Follows instructions and tutorials",
				"Asks questions to understand",
				"Builds foundational mental models",
			},
		},
		{
			ID:       StageExecute,
			Name:     "Execute",
			Order:    2,
			Question: "What can I do?",
			Output:   "Results",
			Description: "Apply knowledge to produce work. Move from theory to practice. " +
				"Requires guidance and supervision but can complete defined tasks.",
			Characteristics: []string{
				"Completes assigned tasks",
				"Applies learned techniques",
				"Requires supervision and guidance",
				"Produces work product with support",
			},
		},
		{
			ID:       StageValue,
			Name:     "Value",
			Order:    3,
			Question: "Can I independently create outcomes?",
			Output:   "Measurable outcomes",
			Description: "Demonstrate independent ability to create value. Operate without " +
				"supervision. Make decisions and own results. This is where practitioners " +
				"become trusted contributors.",
			Characteristics: []string{
				"Works independently",
				"Makes sound decisions",
				"Delivers measurable outcomes",
				"Earns trust through results",
				"Solves problems autonomously",
			},
		},
		{
			ID:       StageEnable,
			Name:     "Enable",
			Order:    4,
			Question: "Can I make others successful?",
			Output:   "Capable people",
			Description: "Develop other people's capabilities. Mentor, coach, teach, and lead. " +
				"Output shifts from personal work to growing others. This is where " +
				"practitioners become leaders.",
			Characteristics: []string{
				"Mentors and coaches others",
				"Teaches effectively",
				"Builds team capability",
				"Multiplies impact through people",
				"Creates psychological safety",
			},
		},
		{
			ID:       StageReplicate,
			Name:     "Replicate",
			Order:    5,
			Question: "Can I make success repeatable at scale?",
			Output:   "Systems",
			Description: "Create systems that multiply capability beyond direct human interaction. " +
				"Build platforms, standards, frameworks, AI agents, automation, and institutions " +
				"that enable many others to succeed. This is the AI-era extension of classical " +
				"mastery models.",
			Characteristics: []string{
				"Creates reusable frameworks",
				"Builds platforms and tools",
				"Defines standards and best practices",
				"Develops AI agents and automation",
				"Establishes institutions and communities",
				"Scales impact through systems",
			},
		},
	}
}

// StageByID returns the Stage for the given ID, or nil if not found.
func StageByID(id StageID) *Stage {
	for _, s := range Stages() {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

// AllStageIDs returns all stage IDs in order.
func AllStageIDs() []StageID {
	return []StageID{
		StageLearn,
		StageExecute,
		StageValue,
		StageEnable,
		StageReplicate,
	}
}
