package models

import (
	"../common"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
)

type ImportRecord struct{
	ID string	`json:"id"`
	Imdate string	`json:"imdate"`
	GoodName string	`jsoimount"`
	Price float32	`json:"imprice"`
	Count int	`json:"imcount"`
	TotalPrice float32	`json:"imtotalprice"`
	Shipper string	`json:"shipper"`
	Phone string	`json:"phone"`
	Detial string	`json:"detial"`
	// Profit float32	`json:"profit"`
}

func ImportList(year int,month int,page int,pageSize int)(importList []*ImportRecord,err error){
	if !common.CheckDate(year,month){
		err=fmt.Errorf("year or month is error:year=%v month=%v",year,month)
		logs.Error(err)
		return
	}
	if page<0||pageSize<0{
		err=fmt.Errorf("page and pageSize error:page=%v,pageSize=%v",page,pageSize)
		logs.Error(err)
		return
	}
	start:=strToTimeStr(year,month)
	end:=strToTimeStr(year,month+1)
	queryString:=fmt.Sprintf("select * from export_records(id,gname,imprice,imtotalprice,imcount,detial,shipper,phone,imdate) where edate>=%v::timestamp and edate<%v::timestamp limit %v offset %v",start,end,pageSize,page*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var importRecord ImportRecord
		err=rows.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.Phone,&importRecord.Imdate)
		if err!=nil{
			logs.Error(err)
			return
		}
		importList=append(importList,&importRecord)
	}
	return
}

func Import(importRecord ImportRecord)(err error){
	if reflect.DeepEqual(importRecord,ImportRecord{}){
		err=fmt.Errorf("Import importRecord is error:importRecord=%v",importRecord)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Insert into import_records values($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(importRecord.ID,importRecord.Imdate,importRecord.GoodName,importRecord.Price,importRecord.Count,importRecord.TotalPrice,importRecord.Shipper,importRecord.Phone,importRecord.Detial)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func UpdateImport(importRecord ImportRecord)(err error){
	if reflect.DeepEqual(importRecord,ImportRecord{}){
		err=fmt.Errorf("UpdateImport importRecord is error:importRecord=%v",importRecord)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Update import_records set gname=$1,imprice=$2,imcount=$3,imtotalprice=$4,shipper=$5,phone=$6,detial=$7 where id =$8")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(importRecord.Imdate,importRecord.GoodName,importRecord.Price,importRecord.Count,importRecord.TotalPrice,importRecord.Shipper,importRecord.Phone,importRecord.Detial,importRecord.ID)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func DelImport(id string)(err error){
	if id==""{
		err=fmt.Errorf("Del Import id=%v",id)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Delete from import_records where imid=$1")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(id)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func SearchImports(id string)(importRecord *ImportRecord,err error){
	if id==""||common.HasSpecialCharacter(id){
		err=fmt.Errorf("SearchImports get id=%v",id)
		logs.Error(err)
		return
	}
	row:=db.QueryRow(`Select * from import_records where id=$1`,id)
	if row!=nil{
		err=row.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.Phone,&importRecord.Imdate)
		if err!=nil{
			logs.Error(err)
			return
		}
	}
	return
}