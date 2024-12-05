package remote

import (
	"github.com/gorilla/websocket"
	"github.com/kunal-sawant/go-syncplay-client/models"
)

func ConnectToServer(sc models.ServerConfig) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(sc.Url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
