package types

type Response struct {
	Code    int16       `json:"code"`
	Payload interface{} `json:"payload"`
	Msg     string      `json:"msg"`
}

type UnsignedResponse struct {
	Msg interface{} `json:"msg"`
}

type SignedResponse struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
}
