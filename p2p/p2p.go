package p2p

import (
	"fmt"
	"net/http"

	"github.com/Hunobas/nomadcoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	for {
		fmt.Println("Waiting 4 msg. . .")
		_, p, err := conn.ReadMessage()
		fmt.Println("Msg arrived!\n\n")
		utils.HandleErr(err)
		fmt.Printf("%s", p)
	}
}
