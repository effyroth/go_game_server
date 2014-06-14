package main

import (
// "strconv"
)

type Room struct {
	players []*Player
	notFull bool
}

// var (
// 	rooms       []room
// 	noEmptyRoom = true
// 	currentRoom room
// )

type RoomManeger struct {
	rooms       []*Room
	haveRoom    bool
	currentRoom *Room
}

func NewRoom() *Room {
	return &Room{
		make([]*Player, 10),
		true,
	}
}

func NewRoomManeger() *RoomManeger {
	return &RoomManeger{
		make([]*Room, 10),
		true,
		NewRoom(),
	}
}

func (r *Room) Info() string {
	return "room info online player NO:" + itoa(cap(r.players))
}

func (r *Room) Join(p *Player) {
	if r.notFull {
		r.players = append(r.players, p)
		p.room = r
	}
	if len(r.players) == cap(r.players) {
		r.notFull = false
		roomManeger.haveRoom = false
		roomManeger.currentRoom = nil
	}
}

func (r *RoomManeger) CreateNewRoom(p *Player) {
	roomManeger.currentRoom = NewRoom()
	roomManeger.currentRoom.Join(p)
	roomManeger.rooms = append(roomManeger.rooms, roomManeger.currentRoom)
	roomManeger.haveRoom = true
}
