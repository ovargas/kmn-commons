package errors

type ApiError struct {
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
