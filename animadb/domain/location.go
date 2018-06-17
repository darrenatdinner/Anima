package domain

// Location - A Single Pinpoint Location
type Location struct {
	AuditData

	Address      string `json:"address,omitempty"`
	StreetLine1  string `json:"streetLine1,omitempty"`
	StreetLine2  string `json:"streetLine2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	ZipExtension string `json:"zipExtension,omitempty"`
	Region       string `json:"region,omitempty"`
	Country      string `json:"country,omitempty"`
	Continent    string `json:"continent,omitempty"`

	UseGPS     bool   `json:"useGPS,omitempty"`
	Latitude   string `json:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty"`
	Elevation  string `json:"elevation,omitempty"`
	ErrorRange string `json:"errorRange,omitempty"`
}
