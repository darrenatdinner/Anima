package domain

import (
	"time"
)

// AuditData - Audit Data Attached To The Perdidted Data Object
type AuditData struct {
	ID int `json:"id,omitempty"`

	CreatedBy  string    `json:"createdBy,omitempty"`
	CreateOn   time.Time `json:"createdOn,omitempty"`
	ModifiedBy string    `json:"modifiedBy,omitempty"`
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	Version    string    `json:"version,omitempty"`
}

// StampData - Stamp The Current Audit Date To Current Time
func (audit AuditData) StampData(userStamp string) {

	if audit.ID == -1 {
		audit.CreatedBy = userStamp
		audit.CreateOn = time.Now()
	}
	audit.ModifiedBy = userStamp
	audit.ModifiedOn = time.Now()
}
