package model

type ResponseMap map[string]interface{}

type Response struct {
	Status    string `json:"stat"`
	Message   string `json:"msg"`
	Data      any    `json:"data"`
	Timestamp string `json:"ts"`
	RequestId string `json:"id"`
}
