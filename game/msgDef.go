package game
// 消息定义
const (
	Login           = 1  // 登陆
	HeartBeat       = 2  // 心跳包

	GetEnergy = 3 // 获得能量

	MessageCmdExit  = 98 // 通知玩家离开游戏
	MessageCmdEnter = 99 // 通知有人进入游戏
)

func getMsgText(code int32) string {
	textMap := map[int32]string{
		Login:           "用户登陆",
		HeartBeat:       "心跳包",

		GetEnergy: "获得能量", // 获得能量


		MessageCmdEnter: "enter",
		MessageCmdExit:  "exit",
	}
	return textMap[code]
}