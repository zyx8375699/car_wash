package login

import (
	"context"
	"encoding/json"
	"net/http"
	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	ROOT          = "/car_wash/v1/login"
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_DELETE = "DELETE"
)

func decodeLoginReqr(_ context.Context, r *http.Request) (interface{}, error) {
	var request db.Login
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeHttpHandler(LoginSvc ILoginSvc, logger kitlog.Logger, router *mux.Router) *mux.Router {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		//kithttp.ServerBefore(kithttp.PopulateRequestContext),
		//kithttp.ServerErrorEncoder(),
	}

	loginEndpoint := MakeLoginEndpoint(LoginSvc)

	router.Methods(METHOD_POST).
		Path(ROOT).
		Handler(kithttp.NewServer(
			loginEndpoint,
			decodeLoginReqr,
			common.EncodeResp,
			opts...,
		))
	return router
}
