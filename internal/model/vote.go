package model

type Poll struct {
	ID        string            `json:"id"`
	CreatorID string            `json:"creator_id"`
	Question  string            `json:"question"`
	Options   []string          `json:"options"`
	Votes     map[string]string `json:"votes"` 
	IsClosed  bool              `json:"is_closed"`
}
