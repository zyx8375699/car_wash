package user

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
	ROOT          = "/car_wash/v1/user"
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_DELETE = "DELETE"
	METHOD_PUT    = "PUT"
	REALM         = "example realm"
)

func decodeAddUserReqr(_ context.Context, r *http.Request) (interface{}, error) {
	var request db.AddUserReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateUserReqr(_ context.Context, r *http.Request) (interface{}, error) {
	var request db.UpdateUserReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAddMoneyReqr(_ context.Context, r *http.Request) (interface{}, error) {
	var request db.AddMoneyReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserReqr(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeDeleteUserReqr(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, errors.New("param err")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	return DeleteUserReq{Id: id}, nil
}

func MakeHttpHandler(userSvc IUserSvc, logger kitlog.Logger, router *mux.Router, cfg *common.LoginConfig) *mux.Router {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}

	addUserEndpoint := MakeAddUserEndpoint(userSvc)
	getUserEndpoint := MakeGetUserEndpoint(userSvc)
	updateUserEndpoint := MakeUpdateUserEndpoint(userSvc)
	deleteUserEndpoint := MakeDeleteUserEndpoint(userSvc)
	addMoneyEndpoint := MakeAddMoneyEndpoint(userSvc)

	router.Methods(METHOD_GET).
		Path(ROOT).
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, REALM)(getUserEndpoint),
			decodeGetUserReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_POST).
		Path(ROOT).
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, REALM)(addUserEndpoint),
			decodeAddUserReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_PUT).
		Path(ROOT).
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, REALM)(updateUserEndpoint),
			decodeUpdateUserReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_DELETE).
		Path(ROOT + "/{id}").
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, REALM)(deleteUserEndpoint),
			decodeDeleteUserReqr,
			common.EncodeResp,
			opts...,
		))
	router.Methods(METHOD_POST).
		Path(ROOT + "/money").
		Handler(kithttp.NewServer(
			basic.AuthMiddleware(cfg.UserName, cfg.Password, REALM)(addMoneyEndpoint),
			decodeAddMoneyReqr,
			common.EncodeResp,
			opts...,
		))
	return router
}
