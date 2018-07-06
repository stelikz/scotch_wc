package main

import (
		"os"
        "log"
        "flag"
        "net/http"
        "github.com/gorilla/websocket"
        "github.com/stelikz/scotch_wc/messages"
        "github.com/stelikz/scotch_wc/db"

)
//
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan messages.Message)  // broadcast channel
// Configure the upgrader
var upgrader = websocket.Upgrader{}

var addr = flag.String(“addr”, “:”+os.Getenv(“PORT”), “http service address”)

func main() {
    // Create a simple file server
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
	// Configure websocket route
    http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	 log.Println("http server started on :8080")
        err := http.ListenAndServe(addr, nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}

func handleConnections(w http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true

	var results messages.Messages

	results = database.Show("localhost", results, "store", "chats")
	
	err = ws.WriteJSON(results)
	if err != nil {
		log.Println(err)
	}

	for {
		var msg messages.Message
		// var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Fatalln(err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		database.Store("localhost", msg, "store", "chats")

		for client := range clients{
			err := client.WriteJSON(msg)
			if err != nil {
				log.Fatalln(err)
				client.Close()
				delete(clients, client)
			}

		}
	}

}