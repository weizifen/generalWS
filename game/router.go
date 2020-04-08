package game

import "wxGameWebSocket/utils/ws"

// Websocket 路由
func WebSocketInit() {
	ws.Register(Login, LoginController)
	ws.Register(HeartBeat, HeartbeatController)
	//ws.Register(websocket.GetEnergy, websocket.GetEnergyController)
}