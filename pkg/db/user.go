package db

import "errors"

type User struct {
	Id             int     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`                    //会员id
	Name           *string `json:"name" gorm:"column:name;type:varchar(8);not null"`        //会员姓名
	Phone          *int64  `json:"phone" gorm:"column:phone;type:bigint;not null;unique"`   //手机号
	Type           *int    `json:"type" gorm:"column:type;type:int;not null"`               //会员类型
	RestMoney      int64   `json:"restMoney" gorm:"column:rest_money;type:bigint;not null"` //余额
	VehicleLicense string  `json:"vehicleLicense" gorm:"column:license;type:varchar(16)"`   //车牌号
}

type AddUserReq struct {
	User *User `json:"user"`
}

type UpdateUserReq struct {
	Phone *int64 `json:"phone"`
	User  *User  `json:"user"`
}

type AddMoneyReq struct {
	Phone int   `json:"phone"`
	Money int64 `json:"money"`
}

func (User) TableName() string {
	return "user"
}

/**
 * function: GetAllUser
 * 获取所有用户
 */
func (c *SqlConn) GetAllUsers() ([]*User, error) {
	var us []*User
	err := c.db.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

/**
 * function: GetUserByPhone
 * 根据电话号码查询用户
 */
func (c *SqlConn) GetUserByPhone(phone int) (*User, error) {
	u := new(User)
	err := c.db.Where("phone = ?", phone).First(u).Error
	return u, err
}

/**
 * function: AddUser
 * 添加新用户
 */
func (c *SqlConn) AddUser(u *User) error {
	if u == nil {
		return errors.New("用户为空")
	}
	return c.db.Save(u).Error
}

/**
 * function: UpdateUser
 * 更新用户
 */
func (c *SqlConn) UpdateUser(u *User, phone int64) error {
	return c.db.Model(User{}).Where("phone = ?", phone).Updates(u).Error
}

/**
 * function: DeleteUser
 * 删除用户
 */
func (c *SqlConn) DeleteUser(id int) error {
	return c.db.Delete(&User{
		Id: id,
	}).Error
}

/**
 * function: ChangeUserMoney
 * 修改用户的储值
 */
func (c *SqlConn) ChangeUserMoney(phone int, money int64) error {
	return c.db.Model(User{}).Where("phone = ?", phone).Updates(User{
		RestMoney: money,
	}).Error
}
