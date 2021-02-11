package table_manager

import (
	"chat-room-go/msg"
	// "github.com/mingz2013/mahjong-table-go/table"
)

type UserSession struct {
	UserId  int
	TableId int
}

func NewUserSession(userId int) UserSession {
	return UserSession{UserId: userId}
}

type TableSession struct {
	Table   Table
	MsgIn   chan msg.Msg
	MsgOut  chan msg.Msg
	TableId int
}

func (s *TableSession) Close() {
	s.Table.Close()
	close(s.MsgIn)
}

func NewTableSession(table Table, msgIn chan msg.Msg, msgOut chan msg.Msg, tableId int) TableSession {
	return TableSession{Table: table, MsgIn: msgIn, MsgOut: msgOut, TableId: tableId}
}
