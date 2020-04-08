package ws

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"sync"
	"wxGameWebSocket/common"
)

type DealFunc func(client *Client, key int32, message []byte) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[int32]DealFunc)
	handlersRWMutex sync.RWMutex
)

// 注册处理函数
func Register(key int32, value DealFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func ProcessData(client *Client, message []byte)  {
	log.Info("client -> server", "v", string(message))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()
	req := &Request{}
	err := jsoniter.Unmarshal(message, req)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))
		return
	}

	requestData, err := jsoniter.Marshal(req.V)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))

		return
	}
	key := req.K
	var (
		code uint32
		msg  string
		data interface{}
	)
	// 采用 map 注册的方式
	if value, ok := getHandlers(key); ok {
		code, msg, data = value(client, key, requestData)
	} else {
		code = common.RoutingNotExist
		log.Info("路由不存在", "key", key)
	}
	msg = common.GetErrorMessage(code, msg)

	responseHead := NewResponseHead(key, code, msg, data)

	headByte, err := jsoniter.Marshal(responseHead)
	if err != nil {
		log.Error("处理数据 json Marshal", "err", err.Error())
		return
	}

	client.SendMsg(headByte)
	log.Info("server -> client", "message", responseHead)
	return

}
// 获得处理方法
func getHandlers(key int32) (value DealFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	value, ok = handlers[key]

	return
}