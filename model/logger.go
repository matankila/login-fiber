package model

type RequestInfo struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Ip     string `json:"ip"`
	UID    string `json:"uid"`
}
