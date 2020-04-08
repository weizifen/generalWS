package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
	"time"
)
const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60  // 多少时间内未进行心跳 断开连接
)
// 用户登录
type Login struct {
	UserId string
	Client *Client
}

// 用户连接
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}
// 初始化
func NewClient(addr string, socket *websocket.Conn) (client *Client) {
	currentTime := uint64(time.Now().Unix())
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     currentTime,
		HeartbeatTime: currentTime,
	}

	return
}
// 用户登录
func (c *Client) Login(userId string) {
	loginTime := uint64(time.Now().Unix())
	c.UserId = userId
	c.LoginTime = loginTime
	// 登录成功=心跳一次
	c.Heartbeat(loginTime)
}

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
	return
}

// 读取客户端数据
func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Println(fmt.Sprintf("读取客户端数据 关闭send client: %v", c))
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println(fmt.Sprintf("读取客户端数据 错误 地址 %v, err: %v", c.Addr, err))

			return
		}

		// 处理程序
		fmt.Println(fmt.Sprintf("读取客户端数据 处理: %v", string(message)))
		ProcessData(c, message)
	}
}
// 向客户端写数据
func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)

		}
	}()

	defer func() {
		clientManager.Unregister <- c
		c.Socket.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				fmt.Println(fmt.Sprintf("Client发送数据 关闭连接 address: %v, ok: %v", c.Addr, ok))

				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
// 发送数据给客户端
func (c *Client) SendMsg(msg []byte) {

	if c == nil {

		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime + heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 是否登录了
func (c *Client) IsLogin() (isLogin bool) {
	// 用户登录了
	if c.UserId != "" {
		isLogin = true
		return
	}
	return
}