package login

import (
	"context"
	"yuxuan/car-wash/pkg/db"
)

type ILoginSvc interface {
	ValidateLogin(ctx context.Context, user string, password string) (string, error)
}

type LoginSvc struct {
	Db *db.SqlConn
}

func NewLoginSvc(db *db.SqlConn) *LoginSvc {
	return &LoginSvc{
		Db: db,
	}
}

func (svc *LoginSvc) ValidateLogin(ctx context.Context, user string, password string) (string, error) {
	return svc.Db.ValidateLogin(user, password)
}
