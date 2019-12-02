package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	host="localhost"
	port="5432"
	user="postgres"
	password="z83313420"
	dbname="uml"
)

func InitDB()(err error){
	config:=fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v",host,port,user,password,dbname)
	db,err=sql.Open("psql",config)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}