package model
// 登录请求数据
type Login struct {
	UserId       string `json:"userId,omitempty"` // UUID
	//Token        string `json:"token"`            // Token
}
// 心跳请求数据
type HeartBeat struct {
}