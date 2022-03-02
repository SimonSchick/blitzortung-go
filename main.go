package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	conn := NewConnection()
	err := conn.Connect()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	input := make(chan *StrikePacket)

	go func() {
		defer close(input)
		for {
			message, err := conn.Read()
			if err != nil {
				break
			}
			// send the trade to the channel
			input <- message
		}
	}()

	for strike := range input {
		jsonRes, _ := json.MarshalIndent(strike, "", "\t")
		fmt.Println(string(jsonRes))
	}
}
