package models

import (
	"github.com/astaxie/beego/logs"
)

type Log struct{
	Time string
	Api string
	Name string
	Level int 
	Detial string
}

func GetLog()(logsList []Log,err error){
	rows,err:=db.Query("Select * from log")
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var mlog Log
		rows.Scan(&mlog.Time,&mlog.Level,&mlog.Detial)
	}
	return
}