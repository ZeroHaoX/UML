package models

import (
	"github.com/astaxie/beego/logs"
	"../common"
	"fmt"
	// "reflect"
	"time"
)

type Good struct{
	Gno	   float64	`json:"gno"`
	ImDate string 	`json:"imdate"`
	GName string	`json:"gname"`
	Shipper string	`json:"shipper"`
	SPhone string  	`json:"sphone"`
	Count float64	`json:"count"`
	Price float64	`json:"price"`
	ImPrice float64	`json:"imprice"`
}

func GoodsList(page int,pageSize int,orderBy string)(goodsList []Good,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf("Select gid,gname,shipper,count,price,sphone,imprice,imdate from goods_message order by imdate %v limit %v offset %v ",orderBy,pageSize,(page-1)*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var good Good
		var localTime time.Time
		err=rows.Scan(&good.Gno,&good.GName,&good.Shipper,&good.Count,&good.Price,&good.SPhone,&good.ImPrice,&localTime)
		if err!=nil{
			logs.Error(err)
			return
		}
		good.ImDate=localTime.Format("2006-01-02")
		goodsList=append(goodsList,good)
	}
	// logs.Debug(goodsList)
	return
}

func SearchGood(parameter string,filter string)(goods []Good,err error){
	if filter==""||common.HasSpecialCharacter(filter){
		err=fmt.Errorf("Search goods name is error:%v",err)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf(`Select gid,gname,shipper,count,price,sphone,imprice,imdate from goods_message where %v='%v'`,parameter,filter)
	// logs.Debug(queryString)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var good Good
		var localTime time.Time
		err=rows.Scan(&good.Gno,&good.GName,&good.Shipper,&good.Count,&good.Price,&good.SPhone,&good.ImPrice,&localTime)
		if err!=nil{
			logs.Error(err)
			return
		}
		good.ImDate=localTime.Format("2006-01-02")
		goods=append(goods,good)
	}
	return
}

func UpdateGoodMes(good *Good)(err error){
	if good==nil{
		err=fmt.Errorf("UpdateGoodMes good is nil")
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Update goods_message set gname=$1,shipper=$2,count=$3,price=$4,imprice=$5 where gid=$6")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(good.GName,good.Shipper,good.Count,good.Price,good.ImPrice,good.Gno)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func DelGoodMes(gname string,shipper string,imdate string)(err error){
	if gname==""||shipper==""||imdate==""{
		err=fmt.Errorf("DelGoodMes parameters nil")
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Delete from goods_message where imdate=$1 and gname=$2 and shipper=$3")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(imdate,gname,shipper)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func SearchnNumber()(length int,err error){
	rows,err:=db.Query("Select count(imdate) from goods_message")
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		err=rows.Scan(&length)
		if err!=nil{
			logs.Error(err)
			return
		}
	}
	return
}