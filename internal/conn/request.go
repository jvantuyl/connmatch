package conn

import "time"

type Request struct {
	When            time.Time
	ResponseChannel chan *Response
}
