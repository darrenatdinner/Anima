package domain

// IdentificationType - Type Of IdentificationRecord Created
type IdentificationType struct {
	AuditData

	Name string `json:"name,omitempty"`
}
