package structures

const (
	One = iota + 1
	Two
	Three
	Four
)

type Detective struct {
	ID         string `json:"rut" db:"rut"`
	Name       string `json:"name" db:"name"`
	Lastname   string `json:"lastname" db:"lastname"`
	Address    string `json: "address" db:"address"`
	City       string `json: "city" db:"city"`
	PostalCode string `json: "postalcode" db:"postalcode"`
	Phone      string `json: "phone" db:"phone"`
	Device     string `json: "device" db:"device"`
	Level      int    `json: "level" db:"level"`
}
