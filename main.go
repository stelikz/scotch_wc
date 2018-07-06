package main

import (
        "log"
        "net/http"
        "github.com/gorilla/websocket"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "fmt"
        "./messages"
        "./db"

)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan messages.Message)           // broadcast channel
// Configure the upgrader
var upgrader = websocket.Upgrader{}

type Message struct {
    Username string `json:"username"`
    Message string `json:"message"`
    ID uint64 `json:"id"`

}

type Messages []Message

func main() {
    // Create a simple file server
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)
	// Configure websocket route
    http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	 log.Println("http server started on :8080")
        err := http.ListenAndServe(":8080", nil)
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
	// var results Messages

	database.Show("localhost", results, "store", "chats")
	// session, _ := mgo.Dial("localhost")

	// anotherSession := session.Copy()
	// defer anotherSession.Close()

	// c := session.DB("store").C("chats")
	// err2 := c.Find(bson.M{}).All(&results)
	// if err2 != nil {
	// 	log.Println(err2)
	// }
	
	err = ws.WriteJSON(results)
	if(err!= nil) {
		fmt.Println("err")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Not err")

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

		// session, _ := mgo.Dial("localhost")

		// anotherSession := session.Copy()
		// defer anotherSession.Close()

		// c := session.DB("store").C("chats")
		// c.Insert(msg)
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