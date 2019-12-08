package models

import (
	// "time"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type Account struct{
	Profit float64	`json:"profit"`
	Turnover float64	`json:"turnover"`
}

func timeToString(year int ,month int)(timeString string){
	var m string
	if month<10{
		m="0"+strconv.Itoa(month)
	}
	y:=strconv.Itoa(year)
	timeString= y+"-"+m
	logs.Debug(timeString)
	return 
}

func GetAccount(year int,month int)(account *Account,err error){
	timeString:=timeToString(year,month)
	rows,err:=db.Query("Select profit,turnover from account_view where accountdate=$1",timeString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		err=rows.Scan(&account.Profit,&account.Turnover)
		if err!=nil{
			logs.Error(err)
			return
		}
	}
	return
}

func AccountList()(accounts []Account,err error){

	return
}