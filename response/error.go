package response

type ErrorResponse struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details"`
}
