package ws

import "encoding/json"

/************************  请求数据  **************************/
// 通用请求数据格式
type Request struct {
	K int32       `json:"k"`
	V interface{} `json:"v,omitempty"` // 数据 json
}

/************************  响应数据  **************************/
type Head struct {
	K int32     `json:"cmd"`      // 消息的cmd 动作  消息的Id
	V *Response `json:"response"` // 消息体
}

type Response struct {
	Code    uint32      `json:"code"`
	CodeMsg string      `json:"codeMsg"`
	Data    interface{} `json:"data"` // 数据 json
}

// 设置返回消息
func NewResponseHead(cmd int32, code uint32, codeMsg string, data interface{}) *Head {
	response := NewResponse(code, codeMsg, data)

	return &Head{K: cmd, V: response}
}

func (h *Head) String() (headStr string) {
	headBytes, _ := json.Marshal(h)
	headStr = string(headBytes)

	return
}

func NewResponse(code uint32, codeMsg string, data interface{}) *Response {
	return &Response{Code: code, CodeMsg: codeMsg, Data: data}
}
