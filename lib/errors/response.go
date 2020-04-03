package errors

type CustomError struct {
	Code    int64  `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
