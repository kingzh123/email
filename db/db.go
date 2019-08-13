package DbsEmail

import (
	"database/sql"
	"feifanEmailCron/configs"
	"feifanEmailCron/public"
	"github.com/go-sql-driver/mysql"
)

type Db struct{
	Dbw *sql.DB
}

var (
	err error
	config *configsEmail.Config
	db Db
)

func init() {
	config, err = publicEmail.ReadYamlConfig("./config.yaml")
	if err != nil {
		panic(err)
	}
	Connect(&db)
}

func GetConfig() *configsEmail.Config {
	return config
}

func GetDb () *sql.DB {
	return db.Dbw
}

func Connect(db *Db){
	//params
	cfg := mysql.NewConfig()
	cfg.User = config.Conf.Mysql.Username
	cfg.Passwd = config.Conf.Mysql.Password
	cfg.Net = config.Conf.Mysql.Net
	cfg.Addr = config.Conf.Mysql.Host
	cfg.DBName = config.Conf.Mysql.Database
	dsn := cfg.FormatDSN()
	//connect database
	db.Dbw, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}