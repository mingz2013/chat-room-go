package chat_room_go

import (
	"chat-room-go/base"
	"chat-room-go/msg"
	"chat-room-go/robot-manager"
	"chat-room-go/table-manager"
	"sync"
)

func StartLocalTest() {
	var wg sync.WaitGroup

	msgIn := make(chan msg.Msg)
	msgOut := make(chan msg.Msg)

	robotManager := robot_manager.NewRobotManager(msgIn, msgOut)
	roomManager := table_manager.NewTableManager("", msgOut, msgIn)

	base.RunProcessor(&wg, roomManager)
	base.RunProcessor(&wg, robotManager)

	wg.Wait()
}
