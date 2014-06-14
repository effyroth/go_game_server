package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

var (
	roomManeger *RoomManeger
)

func main() {
	// log.Print("initing database ...")
	// initDB()
	// log.Print("initing game script data ...")
	// initGameData()

	log.Print("starting socket server ...")
	//TODO read port from config
	http.Handle("/", websocket.Handler(gameServer))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	initRoomManeger()
}

func initRoomManeger() {
	roomManeger = NewRoomManeger()
}
