package domain

import "time"

// IdentificationRecord -
type IdentificationRecord struct {
	AuditData

	RecordNumber   string             `json:"recordNumber,omitempty"`
	Visible        bool               `json:"visible,omitempty"`
	Provider       Organization       `json:"provider,omitempty"`
	RegisteredDate time.Time          `json:"registeredDate,omitempty"`
	Type           IdentificationType `json:"type,omitempty"`
}
