package config

type Response struct {
	Code    int16       `json:"code"`
	Payload interface{} `json:"payload"`
	Msg     string      `json:"msg"`
}
