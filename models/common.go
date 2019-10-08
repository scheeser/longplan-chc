package models

// CodeableConcept HealtheIntent codeable concept
type CodeableConcept struct {
	Text    string   `json:"text"`
	Codings []Coding `json:"codings,omitempty"`
}

// Coding A standard coding
type Coding struct {
	Code    string `json:"code"`
	Display string `json:"display"`
	System  string `json:"system,omitempty"`
}

// Concept A reference to an ontology concept
type Concept struct {
	ContextID string `json:"contextId"`
	Alias     string `json:"alias"`
}

// Category A longitudinal plan resource category.
type Category struct {
	ID      string   `json:"id,omitempty"`
	Text    string   `json:"text"`
	Status  string   `json:"status"`
	Coding  []Coding `json:"coding,omitempty"`
	Version int32    `json:"version"`
}

// Discipline A longitudinal plan resource discipline.
type Discipline struct {
	ID      string   `json:"id,omitempty"`
	Text    string   `json:"text"`
	Status  string   `json:"status"`
	Coding  []Coding `json:"coding,omitempty"`
	Version int32    `json:"version"`
}

// Reference A reference to an identifier
type Reference struct {
	ID string `json:"id,omitempty"`
}

// CreatedBy A struct to hold created by details.
type CreatedBy struct {
	Type      string    `json:"type,omitempty"`
	Display   string    `json:"display,omitempty"`
	Reference Reference `json:"reference,omitempty"`
}
