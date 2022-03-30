package api

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsflate"
	"github.com/gobwas/ws/wsutil"
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	e := wsflate.Extension{
		Parameters: wsflate.Parameters{
			ServerNoContextTakeover: true,
			ClientNoContextTakeover: true,
		},
	}
	fmt.Println("start server at 9999")
	u := ws.HTTPUpgrader{
		Negotiate: e.Negotiate,
	}
	conn, _, _, err := u.Upgrade(r, w)
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	//go func() {
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
	//}()
}
