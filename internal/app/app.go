package app

import (
	"container/heap"
	"net/http"
	"sync"
	"time"

	"github.com/jvantuyl/connmatch/internal/conn"
)

type App struct {
	Timeout     time.Duration
	NextIdent   int
	Mux         *http.ServeMux
	JoinMutex   sync.Mutex
	JoinChannel chan conn.Request
	Connections conn.Cache
}

func New() *App {
	return &App{}
}

func (a *App) Init() {
	a.Timeout = DefaultTimeout

	a.JoinChannel = make(chan conn.Request)
	a.Connections = make(conn.Cache, 0, 50000)

	a.Mux = http.NewServeMux()

	heap.Init(&a.Connections)

	a.Mux.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) { a.Join(w, r) })
	a.Mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) { a.Stats(w, r) })
}

func (a *App) Run() {
	go a.Joiner()
	http.ListenAndServe(":8080", a.Mux)
}
