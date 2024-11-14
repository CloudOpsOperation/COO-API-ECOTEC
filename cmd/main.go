package main

import (
	"log"

	"github.com/CloudOpsOperation/COO-API-ECOTEC/cmd/api"
	"github.com/CloudOpsOperation/COO-API-ECOTEC/config"
	"github.com/CloudOpsOperation/COO-API-ECOTEC/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	log.Println("Connecting to " + config.Envs.PublicHost + ":" + config.Envs.DBAddr + " database name:" + config.Envs.DBName)

	db, err := db.NewSQLConnection(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPwd,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		panic(err)
	}

	srv := api.NewAPI(":"+config.Envs.ApiPort, db)
	if err := srv.Start(); err != nil {
		panic(err)
	}

}
