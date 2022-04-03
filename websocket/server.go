package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

func upper(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err)
			break

		}
		fmt.Println("reveived from client: " + reply)
		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	http.Handle("/test1", websocket.Handler(upper))
	http.HandleFunc("/", index)

	log.Println("Websocket Server start 127.0.0.1:6432")
	if err := http.ListenAndServe(":6432", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
