package main

import (
	// "strconv"
	"log"
)

var (
	RoomIdCenter = 0
)

func GetRoomId() int {
	RoomIdCenter++
	return RoomIdCenter
}

type Room struct {
	NO      int
	players map[int]*Player
	notFull bool
	cap     int
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
		GetRoomId(),
		make(map[int]*Player, 10),
		true,
		10,
	}
}

func NewRoomManeger() *RoomManeger {
	return &RoomManeger{
		make([]*Room, 0),
		true,
		NewRoom(),
	}
}

func (r *Room) Info() string {
	log.Println(r)
	return "room info online player NO:" + itoa(len(r.players)) + "\nroom NO:" + itoa(r.NO)
}

func (r *Room) Join(p *Player) {
	if r.notFull {
		for i := 1; i <= r.cap; i++ {
			_, exist := r.players[i]
			if !exist {
				r.players[i] = p
				p.position = i
				p.room = r
				break
			}
		}

		// r.players = append(r.players, p)
		// p.room = r
	}
	if len(r.players) == r.cap {
		r.notFull = false
		roomManeger.haveRoom = false
		roomManeger.currentRoom = nil
	}
}

func (r *Room) Remove(p *Player) {
	delete(r.players, p.position)
}

func (r *Room) StartGame() {
	r.Publish("game start")
}

func (r *Room) Publish(rtnJson string) {
	for _, p := range r.players {
		Send(p.ws, rtnJson)
	}
}

func (r *Room) GetGameInfo() string {
	return "game info"
}

func (r *RoomManeger) CreateNewRoom(p *Player) {
	roomManeger.currentRoom = NewRoom()
	roomManeger.currentRoom.Join(p)
	p.isHost = true
	roomManeger.rooms = append(roomManeger.rooms, roomManeger.currentRoom)
	roomManeger.haveRoom = true
}

func (r *RoomManeger) GetRoomsInfo() {
	log.Println(r)
	log.Printf("%v", r)
}
