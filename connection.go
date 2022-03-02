package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

type Connection struct {
	url string
	ws  *websocket.Conn
}

var SERVER_IDS = []int{5, 6, 7, 8}

func NewConnection() *Connection {
	serverIndex := rand.Int() % (len(SERVER_IDS) - 1)
	println(serverIndex)
	conn := &Connection{
		url: fmt.Sprintf("wss://ws%d.blitzortung.org:3000/", SERVER_IDS[serverIndex]),
	}

	return conn
}

func (conn *Connection) Connect() error {
	ws, _, err := websocket.DefaultDialer.Dial(conn.url, nil)
	if err != nil {
		return err
	}
	conn.ws = ws
	ws.WriteJSON(&TimeInitPacket{Time: 0})
	ws.WriteJSON(&ServerSelectionPacket{WsServer: conn.url})
	return nil
}

func (conn *Connection) Read() (*StrikePacket, error) {
	if conn.ws == nil {
		return nil, errors.New("not connected")
	}
	var response StrikePacket
	err := conn.ws.ReadJSON(&response)
	if err != nil {
		panic(err)
	}
	return &response, nil
}
