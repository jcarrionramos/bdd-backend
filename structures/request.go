package structures

const (
	Accepted = iota + 1
	Rejected
	StandBy
)

type Request struct {
	ID          string `json:"id,omitempty"`
	DetectiveID string `json:"detective_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Info        string `json:"info,omitempty"`
	Curriculum  string `json:"curriculum,omitempty"`
	Status      int    `json:"status,omitempty"`
}
