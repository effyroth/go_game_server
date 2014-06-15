package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"log"
)

var (
	cmdJoinRoom = `
{
    "Command" : "CMD_JOINROOM",
    	"Param":
        {
            "RoomType" : "Chat"
        }
}
`
	cmdStartGame = `
{
    "Command" : "CMD_STARTGAME",
    	"Param":
        {
            "Username" : "xxx"
        }
}
`

	cmdChat = "CMD_CHAT"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:1234/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(cmd)

	go ReadLoop(ws)

	WriteLoop(ws)
}

func ReadLoop(ws *websocket.Conn) {
	var msg = make([]byte, 1024)
	var n int = 0
	var err error
	for {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
	}

}

func WriteLoop(ws *websocket.Conn) {
	WriteCommand(ws, cmdJoinRoom)
	for {
		command := GetChatCommand()
		log.Println(command)
		WriteCommand(ws, command)
	}

}

func WriteCommand(ws *websocket.Conn, cmd string) {
	if _, err := ws.Write([]byte(cmd)); err != nil {
		log.Fatal(err)
	}
}

func GetChatCommand() string {
	message := GetChatMessage()
	return CommandJson(cmdChat, message)
}

func GetChatMessage() string {
	var s string
	fmt.Scan(&s)
	return s
}

type Param struct {
	Message interface{}
}

type JsonCommand struct {
	Command string
	Param   Param
}

func CommandJson(command string, message interface{}) string {
	msg := JsonCommand{
		Command: command,
		Param: Param{
			Message: message,
		},
	}
	return makeJson(msg)
}

func makeJson(v interface{}) string {
	// fmt.Printf("%v\n", v)
	bin, err := json.Marshal(v)
	if err != nil {
		// log.Println("error!!!!!")
		// log.Println(err.Error())
		return "Marshal json fail!"
	}
	return string(bin)
}
