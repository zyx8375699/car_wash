package user

import (
	"context"
	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"

	"github.com/go-kit/kit/endpoint"
)

type DeleteUserReq struct {
	Id int
}

func MakeAddUserEndpoint(s IUserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(db.AddUserReq)
		err = s.AddUser(ctx, &req)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}

func MakeGetUserEndpoint(s IUserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		Users, err := s.GetUser(ctx)
		return common.NewCommonResponse(true, err, Users), nil
	}
}

func MakeDeleteUserEndpoint(s IUserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(DeleteUserReq)
		err = s.DeleteUser(ctx, req.Id)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}

func MakeUpdateUserEndpoint(s IUserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(db.UpdateUserReq)
		err = s.UpdateUser(ctx, &req)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}

func MakeAddMoneyEndpoint(s IUserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(db.AddMoneyReq)
		err = s.AddMoney(ctx, req.Money, req.Phone)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}
