package main

import (
	"fmt"
	"net/http"
	"strconv"

	"../pkg/utilities"
	"../pkg/websocketserver"
)

func main() {

	// Testing out functionality...

	// This changes all output to go to a log file rather than a console
	utilities.AssignLogFile()

	// Tests out the timer functionality
	for x := 0; x < 10; x++ {
		utilities.SecondSleeper()
		fmt.Println(strconv.Itoa(x+1) + " Timer has slept for a second")
	}

	hub := websocketserver.NewHub("baseHub")

	go hub.Run()
	// start up the websocket server
	http.HandleFunc("ws", func(writer http.ResponseWriter, request *http.Request) {
		websocketserver.ServeWs(hub, writer, request)
	})

}
