package user

import (
	"context"
	"errors"
	"yuxuan/car-wash/pkg/db"
)

type IUserSvc interface {
	AddUser(ctx context.Context, req *db.AddUserReq) error       //新增会员
	GetUser(ctx context.Context) ([]*db.User, error)             //查询所有会员
	AddMoney(ctx context.Context, money int64, phone int) error  //充值
	UpdateUser(ctx context.Context, req *db.UpdateUserReq) error //修改会员信息
	DeleteUser(ctx context.Context, id int) error                //删除会员
}

type UserSvc struct {
	Db *db.SqlConn
}

func NewUserSvc(db *db.SqlConn) *UserSvc {
	return &UserSvc{
		Db: db,
	}
}

/**
 * function: AddUser
 * 添加会员
 */
func (svc *UserSvc) AddUser(ctx context.Context, req *db.AddUserReq) error {
	if req.User.Name == nil {
		return errors.New("会员姓名不能为空")
	}
	if req.User.Phone == nil {
		return errors.New("会员手机不能为空")
	} else if !validatePhone(*req.User.Phone) {
		return errors.New("会员手机号格式错误")
	}
	if req.User.Type == nil {
		return errors.New("会员级别不能为空")
	}
	return svc.Db.AddUser(req.User)
}

/**
 * function: GetUser
 * 获取所有的会员
 */
func (svc *UserSvc) GetUser(ctx context.Context) ([]*db.User, error) {
	return svc.Db.GetAllUsers()
}

/**
 * function: UpdateUser
 * 更新会员信息
 */
func (svc *UserSvc) UpdateUser(ctx context.Context, req *db.UpdateUserReq) error {
	if req.Phone == nil {
		return errors.New("电话为空")
	}
	if req.User == nil {
		return errors.New("用户为空")
	}
	return svc.Db.UpdateUser(req.User, *req.Phone)
}

/**
 * function: DeleteUser
 * 删除会员
 */
func (svc *UserSvc) DeleteUser(ctx context.Context, id int) error {
	return svc.Db.DeleteUser(id)
}

/**
 * function: AddMoney
 * 充值
 */
func (svc *UserSvc) AddMoney(ctx context.Context, money int64, phone int) error {
	u, err := svc.Db.GetUserByPhone(phone)
	if err != nil {
		return err
	}
	return svc.Db.ChangeUserMoney(phone, money+u.RestMoney)
}
