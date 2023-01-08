package bean

type ResponseError struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
