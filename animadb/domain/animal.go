package domain

// Animal - A Full Representation Of An Organism
type Animal struct {
	AuditData

	Life    string `json:"life,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Kingdom string `json:"kingdom,omitempty"`
	Phylum  string `json:"phylum,omitempty"`
	Class   string `json:"class,omitempty"`
	Order   string `json:"order,omitempty"`
	Family  string `json:"family,omitempty"`
	Genus   string `json:"genus,omitempty"`
	Species string `json:"species,omitempty"`

	CommonName string        `json:"commonName,omitempty"`
	Media      []AnimalMedia `json:"media"`
}
