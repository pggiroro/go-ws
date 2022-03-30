package api

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	fmt.Println("start server at 9999")
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			fmt.Println("read msg is " + string(msg))
			if err != nil {
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				// handle error
			}
		}
	}()
}