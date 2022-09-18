package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	darkneos "github.com/sktt1ryze/ygopro-proxy/DarkNeos"
)

const TARGET_PORT = ":8000"
const PROXY_PORT = ":3344"
const CHANNEL_SIZE = 0x1000
const BUFFER_SIZE = 0x1000
const TIME_OUT = 5

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0x1000,
	WriteBufferSize: 0x1000,
}

func ygoEndpoint(w http.ResponseWriter, r *http.Request) {
	defer log.Println("ygoEndpoint finished")

	upgrader.CheckOrigin = wsChecker

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to ws://localhost" + TARGET_PORT + " [websocket] succeeded!")

	tcp, err := net.Dial("tcp", "127.0.0.1"+PROXY_PORT)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to " + "12.0.0.1" + PROXY_PORT + " [tcp] succeeded!")

	wsCh := make(chan []byte, CHANNEL_SIZE)
	tcpCh := make(chan []byte, CHANNEL_SIZE)
	wsStopCh := make(chan bool, CHANNEL_SIZE)
	tcpStopCh := make(chan bool, CHANNEL_SIZE)

	defer func() {
		wsStopCh <- true
		tcpStopCh <- true

		close(wsStopCh)
		close(tcpStopCh)
	}()

	go wsProxy(ws, wsCh, wsStopCh)
	go tcpProxy(tcp, tcpCh, tcpStopCh)

	for {
		select {
		case wsBuf, ok := <-wsCh:
			if !ok {
				return
			}

			if _, err = tcp.Write(wsBuf); err != nil {
				log.Println(err)
				return
			}
		case tcpBuf, ok := <-tcpCh:
			if !ok {
				return
			}

			if err = ws.WriteMessage(websocket.BinaryMessage, tcpBuf); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// todo: generic
func wsProxy(ws *websocket.Conn, Ch chan<- []byte, stopCh <-chan bool) {
	defer ws.Close()
	defer close(Ch)

	for {
		select {
		case _, ok := <-stopCh:
			log.Println("wsProxy recv stop singal, exit. channel closed: ", ok)
			return
		default:
			// if err := ws.SetReadDeadline(time.Now().Add(time.Second * TIME_OUT)); err != nil {
			//   log.Println(err)
			//   return
			// }

			messageType, buffer, err := ws.ReadMessage()
			if err != nil {
				//   if err, ok := err.(net.Error); ok && err.Timeout() {
				//     continue
				//   }

				log.Println(err)
				return
			}

			if messageType == websocket.CloseMessage {
				log.Println("Websocket closed")
				return
			}

			buffer, err = darkneos.Transform(buffer, darkneos.ProtobufToRawBuf)
			if err != nil {
				log.Println(err)
				return
			}

			Ch <- buffer
		}
	}
}

func tcpProxy(tcp net.Conn, Ch chan<- []byte, stopCh <-chan bool) {
	defer tcp.Close()
	defer close(Ch)

	reader := bufio.NewReader(tcp)
	buffer := make([]byte, BUFFER_SIZE)

	for {
		select {
		case _, ok := <-stopCh:
			log.Println("tcpProxy recv stop singal, exit. channel closed: ", ok)
			return
		default:
			if err := tcp.SetReadDeadline(time.Now().Add(time.Second * TIME_OUT)); err != nil {
				log.Println(err)
				return
			}

			_, err := reader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					continue
				}

				if err, ok := err.(net.Error); ok && err.Timeout() {
					continue
				}

				log.Println(err)
				return
			}

			buffer, err = darkneos.Transform(buffer, darkneos.RawBufToProtobuf)
			if err != nil {
				log.Println(err)
				return
			}

			Ch <- buffer
		}
	}
}

func wsChecker(r *http.Request) bool { return true }

func setupRoutes() {
	http.HandleFunc("/", ygoEndpoint)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Llongfile)

	setupRoutes()

	log.Fatal(http.ListenAndServe(TARGET_PORT, nil))
}
