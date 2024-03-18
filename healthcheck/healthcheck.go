package healthcheck

import "net/http"

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Application is alive and healthy\n"))
}
