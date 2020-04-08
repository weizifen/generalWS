package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"wxGameWebSocket/config"
	"wxGameWebSocket/utils/helper"
	"wxGameWebSocket/utils/logger"
)
var log = logger.New("ws")
var (
	clientManager = NewClientManager() // 管理者
	serverIp      string
	serverPort    string
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	HandshakeTimeout:  5 * time.Second,
	// CheckOrigin: 处理跨域问题，线上环境慎用
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	},
}

// 启动程序
func StartWebSocket() {
	serverIp = helper.GetServerIp()
	webSocketPort := config.ServerConfig.WebSocketPort
	rpcPort := config.ServerConfig.RPCPort

	serverPort = rpcPort

	http.HandleFunc("/ws", handleWSConnect)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	// 添加处理程序
	go clientManager.start()
	fmt.Println(fmt.Sprintf("WebSocket 启动程序成功rpc地址 %v:%v", serverIp, serverPort))
	fmt.Println(fmt.Sprintf("WebSocket 启动程序成功ws地址 %v:%v", serverIp, webSocketPort))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", webSocketPort), nil); err != nil {
		fmt.Println("ws服务启动失败")
	}
}
func handleWSConnect(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&wsUpgrader).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())


	client := NewClient(conn.RemoteAddr().String(), conn)

	// 读取分离，减少收发数据堵塞的可能
	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}
