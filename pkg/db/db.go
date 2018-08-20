package db

import (
	"yuxuan/car-wash/pkg/common"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlConn struct {
	dbPath string
	db     *gorm.DB
	cfg    *common.LoginConfig
}

func NewSqlConn(path string, cfg *common.LoginConfig) *SqlConn {
	return &SqlConn{
		dbPath: path,
		cfg:    cfg,
	}
}

/**
 * function: InitDb
 * 初始化数据库
 */
func (c *SqlConn) InitDb() error {
	db, err := gorm.Open("mysql", c.dbPath)
	if err != nil {
		return err
	}
	c.db = db
	err = c.CreateTables()
	if err != nil {
		return err
	}
	return nil
}

/**
 * function: CreateTables
 * 建表
 */
func (c *SqlConn) CreateTables() error {
	if !c.db.HasTable(&Transaction{}) {
		err := c.db.CreateTable(&Transaction{}).Error
		if err != nil {
			return err
		}
	}
	if !c.db.HasTable(&User{}) {
		err := c.db.CreateTable(&User{}).Error
		if err != nil {
			return err
		}
	}
	if !c.db.HasTable(&Login{}) {
		err := c.db.CreateTable(&Login{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
