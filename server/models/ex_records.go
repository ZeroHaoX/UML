package models

import (
	"github.com/astaxie/beego/logs"
	"../common"
	"fmt"
	// "time"
	"strconv"
	// "reflect"
)

type ExportRecord struct{
	ID string	`json:"id"`
	Edate string	`json:"edate"`
	GID string	`json:"gid"`
	GoodName string	`json:"gname"`
	Count int	`json:"ecount"`
	Price float32	`json:"eprice"`
	TotalPrice float32	`json:"etotalprice"`
	Buyer string	`json:"buyer"`
	Phone string	`json:"phone"`
	Detial string	`json:"detial"`
	// Profit float32	`json:"profit"`
}

func strToTimeStr(year int,month int)(timeString string){
	var m string
	if month<10{
		m="0"+strconv.Itoa(month)
	}
	if month==13{
		year++
		m="0"+strconv.Itoa(1)
	}
	// if day<10{
	// 	d="0"+strconv.Itoa(month)
	// }
	timeString=fmt.Sprintf("%v-%v-01",year,m)
	return
}

func ExportList(year int,month int)(exportList []*ExportRecord,err error){
	if !common.CheckDate(year,month){
		err=fmt.Errorf("year or month is error:year=%v month=%v",year,month)
		logs.Error(err)
		return
	}
	start:=strToTimeStr(year,month)
	end:=strToTimeStr(year,month+1)
	rows,err:=db.Query("select * from export_records_view(id,gid,gname,eprice,etotalprice,ecount,detial,buyer,phone,edate) where edate>=$1::timestamp and edate<$2::timestamp",start,end)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var exoprtRecord ExportRecord
		err=rows.Scan(&exoprtRecord.ID,&exoprtRecord.GID,&exoprtRecord.GoodName,&exoprtRecord.Price,&exoprtRecord.TotalPrice,&exoprtRecord.Count,&exoprtRecord.Detial,&exoprtRecord.Buyer,&exoprtRecord.Phone,exoprtRecord.Edate)
		if err!=nil{
			logs.Error(err)
			return
		}
	}
	return
}

func Exoprt(exportRecord *ExportRecord)(err error){
	if exportRecord==nil{
		err=fmt.Errorf("Export exportRecort is error:exportRecord=%v",err)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Insert into export_records values($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(exportRecord.ID,exportRecord.GID,exportRecord.Count,exportRecord.Price,exportRecord.TotalPrice,exportRecord.Buyer,exportRecord.Phone,exportRecord.Detial)
	if err!=nil{
		err=fmt.Errorf("Export stmt error:%v",err)
		logs.Error(err)
		return
	}
	return
}

func UpdateExport(exportRecord *ExportRecord)(err error){
	if exportRecord==nil{
		err=fmt.Errorf("UpdateExport exportRecord is error:exportRecord=%v",exportRecord)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Update export_records set gid=$1,eprice=$2,ecount=$3,etotalprice=$4,buyer=$5,phone=$6,detial=$7 where id=$8")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(exportRecord.GID,exportRecord.GoodName,exportRecord.Price,exportRecord.Count,exportRecord.TotalPrice,exportRecord.Buyer,exportRecord.Phone,exportRecord.Detial,exportRecord.ID)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func DelExport(id string)(err error){
	if id==""||common.HasSpecialCharacter(id){
		err=fmt.Errorf("Del Export id=%v",id)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Delete from export_records where eid=$1")
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

func SearchExportList(id string)(exportRecord *ExportRecord,err error){
	if id==""||common.HasSpecialCharacter(id){

	}
	exportRecord=new(ExportRecord)
	row:=db.QueryRow("Select * from export_records_view where eid=$1",id)
	if row==nil{
		return
	}
	err=row.Scan(&exportRecord.ID,&exportRecord.GID,exportRecord.GoodName,exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&exportRecord.Edate)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}
