package conn

import (
	"fmt"
	"time"
)

type Connection struct {
	Ident int
	Delay time.Duration
}

func (conn Connection) Identity() string {
	return fmt.Sprintf("id%d", conn.Ident)
}
