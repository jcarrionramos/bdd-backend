package structures

type Response struct {
	Status int         `json:"status,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
