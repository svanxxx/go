package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	routeRegister()
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
