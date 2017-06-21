package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "sitepointgoapp/routers"
)

var RunMode string
var Cfg = beego.AppConfig

func main() {
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	dbSslmode := Cfg.String("db_sslmode")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")
	dbLink := ("user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSslmode)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbLink, maxIdleConn, maxOpenConn)
	// orm.RunSyncdb("default", true, true)

	RunMode = Cfg.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}

	beego.Run()
}
