package healthcheck

import "net/http"

func Healthcheck() string {

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Application is alive and healthy\n"))
	})

	// Start the HTTP server on port 8080 in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()
	return "health check enabled [ port=8080, path=/alive ]"
}
