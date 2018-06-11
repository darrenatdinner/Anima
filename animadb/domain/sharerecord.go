package domain

// ShareRecord - A Record Containing The Opt-in Staus
type ShareRecord struct {
	AuditData

	Person    Person    `json:"person,omitempty"`
	OptInType OptInType `json:"optInType,omitempty"`
	Data      string    `json:"data,omitempty"`
	OptIn     bool      `json:"optIn,omitempty"`
}
