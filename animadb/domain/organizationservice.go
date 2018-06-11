package domain

// OrganizationService - A Record Containing A Type Of Service
type OrganizationService struct {
	AuditData

	Name string `json:"name,omitempty"`
}
