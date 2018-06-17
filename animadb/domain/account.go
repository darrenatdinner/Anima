package domain

// Account - Single Account For Associating Contacts
type Account struct {
	AuditData

	AccountIdentifier   string        `json:"accountIdentifier,omitempty"`
	Pets                []Pet         `json:"pets,omitempty"`
	PrimaryContact      Person        `json:"primaryContact,omitempty"`
	Contacts            []Person      `json:"contacts,omitempty"`
	Animals             []Pet         `json:"animals,omitempty"`
	DataShare           []ShareRecord `json:"datashare,omitempty"`
	LocalDatastoreOptin ShareRecord   `json:"localDatastoreOptin,omitempty"`
}
