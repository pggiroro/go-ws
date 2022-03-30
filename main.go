package main

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start server at 9999")
	http.HandleFunc("/test",func(w http.ResponseWriter, r *http.Request) {
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
	})
	log.Fatal(http.ListenAndServe(":9999", nil))
}
