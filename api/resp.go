package api

type Resp struct {
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	DevMessage string      `json:"dev_message"`
}
