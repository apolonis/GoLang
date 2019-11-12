package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//	fmt.Println("DEBUG:Error upgrading", err)
		fmt.Fprintf(w, "DEBUG:Error upgrading", err)
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	fmt.Printf("DEBUG ERROR IN FUNCTION WRITER \n")
	//for {

	countryList, err := getCountries()
	// fmt.Println(countryList)
	if err != nil {
		fmt.Println("DEBUG:Error getCountries() in writer()", err)
	}
	for i := 0; i < len(countryList); i++ {
		//fmt.Println(country[i])
		recvMsgType, recvMsg, recvErr := conn.ReadMessage()
		fmt.Println("RECV:", recvMsgType, string(recvMsg), recvErr)

		jsonString, err := json.Marshal(countryList[i])

		fmt.Println("DEBUG: Printing jsonString", jsonString)

		if err != nil {
			fmt.Println("DEBUG:ERROR MARSHALING", err)
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
			fmt.Println("DEBUG:ERROR write message", err)
			//return
			//}
		}

	}

}
