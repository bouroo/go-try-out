package handler

type JSendStatus string

const (
	JSendStatusSuccess JSendStatus = "success"
	JSendStatusError   JSendStatus = "error"
	JSendStatusFail    JSendStatus = "fail"
)

type JSend struct {
	Status  JSendStatus `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    any         `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
}
