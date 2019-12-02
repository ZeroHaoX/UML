package models

import (
	"github.com/astaxie/beego/logs"
	"../common"
	"fmt"
	// "reflect"
)

type Good struct{
	ID int	`json:"gid"`
	Name string		`json:"gname"`
	Shipper string	`json:"shipper"`
	Count int	`json:"count"`
	Price float32	`json:"price"`
	ImPrice float32	`json:"imprice"`
}

func GoodsList(page int,pageSize int,orderBy string)(goodsList []*Good,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf("Select * from goods order by %v limit %v offset %v",orderBy,pageSize,page*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var good Good
		err=rows.Scan(&good.ID,&good.Name,&good.Shipper,&good.Count,&good.Price,&good.ImPrice)
		if err!=nil{
			logs.Error(err)
			return
		}
		goodsList=append(goodsList,&good)
	}
	return
}

func SearchGood(filter string,parameter string)(goods []Good,err error){
	if filter==""||common.HasSpecialCharacter(filter){
		err=fmt.Errorf("Search goods name is error:%v",err)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf(`Select * from goods where %v like '%%v%'`,filter,parameter)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var good Good
		err=rows.Scan(&good.ID,&good.Name,&good.Shipper,&good.Count,&good.Price,&good.ImPrice)
		if err!=nil{
			logs.Error(err)
			return
		}
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
	stmt,err:=db.Prepare("Update goods set gid=$1,gname=$2,shipper=$3,count=$4,price=$5,imprice=$6")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(good.ID,good.Name,good.Shipper,good.Count,good.Price,good.ImPrice)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func DelGoodMes(gid string)(err error){

	return
}