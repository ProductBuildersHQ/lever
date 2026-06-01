package lever

// ProfileID identifies a domain-specific LEVER profile.
type ProfileID string

// Profile represents a domain-specific interpretation of LEVER stages.
// Other frameworks (PBMM, ASDM) implement this to map their domain
// to the universal LEVER progression.
type Profile struct {
	// ID is the unique identifier for this profile.
	ID ProfileID `json:"id"`

	// Name is the human-readable profile name.
	Name string `json:"name"`

	// Description explains what domain this profile covers.
	Description string `json:"description"`

	// Domain categorizes the profile (e.g., "product", "engineering", "leadership").
	Domain string `json:"domain"`

	// StageInterpretations provides domain-specific meaning for each LEVER stage.
	StageInterpretations []StageInterpretation `json:"stage_interpretations"`
}

// StageInterpretation provides domain-specific meaning for a LEVER stage.
type StageInterpretation struct {
	// StageID identifies the LEVER stage being interpreted.
	StageID StageID `json:"stage_id"`

	// Title is the domain-specific title for this stage (optional).
	Title string `json:"title,omitempty"`

	// Description explains what this stage means in this domain.
	Description string `json:"description"`

	// Competencies are domain-specific capabilities at this stage.
	Competencies []Competency `json:"competencies,omitempty"`

	// Evidence describes observable indicators of this stage.
	Evidence []string `json:"evidence,omitempty"`

	// Examples provide concrete illustrations.
	Examples []string `json:"examples,omitempty"`
}

// Competency represents a specific capability within a stage.
type Competency struct {
	// ID is the unique identifier for this competency.
	ID string `json:"id"`

	// Name is the competency name.
	Name string `json:"name"`

	// Description explains the competency.
	Description string `json:"description"`

	// Indicators are observable behaviors demonstrating this competency.
	Indicators []string `json:"indicators,omitempty"`
}

// ProductBuilderProfile returns a reference Product Builder profile.
// This demonstrates how PBMM would implement LEVER.
// The actual PBMM spec would be more detailed.
func ProductBuilderProfile() Profile {
	return Profile{
		ID:          "product-builder",
		Name:        "Product Builder",
		Description: "LEVER interpretation for Product Builders - individuals who own end-to-end product outcomes.",
		Domain:      "product",
		StageInterpretations: []StageInterpretation{
			{
				StageID:     StageLearn,
				Title:       "Learner",
				Description: "Acquiring product, technology, and business knowledge.",
				Evidence: []string{
					"Completes relevant courses and certifications",
					"Studies successful products and case studies",
					"Learns technical fundamentals (coding, AI, tools)",
					"Understands business models and metrics",
				},
				Examples: []string{
					"Takes product management courses",
					"Learns to use AI coding assistants",
					"Studies customer discovery techniques",
				},
			},
			{
				StageID:     StageExecute,
				Title:       "Practitioner",
				Description: "Building prototypes and MVPs with guidance.",
				Evidence: []string{
					"Creates functional prototypes",
					"Uses AI tools to accelerate development",
					"Completes features with supervision",
					"Participates in product delivery",
				},
				Examples: []string{
					"Builds an internal tool with AI assistance",
					"Ships a feature to production",
					"Conducts user interviews",
				},
			},
			{
				StageID:     StageValue,
				Title:       "Builder",
				Description: "Independently delivering products that create business value.",
				Evidence: []string{
					"Ships products without supervision",
					"Drives measurable business outcomes",
					"Makes sound product decisions",
					"Owns the full product lifecycle",
				},
				Examples: []string{
					"Launches a product that acquires users",
					"Improves key metrics through product changes",
					"Manages a product roadmap independently",
				},
			},
			{
				StageID:     StageEnable,
				Title:       "Leader",
				Description: "Developing other Product Builders.",
				Evidence: []string{
					"Mentors junior builders",
					"Creates learning opportunities",
					"Builds high-performing teams",
					"Establishes product practices",
				},
				Examples: []string{
					"Mentors a practitioner to builder level",
					"Runs product craft sessions",
					"Defines team product processes",
				},
			},
			{
				StageID:     StageReplicate,
				Title:       "Multiplier",
				Description: "Creating systems that scale product capability.",
				Evidence: []string{
					"Creates reusable frameworks and playbooks",
					"Builds AI-powered product tools",
					"Defines organizational standards",
					"Establishes product communities",
				},
				Examples: []string{
					"Creates a product requirements AI agent",
					"Writes a product playbook adopted by teams",
					"Establishes a product guild or community",
					"Builds a platform used by other builders",
				},
			},
		},
	}
}

// EngineerProfile returns a reference Software Engineer profile.
func EngineerProfile() Profile {
	return Profile{
		ID:          "engineer",
		Name:        "Software Engineer",
		Description: "LEVER interpretation for software engineers building production systems.",
		Domain:      "engineering",
		StageInterpretations: []StageInterpretation{
			{
				StageID:     StageLearn,
				Title:       "Junior Engineer",
				Description: "Learning fundamentals of software development.",
				Evidence: []string{
					"Completes coding tutorials and courses",
					"Understands basic CS concepts",
					"Learns team tools and practices",
				},
			},
			{
				StageID:     StageExecute,
				Title:       "Engineer",
				Description: "Implementing features and fixing bugs with guidance.",
				Evidence: []string{
					"Completes assigned tasks",
					"Writes code that passes review",
					"Debugs issues with support",
				},
			},
			{
				StageID:     StageValue,
				Title:       "Senior Engineer",
				Description: "Independently delivering significant technical work.",
				Evidence: []string{
					"Owns technical decisions",
					"Delivers complex features",
					"Mentors junior engineers informally",
				},
			},
			{
				StageID:     StageEnable,
				Title:       "Staff Engineer",
				Description: "Elevating team capability and technical direction.",
				Evidence: []string{
					"Sets technical direction",
					"Mentors multiple engineers",
					"Leads technical initiatives",
				},
			},
			{
				StageID:     StageReplicate,
				Title:       "Principal/Architect",
				Description: "Creating systems that multiply engineering capability.",
				Evidence: []string{
					"Creates reusable platforms",
					"Defines engineering standards",
					"Builds tools used across teams",
					"Establishes technical communities",
				},
			},
		},
	}
}

// ReferenceProfiles returns all reference profiles.
func ReferenceProfiles() []Profile {
	return []Profile{
		ProductBuilderProfile(),
		EngineerProfile(),
	}
}
