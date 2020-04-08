package game

import (
	jsoniter "github.com/json-iterator/go"
	"time"
	"wxGameWebSocket/common"
	"wxGameWebSocket/game/model"
	"wxGameWebSocket/utils/logger"
	"wxGameWebSocket/utils/ws"
)
var log = logger.New("game")

func LoginController(client *ws.Client, key int32, message []byte) (code uint32, msg string, data interface{})  {
	code = common.OK
	request := &model.Login{}
	err := jsoniter.Unmarshal(message, request)
	if err != nil {
		log.Error("解析数据失败", "err", err.Error())
	}
	if request.UserId == "" {
		code = common.UnauthorizedUserId
		log.Error("非法的用户ID")
		return
	}

	if client.IsLogin() {
		log.Info("用户登录 用户已经登录", "key", key, "userId", client.UserId)
		code = common.OperationFailure

		return
	}

	client.Login(request.UserId)

	login := &ws.Login{
		UserId: request.UserId,
		Client: client,
	}
	ws.AddLoginUser(login)
	return
}
func HeartbeatController(client *ws.Client, key int32, message []byte) (code uint32, msg string, data interface{})  {
	code = common.OK
	currentTime := uint64(time.Now().Unix())
	request := &model.HeartBeat{}
	err := jsoniter.Unmarshal(message, request)
	if err != nil {
		log.Error("HeartbeatController 解析数据失败", "err", err.Error())
	}

	if !client.IsLogin() {
		log.Info("心跳接口 用户未登录")
		code = common.NotLoggedIn
		return
	}

	client.Heartbeat(currentTime)

	return
}