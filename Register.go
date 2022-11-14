package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type registerRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type registerResponse struct {
	ID  uint64 `json:"id"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request registerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeRegisterEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		var man UserManager
		user, err := man.RegisterUser(req.Name, req.Password, req.Email)
		if err != nil {
			return registerResponse{0, err.Error()}, nil
		}
		return registerResponse{user.id, ""}, nil
	}
}

func routeRegister() {
	registerHandler := httptransport.NewServer(
		makeRegisterEndpoint(),
		decodeRegisterRequest,
		encodeResponse,
	)

	http.Handle("/register", registerHandler)

}
