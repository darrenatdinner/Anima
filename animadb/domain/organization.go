package domain

// Organization - Data For String Information About An Organization
type Organization struct {
	AuditData

	Name            string                `json:"name,omitempty"`
	PrimaryLocation Location              `json:"primaryLocation,omitempty"`
	PrimaryContact  ContactMethod         `json:"primaryContact,omitempty"`
	Employees       []Person              `json:"employees,omitempty"`
	Services        []OrganizationService `json:"services,omitempty"`
}
