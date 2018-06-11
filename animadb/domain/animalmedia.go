package domain

// AnimalMedia - A File Such As A Picture Or Video Of The Animal
type AnimalMedia struct {
	AuditData

	FileName     string `json:"fileName,omitempty"`
	FileLocation string `json:"fileLocation,omitempty"`
	MediaType    string `json:"mediaType,omitempty"`
	Metadata     string `json:"metadata,omitempty"`
}
