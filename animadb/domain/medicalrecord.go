package domain

import "time"

// MedicalRecord - A Single Medical Record For An Animal
type MedicalRecord struct {
	AuditData

	Client       Account      `json:"client,omitempty"`
	Pet          Pet          `json:"pet,omitempty"`
	Organization Organization `json:"organization,omitempty"`
	Doctor       Person       `json:"doctor,omitempty"`
	PetData      string       `json:"petData,omitempty"`
	Date         time.Time    `json:"date,omitempty"`
	Notes        string       `json:"notes,omitempty"`
	Outcome      string       `json:"outcome,omitempty"`
}
