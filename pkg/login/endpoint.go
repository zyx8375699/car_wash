package login

import (
	"context"
	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"

	"github.com/go-kit/kit/endpoint"
)

func MakeLoginEndpoint(s ILoginSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(db.Login)
		token, err := s.ValidateLogin(ctx, req.User, req.Password)
		return common.NewCommonResponse(ok, err, token), nil
	}
}
