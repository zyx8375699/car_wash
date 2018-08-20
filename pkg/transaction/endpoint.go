package transaction

import (
	"context"

	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"

	"github.com/go-kit/kit/endpoint"
)

type DeleteTransactionReq struct {
	Id int
}

func MakeAddTransactionEndpoint(s ITransactionSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(db.AddTransactionReq)
		err = s.AddTransaction(ctx, &req)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}

func MakeGetTransactionEndpoint(s ITransactionSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		Transactions, err := s.GetTransaction(ctx)
		return common.NewCommonResponse(true, err, Transactions), nil
	}
}

func MakeDeleteTransactionEndpoint(s ITransactionSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(DeleteTransactionReq)
		err = s.DeleteTransaction(ctx, req.Id)
		return common.NewCommonResponse(ok, err, nil), nil
	}
}
