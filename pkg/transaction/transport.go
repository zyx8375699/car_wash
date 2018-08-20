package transaction

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"

	"github.com/go-kit/kit/auth/basic"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	ROOT          = "/car_wash/v1/transaction"
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

func decodeAddTransactionReqr(_ context.Context, r *http.Request) (interface{}, error) {
	var request db.AddTransactionReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetTransactionReqr(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeDeleteTransactionReqr(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.New("param err")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return DeleteTransactionReq{Id: id}, nil
}

func MakeHttpHandler(TransactionSvc ITransactionSvc, logger kitlog.Logger, router *mux.Router, cfg *common.LoginConfig) *mux.Router {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		//kithttp.ServerErrorEncoder(),
	}

	addTransactionEndpoint := MakeAddTransactionEndpoint(TransactionSvc)
	getTransactionEndpoint := MakeGetTransactionEndpoint(TransactionSvc)
	deleteTransactionEndpoint := MakeDeleteTransactionEndpoint(TransactionSvc)

	router.Methods(METHOD_GET).
		Path(ROOT).
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, "Example Realm")(getTransactionEndpoint),
			decodeGetTransactionReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_POST).
		Path(ROOT).
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, "Example Realm")(addTransactionEndpoint),
			decodeAddTransactionReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_DELETE).
		Path(ROOT + "/{id}").
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, "Example Realm")(deleteTransactionEndpoint),
			decodeDeleteTransactionReqr,
			common.EncodeResp,
			opts...,
		))
	return router
}
