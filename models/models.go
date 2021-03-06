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
	tables = append(tables, new(User), new(GithubAccount), new(DockerServer),
		new(Workspace), new(Repository), new(Build), new(Workflow), new(GithubWebhook),
		new(Registry), new(DockerImage), new(DockerImageVersion))
	LoadSetting()
}

func LoadSetting() {
	DBCfg.Host = settings.GetOrDefault("DB_HOST", "localhost:3306")
	DBCfg.Name = settings.GetOrDefault("DB_NAME", "captain")
	DBCfg.User = settings.GetOrDefault("DB_USER", "root")
	DBCfg.Passwd = settings.GetOrDefault("DB_PASSWORD", "")
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
	logPath := path.Join(settings.GetOrDefault("LOG_PATH", "logs"), "xorm.log")
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

func GetByID(id int64, obj interface{}) (has bool, err error) {
	return x.Id(id).Get(obj)
}

func Insert(beans ...interface{}) (int64, error) {
	return x.Insert(beans...)
}
