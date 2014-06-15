package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

var (
	PlayerIdCenter = 0
)

func GetPlayerId() int {
	PlayerIdCenter++
	return PlayerIdCenter
}

type Player struct {
	ws        *websocket.Conn // ws of this Player
	playerId  int             // uid of table userinfo
	charID    uint32          // cid of table character
	character string
	room      *Room
	position  int
	isHost    bool
}

func (p *Player) GetRoomInfo() string {
	return p.room.Info()
}

func (p *Player) Info() string {
	log.Println(p)
	return "Player info"
}

func (p *Player) Leave() {
	p.room.Remove(p)
}

func NewPlayer(ws *websocket.Conn) *Player {
	return &Player{
		ws:       ws,
		playerId: GetPlayerId(),
		isHost:   false,
	}
}

func (p *Player) IsHost() bool {
	return p.isHost
}
