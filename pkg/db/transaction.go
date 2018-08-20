package db

import "errors"

type Transaction struct {
	Id        int     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`                    //主键，无业务意义
	Type      *string `json:"type" gorm:"column:type;type:varchar(16);not null"`       //交易类型
	Cost      *int64  `json:"cost" gorm:"column:cost;type:int;not null"`               //花费
	DateTime  *string `json:"date" gorm:"column:date;type:datetime;not null"`          //时间
	License   string  `json:"license" gorm:"column:license;type:varchar(16);not null"` //车牌号码
	Phone     *int    `json:"phone" gorm:"column:phone;type:bigint"`                   //用户电话
	PayMethod *int    `json:"payMethod" gorm:"column:method;type:int;not null"`        //付费方法
}

type AddTransactionReq struct {
	Transaction *Transaction `json:"transaction"`
}

type UpdateTransactionReq struct {
	Id          *int         `json:"id"`
	Transaction *Transaction `json:"transaction"`
}

func (Transaction) TableName() string {
	return "transaction"
}

/**
 * function: GetAllTransaction
 * 获取所有交易
 */
func (c *SqlConn) GetAllTransactions() ([]*Transaction, error) {
	var ts []*Transaction
	err := c.db.Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

/**
 * function: AddTransaction
 * 添加一条新交易
 */
func (c *SqlConn) AddTransaction(t *Transaction) error {
	if t == nil {
		return errors.New("交易为空")
	}
	return c.db.Save(t).Error
}

/**
 * function: UpdateTransaction
 * 更新交易信息
 */
func (c *SqlConn) UpdateTransaction(t *Transaction, id int) error {
	return c.db.Model(Transaction{}).Where("id = ?", id).Updates(t).Error
}

/**
 * function DeleteTransaction
 * 删除交易
 */
func (c *SqlConn) DeleteTransaction(id int) error {
	return c.db.Delete(&Transaction{
		Id: id,
	}).Error
}
