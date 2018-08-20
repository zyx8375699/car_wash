package db

import (
	"encoding/base64"
	"errors"
	"fmt"
)

const (
	RETURN_VALUE = "Basic %s"
)

type Login struct {
	Id       int    `json:"-" gorm:"primary_key;AUTO_INCREMENT"`
	User     string `json:"user" gorm:"column:name;type:varchar(16);not null;unique"`
	Password string `json:"password" gorm:"column:password;type:varchar(256);not null"`
}

func (Login) TableName() string {
	return "login"
}

func (c *SqlConn) ValidateLogin(user string, p string) (string, error) {
	password, err := c.getPassWordByName(user)
	if err != nil {
		return "", err
	}
	if password != p {
		return "", errors.New("用户名密码错误")
	}
	input := []byte(fmt.Sprintf("%s:%s", c.cfg.UserName, c.cfg.Password))
	encode := base64.StdEncoding.EncodeToString(input)
	return fmt.Sprintf(RETURN_VALUE, encode), nil
}

func (c *SqlConn) getPassWordByName(name string) (string, error) {
	u := new(Login)
	err := c.db.Where("name = ?", name).First(u).Error
	return u.Password, err
}
