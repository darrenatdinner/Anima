package domain

// ContactType - Contact Type Of The Contact Methid
type ContactType struct {
	AuditData

	Name string `json:"name,omitempty"`
}
