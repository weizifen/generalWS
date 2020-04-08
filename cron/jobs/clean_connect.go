package jobs

import (
	"fmt"
	"time"
	"wxGameWebSocket/utils/ws"
)

func init() {
	remind := &CleanConnectRunner{}
	addJobs(remind)
}

// RemindRunner 任务提醒
type CleanConnectRunner struct {
}

// Name 任务名
func (*CleanConnectRunner) Name() string {
	return "RemindRunner"
}

// Cron 任务执行间隔 每分钟执行一次
func (*CleanConnectRunner) Cron() string {
	// https://www.cnblogs.com/zuxingyu/p/6023919.html
	return "0 */1 * * * *"
}

// Run 任务执行体
func (*CleanConnectRunner) Run() {
	fmt.Println(fmt.Sprintf("定时清理玩家链接 %v", time.Now()))
	// 定时清理超时连接
	ws.ClearTimeoutConnections()
}
