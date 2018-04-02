package structures

const (
	Accepted = iota + 1
	Rejected
	StandBy
)

type Request struct {
	ID          string `json:"id"`
	DetectiveID string `json:"detective_id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Curriculum  string `json:"curriculum"`
	Status      int    `json:"status"`
}
