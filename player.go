package main

import (
	"code.google.com/p/go.net/websocket"
)

type Player struct {
	ws        *websocket.Conn // ws of this Player
	userID    uint32          // uid of table userinfo
	charID    uint32          // cid of table character
	character string
	room      *Room
}

func (p *Player) GetRoomInfo() string {
	return p.room.Info()
}

func (p *Player) Info() string {
	return "Player info"
}
