package transaction

import (
	"context"
	"errors"
	"yuxuan/car-wash/pkg/db"
)

const (
	PAY_CASH = 1 //非会员消费
	PAY_USER = 2 //会员消费
)

type ITransactionSvc interface {
	AddTransaction(ctx context.Context, req *db.AddTransactionReq) error
	UpdateTransaction(ctx context.Context, req *db.UpdateTransactionReq) error
	GetTransaction(ctx context.Context) ([]*db.Transaction, error)
	DeleteTransaction(ctx context.Context, id int) error
}

type TransactionSvc struct {
	Db *db.SqlConn
}

func NewTransactionSvc(db *db.SqlConn) *TransactionSvc {
	return &TransactionSvc{
		Db: db,
	}
}

func (svc *TransactionSvc) validateRequest(req *db.AddTransactionReq) error {
	if req.Transaction == nil {
		return errors.New("交易为空")
	}
	if req.Transaction.Cost == nil {
		return errors.New("费用为空")
	}
	if req.Transaction.DateTime == nil {
		return errors.New("日期为空")
	}
	if req.Transaction.PayMethod == nil {
		return errors.New("支付方式为空")
	}
	return nil
}

/**
 * function: AddTransaction
 * 添加交易
 */
func (svc *TransactionSvc) AddTransaction(ctx context.Context, req *db.AddTransactionReq) error {
	err := svc.validateRequest(req)
	if err != nil {
		return err
	}
	//会员卡消费
	if *req.Transaction.PayMethod == PAY_USER {
		if req.Transaction.Phone == nil {
			return errors.New("会员消费手机不能为空")
		}
		user, err := svc.Db.GetUserByPhone(*req.Transaction.Phone)
		if err != nil {
			return errors.New("获取会员信息失败")
		}
		if user.RestMoney < *req.Transaction.Cost {
			return errors.New("会员余额不足")
		}
		//更新余额
		err = svc.Db.ChangeUserMoney(*req.Transaction.Phone, user.RestMoney-*req.Transaction.Cost)
		if err != nil {
			return errors.New("更新用户余额错误")
		}
	}
	return svc.Db.AddTransaction(req.Transaction)
}

/**
 * function: GetTransaction
 * 获取所有的交易
 */
func (svc *TransactionSvc) GetTransaction(ctx context.Context) ([]*db.Transaction, error) {
	return svc.Db.GetAllTransactions()
}

/**
 * function: UpdateTransaction
 * 修改交易
 */
func (svc *TransactionSvc) UpdateTransaction(ctx context.Context, req *db.UpdateTransactionReq) error {
	if req.Id == nil {
		return errors.New("id为空")
	}
	if req.Transaction == nil {
		return errors.New("交易为空")
	}
	return svc.Db.UpdateTransaction(req.Transaction, *req.Id)
}

/**
 * function: DeleteTransaction
 * 删除指定交易
 */
func (svc *TransactionSvc) DeleteTransaction(ctx context.Context, id int) error {
	return svc.Db.DeleteTransaction(id)
}
