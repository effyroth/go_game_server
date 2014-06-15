package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
)

var (
	cmd = `
{
    "Command" : "CMD_JOINROOM",
    	"Param":
        {
            "Username" : "xxx",
            "Password" : "aa",
            "Email" : "a@x.com"
        }
}
`
	cmdStartGame = `
{
    "Command" : "CMD_STARTGAME",
    	"Param":
        {
            "Username" : "xxx",
            "Password" : "aa",
            "Email" : "a@x.com"
        }
}
`
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:1234/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(cmd)
	if _, err := ws.Write([]byte(cmd)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 1024)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	// log.Println(msg)
	fmt.Printf("Received: %s.\n", msg[:n])

	if _, err := ws.Write([]byte(cmdStartGame)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 1024)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	// log.Println(msg)
	fmt.Printf("Received: %s.\n", msg[:n])
}
