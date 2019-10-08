package models

// GoalDefinitionTemplate A goal definition template
type GoalDefinitionTemplate struct {
	ID string `json:"id"`
}

// HealthConcernDefinitionTemplate A health concern definition template
type HealthConcernDefinitionTemplate struct {
	ID              string                   `json:"id"`
	GoalDefinitions []GoalDefinitionTemplate `json:"goalDefinitions,omitempty"`
}

// PlanTemplate A longitudinal plan template
type PlanTemplate struct {
	ID                       string                            `json:"id,omitempty"`
	Description              string                            `json:"description"`
	Status                   string                            `json:"status"`
	HealthConcernDefinitions []HealthConcernDefinitionTemplate `json:"healthConcernDefinitions,omitempty"`
	CreatedAt                string                            `json:"createdAt,omitempty"`
	UpdatedAt                string                            `json:"updatedAt,omitempty"`
}
