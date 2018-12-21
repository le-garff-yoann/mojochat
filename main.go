package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	botUUID = "root"
)

// Msg : This is a chat message
type Msg struct {
	Datetime time.Time `json:"datetime"`
	Body     string    `json:"body"`
	UUID     string    `json:"uuid"`
	Me       bool      `json:"me"`
}

// App : The application with it's *mux.Router and *pg.DB
type App struct {
	Router *mux.Router

	DB         *pg.DB
	wsUpgrader *websocket.Upgrader
}

// NewApp : Initialize the application
func NewApp(pgURI string) *App {
	r := mux.NewRouter()

	dbOpts, err := pg.ParseURL(pgURI)
	if err != nil {
		log.Print(err)
	}

	db := pg.Connect(dbOpts)

	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  0,
		WriteBufferSize: 0,
	}

	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		c, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print(err)

			return
		}
		defer c.Close()

		uuid := uuid.NewV4().String()

		gb := make(chan interface{})

		pt := time.NewTicker(time.Second)
		defer pt.Stop()

		go func() {
			ln := db.Listen("chat")
			defer ln.Close()

			cch := ln.Channel()

			for {
				select {
				case n := <-cch:
					var m Msg

					if err := json.Unmarshal([]byte(n.Payload), &m); err != nil {
						log.Print(err)
					}

					m.Me = m.UUID == uuid

					err = c.WriteJSON(&m)
					if err != nil {
						gb <- nil
					}
				case <-pt.C:
					err := c.WriteMessage(websocket.PingMessage, []byte("pong"))
					if err != nil {
						gb <- nil
					}
				case <-gb:
					log.Printf("%s just left the chat!", uuid)

					return
				default:
				}
			}
		}()

		log.Printf("%s just entered the chat!", uuid)

		for {
			messageType, p, err := c.ReadMessage()
			if err != nil {
				gb <- nil

				break
			}

			switch messageType {
			case websocket.TextMessage:
				json, _ := json.Marshal(&Msg{time.Now(), string(p), uuid, false})

				go func() {
					if _, err := db.Exec("NOTIFY chat, ?", string(json)); err != nil {
						log.Print(err)
					}
				}()
			default:
			}
		}
	}).Name("chat")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("vue/mojochat"))).Methods("GET").Name("index")

	return &App{r, db, &wsUpgrader}
}

// Run : http.ListenAndServe the application
func (a *App) Run(port string) {
	http.Handle("/", a.Router)

	log.Printf("Listening on %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	a := NewApp(os.Getenv("MOJOCHAT_PG"))

	a.Run(os.Getenv("MOJOCHAT_PORT"))
}
