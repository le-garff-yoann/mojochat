package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"testing"

	"github.com/gorilla/websocket"
	"github.com/posener/wstest"
)

func TestIndex(t *testing.T) {
	a := NewApp("postgresql://x:x@x/x")

	req, _ := http.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()
	h := http.Handler(a.Router.Get("index").GetHandler())

	h.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusOK; got != want {
		t.Errorf("rr.Code: got %q, want %q", got, want)
	}
}

func TestChat(t *testing.T) {
	pgURI := os.Getenv("TEST_PG_DSN")
	if len(pgURI) == 0 {
		t.Skip("Assign a PostgreSQL connection string (eg., postgresql://pg:pg@db/pg?sslmode=disable) to $TEST_PG_DSN")
	}

	a := NewApp(pgURI)

	if _, err := a.DB.Exec("SELECT 1"); err != nil {
		t.Fatal(err)
	}

	h := http.Handler(a.Router.Get("chat").GetHandler())

	d := wstest.NewDialer(h)
	c, resp, err := d.Dial("ws://x/x", nil)
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	if got, want := resp.StatusCode, http.StatusSwitchingProtocols; got != want {
		t.Errorf("resp.StatusCode: got %q, want %q", got, want)
	}

	time.Sleep(time.Second) // FIXME: The WebSocket Protocol Upgrade isn't done when we send the first message?

	body := "Hello World!"
	var msg Msg

	c.WriteMessage(websocket.TextMessage, []byte(body))
	c.ReadJSON(&msg)

	if got, want := msg.Body, body; got != want {
		t.Errorf("msg.Body: got %q, want %q", got, want)
	}
}
