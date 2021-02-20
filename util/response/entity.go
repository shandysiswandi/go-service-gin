package response

// SuccessResponse is
type SuccessResponse struct {
	Error   bool        `json:"error" xml:"error"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

// ErrorResponse is
type ErrorResponse struct {
	Error   bool   `json:"error" xml:"error"`
	Message string `json:"message" xml:"message"`
	Trace   error  `json:"trace" xml:"trace"`
}
