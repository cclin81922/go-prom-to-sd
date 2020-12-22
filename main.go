package main

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main () {
	Collect()

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
