package structures

type Detective struct {
	ID         string `json:"rut,omitempty"`
	FirstName  string `json:"name,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postalcode,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Device     string `json:"device,omitempty"`
	Position   int    `json:"position,omitempty"`
}
