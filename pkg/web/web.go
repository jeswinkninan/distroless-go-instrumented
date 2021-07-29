package web

import (
	"net/http"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		a := []byte("Hello World Telemetry")
		w.Write(a)
	}
}
