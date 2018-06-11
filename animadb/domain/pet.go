package domain

// Pet - A Single Pet Registered With The System
type Pet struct {
	AuditData

	Name            string                 `json:"name,omitempty"`
	AnimalType      Animal                 `json:"animalType,omitempty"`
	PetMedia        []AnimalMedia          `json:"media,omitempty"`
	PetRecord       IdentificationRecord   `json:"record,omitempty"`
	ExternalRecords []IdentificationRecord `json:"externalRecords,omitempty"`
	BehaviorRecords []BehaviorRecord       `json:"behaviorRecords,omitempty"`
	MedicalRecords  []MedicalRecord        `json:"medicalRecords,omitempty"`
	StatisticOptIn  ShareRecord            `json:"statisticoptin,omitempty"`
}
