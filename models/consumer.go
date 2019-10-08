package models

// ConsumerEntity A reference to a HealtheIntent consumer
type ConsumerEntity struct {
	ID                string             `json:"id,omitempty"`
	Name              ConsumerName       `json:"name"`
	SourceIdentifiers []SourceIdentifier `json:"sourceIdentifiers,omitempty"`
}

// SourceIdentifier A representation of a source identifer.
type SourceIdentifier struct {
	ID              string `json:"id"`
	DataPartitionID string `json:"dataPartitionId"`
}

// ConsumerName A structure to store a consumer's name
type ConsumerName struct {
	Prefix    string `json:"prefix,omitempty"`
	Given     string `json:"given"`
	Middle    string `json:"middle,omitempty"`
	Family    string `json:"family"`
	Suffix    string `json:"suffix,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

// ConsumerList A list of Consumers
type ConsumerList struct {
	Items        []ConsumerEntity `json:"items"`
	FirstLink    string           `json:"firstLink"`
	LastLink     string           `json:"lastLink,omitempty"`
	PrevLink     string           `json:"previousLink,omitempty"`
	NextLink     string           `json:"nextLink,omitempty"`
	TotalResults int32            `json:"totalResults"`
}
