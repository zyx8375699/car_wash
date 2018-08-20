package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"yuxuan/car-wash/pkg/common"
	"yuxuan/car-wash/pkg/db"
	"yuxuan/car-wash/pkg/login"
	"yuxuan/car-wash/pkg/transaction"
	"yuxuan/car-wash/pkg/user"

	kitlog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

const (
	MYSQL_CONN = "%s:%s@tcp(%s:%s)/%s?charset=utf8"
)

var (
	ArgMysqlHost     = flag.String("mysqlHost", "119.29.224.144", "mysql host")
	ArgMysqlPort     = flag.String("mysqlPort", "3306", "mysql port")
	ArgMysqlUserName = flag.String("mysqlUserName", "root", "mysql username")
	ArgMysqlPassword = flag.String("mysqlPassword", "woshizyx333", "mysql password")
	ArgMysqlDb       = flag.String("mysqlDb", "car_wash", "mysql db name")
	ArgHttpAddr      = flag.String("addr", ":9090", "ip address")
	ArgLoginConf     = flag.String("loginConfig", "login.json", "login config file")
)

func main() {
	flag.Parse()
	errs := make(chan error, 10)

	var logger kitlog.Logger
	logger = kitlog.NewJSONLogger(os.Stderr)
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	cfg, err := readConfig(*ArgLoginConf)
	if err != nil {
		errs <- err
	}
	mysqlConn := fmt.Sprintf(MYSQL_CONN, *ArgMysqlUserName, *ArgMysqlPassword, *ArgMysqlHost, *ArgMysqlPort, *ArgMysqlDb)
	db := db.NewSqlConn(mysqlConn, cfg)
	err = db.InitDb()

	var transactionSvc transaction.ITransactionSvc
	transactionSvc = &transaction.TransactionSvc{
		Db: db,
	}
	var userSvc user.IUserSvc
	userSvc = &user.UserSvc{
		Db: db,
	}
	var loginSvc login.ILoginSvc
	loginSvc = &login.LoginSvc{
		Db: db,
	}

	if err != nil {
		errs <- err
	}
	router := mux.NewRouter()
	router = transaction.MakeHttpHandler(transactionSvc, logger, router, cfg)
	router = user.MakeHttpHandler(userSvc, logger, router, cfg)
	router = login.MakeHttpHandler(loginSvc, logger, router)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("level", "info", "msg", fmt.Sprintf("car wash service has been started! Listening on %s", *ArgHttpAddr))
		errs <- http.ListenAndServe(*ArgHttpAddr, router)
	}()

	if e := <-errs; e != nil {
		logger.Log("level", "error", "msg", e.Error())
		panic(e)
	}
}

func readConfig(file string) (*common.LoginConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	res := new(common.LoginConfig)
	err = json.Unmarshal(data, res)
	return res, err
}
