package app

import (
	"log"
	"time"

	"github.com/jvantuyl/connmatch/internal/conn"
)

func (a *App) Joiner() {
	for {
		// Wait for first connection
		log.Println("joiner waiting for first connection")
		first := <-a.JoinChannel

		// Got first connection
		log.Println("joiner received first connection")
		timer := time.NewTimer(a.Timeout)

		// Wait for second connection
		select {
		case second := <-a.JoinChannel:
			// calculate delay
			delay := second.When.Sub(first.When)
			log.Printf("joiner received second connection after %d µs", delay)

			// build connection
			c := &conn.Connection{Ident: a.NextIdent, Delay: delay}

			// build responses
			cr1 := &conn.Response{Conn: *c, First: true}
			cr2 := &conn.Response{Conn: *c, First: false}

			// send responses
			first.ResponseChannel <- cr1
			second.ResponseChannel <- cr2

			// save connection and manage connection cache
			a.Connections.Push(c)
			for a.Connections.Len() > 50000 {
				a.Connections.Pop()
			}

			// Clean up for next iteration
			timer.Stop()
			a.NextIdent++

		case <-timer.C:
			// Time Out Waiting
			log.Println("joiner timed out waiting for second connection")
			first.ResponseChannel <- nil
		}
	}
}
