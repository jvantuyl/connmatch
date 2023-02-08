package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jvantuyl/connmatch/internal/conn"
)

func (a *App) Join(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST accepted", http.StatusMethodNotAllowed)
		return
	}

	req := &conn.Request{When: time.Now(), ResponseChannel: make(chan *conn.Response)}
	defer close(req.ResponseChannel)

	log.Println("received join request, sending to joiner")
	a.JoinChannel <- *req

	select {
	case resp := <-req.ResponseChannel:
		if resp == nil {
			log.Println("timed out waiting for second request")
			http.Error(w, "Timeout", http.StatusRequestTimeout)
		} else {
			var (
				order string
			)

			if resp.First {
				order = "First"
			} else {
				order = "Last"
			}

			log.Printf("received response to %s connection, assigned id %s, after %d µs", order, resp.Conn.Identity(), resp.Conn.Delay.Microseconds())

			w.WriteHeader(http.StatusOK)
			_, err := fmt.Fprintf(w, "%s %s\n", resp.Conn.Identity(), order)
			if err != nil {
				log.Printf("Error %v writing connection response %s (%s)", err, resp.Conn.Identity(), order)
			}
		}
	}
}
