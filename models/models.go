package models

import (
	"fmt"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/senghoo/captain/modules/settings"
)

var (
	x      *xorm.Engine
	tables []interface{}
	DBCfg  struct {
		Host, Name, User, Passwd string
	}
)

func init() {
	tables = append(tables, new(User), new(GithubAccount), new(DockerServer))
	LoadSetting()
}

func LoadSetting() {
	DBCfg.Host = settings.GetOrDefault("db.host", "localhost:3306")
	DBCfg.Name = settings.GetOrDefault("db.name", "captain")
	DBCfg.User = settings.GetOrDefault("db.user", "root")
	DBCfg.Passwd = settings.GetOrDefault("db.password", "")
}

func getEngine() (*xorm.Engine, error) {
	var cnnstr string
	if DBCfg.Host[0] == '/' { // looks like a unix socket
		cnnstr = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8&parseTime=true",
			DBCfg.User, DBCfg.Passwd, DBCfg.Host, DBCfg.Name)
	} else {
		cnnstr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			DBCfg.User, DBCfg.Passwd, DBCfg.Host, DBCfg.Name)
	}
	return xorm.NewEngine("mysql", cnnstr)
}

func SetEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("Fail to connect to database: %v", err)
	}

	x.SetMapper(core.GonicMapper{})

	// WARNING: for serv command, MUST remove the output to os.stdout,
	// so use log file to instead print to stdout.
	logPath := path.Join(settings.GetOrDefault("log.path", "logs"), "xorm.log")
	os.MkdirAll(path.Dir(logPath), os.ModePerm)
	f, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("Fail to create xorm.log: %v", err)
	}
	x.SetLogger(xorm.NewSimpleLogger(f))
	x.ShowSQL(true)
	return nil
}

func NewEngine() (err error) {
	if err = SetEngine(); err != nil {
		return err
	}

	if err = x.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("sync database struct error: %v\n", err)
	}

	return nil
}
