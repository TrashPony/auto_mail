package webSocket

import (
	"net/http"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {

	ReadSocket(w, r, r.URL.Path)
}
