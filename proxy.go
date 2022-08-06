package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const TARGET_PORT = ":8000"
const PROXY_PORT = ":3344"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0x1000,
	WriteBufferSize: 0x1000,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func ygoEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = wsChecker

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	log.Println("Websocket connected")

	proxyHandler(ws)
}

func proxyHandler(ws *websocket.Conn) {
	// todo: spawn channel

	for {
		messageType, buf, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		if messageType == websocket.CloseMessage {
			log.Println("Websocket closed")
			return
		}

		fmt.Println(string(buf))

		if err := ws.WriteMessage(messageType, buf); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func wsChecker(r *http.Request) bool { return true }

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ygo", ygoEndpoint)
}

func main() {
	setupRoutes()

	log.Fatal(http.ListenAndServe(TARGET_PORT, nil))
}
