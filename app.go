package chat_room_go

import (
	"github.com/mingz2013/game-table-go/base"
	"github.com/mingz2013/game-table-go/msg"
	"github.com/mingz2013/game-table-go/robot-manager"
	"github.com/mingz2013/game-table-go/table-manager"
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
