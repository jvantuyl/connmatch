package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) Stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET accepted", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	result := make(map[string]int64)

	for _, conn := range a.Connections {
		result[conn.Identity()] = conn.Delay.Microseconds()
	}

	encoder := json.NewEncoder(w)

	encoder.SetIndent("", "  ")
	encoder.Encode(result)
}
