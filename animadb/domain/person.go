package domain

import "time"

// Person - A Single Person
type Person struct {
	AuditData

	GivenName      string               `json:"givenName,omitempty"`
	MiddleName     string               `json:"middleName,omitempty"`
	FamilyName     string               `json:"familyName,omitempty"`
	Title          string               `json:"title,omitempty"`
	BirthDate      time.Time            `json:"birthDate,omitempty"`
	Identification IdentificationRecord `json:"identification,omitempty"`
	Contact        ContactMethod        `json:"contact,omitempty"`
}
