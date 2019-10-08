package models

// HealthConcern A consumer's longitudinal plan health concern
type HealthConcern struct {
	ID                      string          `json:"id,omitempty"`
	Consumer                Reference       `json:"consumer,omitempty"`
	Onset                   string          `json:"onset,omitempty"`
	Abatement               string          `json:"abatement,omitempty"`
	Status                  string          `json:"status"` // TODO: Improve how ennumerations are represented
	Code                    CodeableConcept `json:"code"`
	CreatedBy               CreatedBy       `json:"createdBy"`
	HealthConcernDefinition Reference       `json:"healthConcernDefinition,omitempty"`
	CreatedAt               string          `json:"createdAt,omitempty"`
	UpdatedAt               string          `json:"updatedAt,omitempty"`
	Version                 int32           `json:"version,omitempty"`
}

// HealthConcernDefinition A longitudinal plan health concern definition
type HealthConcernDefinition struct {
	ID        string  `json:"id,omitempty"`
	Text      string  `json:"text"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt,omitempty"`
	UpdatedAt string  `json:"updatedAt,omitempty"`
	Concept   Concept `json:"concept,omitempty"`
}
