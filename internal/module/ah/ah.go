// ah is the application health module
package ah

import (
	"net/http"
)

// Mount mounts ah module into mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", health)
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}
