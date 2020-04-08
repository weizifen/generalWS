package main

import (
	"sync"
	task "wxGameWebSocket/cron"
	"wxGameWebSocket/game"
	"wxGameWebSocket/utils/ws"
)
var wg sync.WaitGroup

func main() {
	wg.Add(3)
	// 注册ws路由
	go registerRouter()
	// 启动ws服务器
	go wsInit()
	// cron任务
	go taskExec()

	wg.Wait()
}
func registerRouter()  {
	game.WebSocketInit()
	wg.Done()
}
func wsInit()  {
	ws.StartWebSocket()
	wg.Done()
}
func taskExec()  {
	task.TaskInit()
	wg.Done()
}