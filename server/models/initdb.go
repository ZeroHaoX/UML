package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	host="119.23.67.53"
	port="5432"
	user="postgres"
	password="postgres"
	dbname="uml"
)

func InitDB()(err error){
	config:=fmt.Sprintf("host=%v port=%v user=%v password=%v sslmodle=disable dbname=%v",host,port,user,password,dbname)
	fmt.Println("config:",config)
	db,err=sql.Open("postgres",config)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}