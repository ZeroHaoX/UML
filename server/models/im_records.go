package models

import (
	"../common"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"time"
)

type ImportRecord struct{
	ID string	`json:"id"`
	Imdate string	`json:"imdate"`
	GoodName string	`json:"gname"`
	Price float64	`json:"imprice"`
	Count int	`json:"imcount"`
	TotalPrice float64	`json:"imtotalprice"`
	Shipper string	`json:"shipper"`
	SPhone string	`json:"sphone"`
	Detial string	`json:"detial"`
	// Profit float32	`json:"profit"`
}

func ImportList(page int,pageSize int,orderBy string)(importList []ImportRecord,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	// start:=nextMonth(year,month)
	// end:=nextMonth(year,month+1)
	queryString:=fmt.Sprintf("Select imid,gname,imprice,imtotalprice,imcount,detial,shipper,sphone,imdate from import_records order by imdate %v limit %v offset %v ",orderBy,pageSize,(page-1)*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var importRecord ImportRecord
		var imlocalTime time.Time
		err=rows.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.SPhone,&imlocalTime)
		if err!=nil{
			logs.Error(err)
			return
		}
		importRecord.Imdate=imlocalTime.Format("2006-01-02")
		importList=append(importList,importRecord)
	}
	return
}

func Import(importRecord ImportRecord)(err error){
	if reflect.DeepEqual(importRecord,ImportRecord{}){
		err=fmt.Errorf("Import importRecord is error:importRecord=%v",importRecord)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Insert into import_records(imid,gname,imprice,imcount,imtotalprice,shipper,sphone,detial) values($1,$2,$3,$4,$5,$6,$7,$8)")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(importRecord.ID,importRecord.GoodName,importRecord.Price,importRecord.Count,importRecord.TotalPrice,importRecord.Shipper,importRecord.SPhone,importRecord.Detial)
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
	_,err=stmt.Exec(importRecord.Imdate,importRecord.GoodName,importRecord.Price,importRecord.Count,importRecord.TotalPrice,importRecord.Shipper,importRecord.SPhone,importRecord.Detial,importRecord.ID)
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
	importRecord=new(ImportRecord)
	rows,err:=db.Query(`Select * from import_records where imid=$1`,id)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		err=rows.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.SPhone,&importRecord.Imdate)
		if err!=nil{
			logs.Error(err)
			return
		}
	}
	return
}

func SearchImportList(id string)(importRecords []ImportRecord,err error){
	if id==""||common.HasSpecialCharacter(id){
		err=fmt.Errorf("SearchImports get id=%v",id)
		logs.Error(err)
		return
	}
	// importRecord=new(ImportRecord)
	rows,err:=db.Query("Select imid,gname,imprice,imtotalprice,imcount,detial,shipper,sphone,imdate from import_records where imid=$1",id)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var importRecord ImportRecord
		var imlocalTime time.Time
		err=rows.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.SPhone,&imlocalTime)
		if err!=nil{
			logs.Error(err)
			return
		}
		importRecord.Imdate=imlocalTime.Format("2006-01-02")
		importRecords=append(importRecords,importRecord)
	}
	return
}

func SearchImportsByTime(year int,month int)(importRecords []ImportRecord,err error){
	nextDate:=nextMonth(year,month)
	nowDate:=timeToString(year,month)
	queryString:=fmt.Sprintf(`Select imid,gname,imprice,imtotalprice,imcount,detial,shipper,sphone,imdate from import_records where imdate>='%v' and imdate<'%v' `,nowDate,nextDate)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var importRecord ImportRecord
		var imlocalTime time.Time
		err=rows.Scan(&importRecord.ID,&importRecord.GoodName,&importRecord.Price,&importRecord.TotalPrice,&importRecord.Count,&importRecord.Detial,&importRecord.Shipper,&importRecord.SPhone,&imlocalTime)
		if err!=nil{
			logs.Error(err)
			return
		}
		importRecord.Imdate=imlocalTime.Format("2006-01-02")
		importRecords=append(importRecords,importRecord)
	}
	return
}


func SearchImportNumber()(length int,err error){
	rows,err:=db.Query("Select count(imid) from import_records")
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