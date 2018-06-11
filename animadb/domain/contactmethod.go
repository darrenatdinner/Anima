package domain

// ContactMethod - A Single Method For Contacting
type ContactMethod struct {
	AuditData

	Name string      `json:"name,omitempty"`
	Data string      `json:"data,omitempty"`
	Type ContactType `json:"type,omitempty"`
}
