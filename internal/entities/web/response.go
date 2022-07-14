package web

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
