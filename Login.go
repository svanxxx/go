package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type loginResponse struct {
	JWT string `json:"jwt"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeLoginEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		var man UserManager
		jwt, err := man.Login(req.Name, req.Password)
		if err != nil {
			return loginResponse{"", err.Error()}, nil
		}
		return loginResponse{jwt, ""}, nil
	}
}

func routeLogin() {
	loginHandler := httptransport.NewServer(
		makeLoginEndpoint(),
		decodeLoginRequest,
		encodeResponse,
	)

	http.Handle("/login", loginHandler)
}
