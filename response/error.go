package response

type Error struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details"`
}
