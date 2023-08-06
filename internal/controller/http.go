package controller

type HTTPBody struct {
	Message string `json:"message"`
}

type HTTPResponse struct {
	StatusCode int          `json:"statusCode"`
	Body       *HTTPBody     `json:"body"`
	Error      string `json:"error"`
}
