package main

import (
	// "io"
	"code.google.com/p/go.net/websocket"
	"github.com/bitly/go-simplejson"
	"log"
)

type commandHandler func(p *Player, command string, param *simplejson.Json)

// handler map for "Command"
var cmdHandlers = map[string]commandHandler{
	// "CMD_REGISTER":    cmdRegisterHander,
	"CMD_JOINROOM": cmdJoinRoomHandler,
}

func cmdJoinRoomHandler(p *Player, command string, param *simplejson.Json) {
	rtnCode := 0
	var rtnMsg interface{}
	defer func() {

		if 0 == rtnCode && 0 == p.charID {
			// client need display create character UI
			// rtnMsg = "JoinRoom"
		} else {
			rtnMsg, _ = errCodes[rtnCode]
		}
		rtnJson := responseJson(command, rtnCode, rtnMsg)
		if err := websocket.Message.Send(p.ws, rtnJson); err != nil {
			log.Printf("Send fail for cmLoginHander")
		}
	}()

	if roomManeger.haveRoom {
		JoinCurrentRoom(p)
	} else {
		roomManeger.CreateNewRoom(p)
	}
	rtnMsg = p.GetRoomInfo()
}

func JoinCurrentRoom(p *Player) {
	roomManeger.currentRoom.Join(p)
}

type jsonReturn struct {
	Code    int
	Message interface{}
}

type jsonResponse struct {
	Command string
	Return  jsonReturn
}

func responseJson(command string, code int, message interface{}) string {
	msg := jsonResponse{
		Command: command,
		Return: jsonReturn{
			Code:    code,
			Message: message,
		},
	}
	return makeJson(msg)
}

func commandDispatcher(p *Player, js *simplejson.Json) {
	rtnCode := 0
	command := ""

	// defer need to be placed at header of func
	defer func() {
		// only send error message here
		if 0 != rtnCode {
			rtnMsg, _ := errCodes[rtnCode]
			rtnJson := responseJson(command, rtnCode, rtnMsg)
			if err := websocket.Message.Send(p.ws, rtnJson); err != nil {
				log.Printf("Send fail for commandDispatcher")
			}
		}
	}()

	command, err := js.Get("Command").String()
	if err != nil {
		rtnCode = 1
		return
	}

	param, ok := js.CheckGet("Param")
	if ok {
		handler, ok := cmdHandlers[command]
		if ok {
			handler(p, command, param)
			return
		}
	}

	rtnCode = 1

}
