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
	ID string	`json:"eid"`
	Edate string	`json:"edate"`
	// GID string	`json:"gid"`
	ImDate string 	`json:"imdate"`
	Shipper string	`json:"shipper"`
	GoodName string	`json:"gname"`
	Count int		`json:"ecount"`
	Price float32	`json:"eprice"`
	TotalPrice float32	`json:"etotalprice"`
	Buyer string	`json:"buyer"`
	Phone string	`json:"bphone"`
	Detial string	`json:"detial"`
	Profit float64	`json:"profit"`
}

func nextMonth(year int,month int)(timeString string){
	var m string
	if month<10{
		m="0"+strconv.Itoa(month)
	}else if month >=10 && month<=11{
		m=strconv.Itoa(month)
	}
	if month==12{
		year++
		m="0"+strconv.Itoa(1)
	}
	// if day<10{
	// 	d="0"+strconv.Itoa(month)
	// }
	timeString=fmt.Sprintf("%v-%v-01",year,m)
	logs.Debug(timeString)
	return
}

func ExportList(year int,month int)(exportList []*ExportRecord,err error){
	if !common.CheckDate(year,month){
		err=fmt.Errorf("year or month is error:year=%v month=%v",year,month)
		logs.Error(err)
		return
	}
	start:=nextMonth(year,month)
	end:=nextMonth(year,month+1)
	rows,err:=db.Query("select * from export_records_view(id,gid,gname,eprice,etotalprice,ecount,detial,buyer,phone,edate) where edate>=$1::timestamp and edate<$2::timestamp",start,end)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var exoprtRecord ExportRecord
		err=rows.Scan(&exoprtRecord.ID,&exoprtRecord.ImDate,&exoprtRecord.GoodName,&exoprtRecord.Price,&exoprtRecord.TotalPrice,&exoprtRecord.Count,&exoprtRecord.Detial,&exoprtRecord.Buyer,&exoprtRecord.Phone,exoprtRecord.Edate)
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
	stmt,err:=db.Prepare("Insert into export_records(eid,imdate,gname,shipper,ecount,eprice,etotalprice,buyer,bphone,detial,edate) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(&exportRecord.ID,&exportRecord.ImDate,exportRecord.GoodName,&exportRecord.Shipper,&exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&exportRecord.Edate)
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
	_,err=stmt.Exec(exportRecord.ImDate,exportRecord.GoodName,exportRecord.Price,exportRecord.Count,exportRecord.TotalPrice,exportRecord.Buyer,exportRecord.Phone,exportRecord.Detial,exportRecord.ID)
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
	err=row.Scan(&exportRecord.ID,&exportRecord.ImDate,exportRecord.GoodName,exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&exportRecord.Edate)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func SearchExportListByTime(year int,month int)(exportRecords []ExportRecord,err error){
	nextDate:=nextMonth(year,month)
	nowDate:=timeToString(year,month)
	rows,err:=db.Query("Select eid,imdate,gname,shipper,ecount,eprice,etotalprice,buyer,bphone,detial,edate,profit from export_records where edate>=$1 and edate<$2 order by eid asc",nowDate,nextDate)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var exportRecord ExportRecord
		err=rows.Scan(&exportRecord.ID,&exportRecord.ImDate,&exportRecord.GoodName,&exportRecord.Shipper,&exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&exportRecord.Edate,&exportRecord.Profit)
		if err!=nil{
			logs.Error(err)
			return
		}
		exportRecords=append(exportRecords,exportRecord)
	}
	return
}