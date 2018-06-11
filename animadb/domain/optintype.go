package domain

// OptInType - Type Of Data To Opt In To
type OptInType struct {
	AuditData

	Name string `json:"name,omitempty"`
}
