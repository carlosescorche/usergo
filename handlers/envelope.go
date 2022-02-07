package handlers

type HTTPResponseEnvelope struct {
	HTTPStatus int         `json:"httpStatus"`
	Data       interface{} `json:"data"`
}

type HTTPResponseError struct {
	Code    string      `json:"code"`
	Op      string      `json:"op"`
	Message interface{} `json:"message"`
	Status  int         `json:"status"`
	Extra   interface{} `json:"extra"`
}

type ErrorInfo struct {
	Code    string
	Op      string
	Message interface{}
}
