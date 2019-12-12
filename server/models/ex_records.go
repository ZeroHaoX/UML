package models

import (
	"github.com/astaxie/beego/logs"
	"../common"
	"fmt"
	"time"
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
	Price float64	`json:"eprice"`
	TotalPrice float64	`json:"etotalprice"`
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

func ExportList(page int,pageSize int,orderBy string)(exportList []ExportRecord,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf("Select eid,imdate,gname,shipper,ecount,eprice,etotalprice,buyer,bphone,detial,edate,profit from export_records order by edate %v limit %v offset %v ",orderBy,pageSize,(page-1)*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var exportRecord ExportRecord
		var elocalTime time.Time
		var imlocalTime time.Time
		err=rows.Scan(&exportRecord.ID,&imlocalTime,&exportRecord.GoodName,&exportRecord.Shipper,&exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&elocalTime,&exportRecord.Profit)
		if err!=nil{
			logs.Error(err)
			return
		}
		exportRecord.Edate=elocalTime.Format("2006-01-02")
		exportRecord.ImDate=imlocalTime.Format("2006-01-02")
		exportList=append(exportList,exportRecord)
	}
	return
}

func Exoprt(exportRecord *ExportRecord)(err error){
	if exportRecord==nil{
		err=fmt.Errorf("Export exportRecort is error:exportRecord=%v",err)
		logs.Error(err)
		return
	}
	logs.Debug(exportRecord)
	stmt,err:=db.Prepare("Insert into export_records(eid,imdate,gname,shipper,ecount,eprice,etotalprice,buyer,bphone,detial) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(exportRecord.ID,exportRecord.ImDate,exportRecord.GoodName,exportRecord.Shipper,exportRecord.Count,exportRecord.Price,exportRecord.TotalPrice,exportRecord.Buyer,exportRecord.Phone,exportRecord.Detial)
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

func SearchExport(id string)(exportRecords []ExportRecord,err error){
	if id==""||common.HasSpecialCharacter(id){
		err=fmt.Errorf("id 有误")
		logs.Error(err)
		return
	}
	// exportRecord=new(ExportRecord)
	rows,err:=db.Query("Select eid,imdate,gname,shipper,ecount,eprice,etotalprice,buyer,bphone,detial,edate,profit from export_records where eid=$1 ",id)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var elocalTime time.Time
		var imlocalTime time.Time
		var exportRecord ExportRecord
		err=rows.Scan(&exportRecord.ID,&imlocalTime,&exportRecord.GoodName,&exportRecord.Shipper,&exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&elocalTime,&exportRecord.Profit)
		if err!=nil{
			logs.Error(err)
			return
		}
		exportRecord.Edate=elocalTime.Format("2006-01-02")
		exportRecord.ImDate=imlocalTime.Format("2006-01-02")
		exportRecords=append(exportRecords,exportRecord)
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
		var elocalTime time.Time
		var imlocalTime time.Time
		err=rows.Scan(&exportRecord.ID,&imlocalTime,&exportRecord.GoodName,&exportRecord.Shipper,&exportRecord.Count,&exportRecord.Price,&exportRecord.TotalPrice,&exportRecord.Buyer,&exportRecord.Phone,&exportRecord.Detial,&elocalTime,&exportRecord.Profit)
		if err!=nil{
			logs.Error(err)
			return
		}
		exportRecord.Edate=elocalTime.Format("2006-01-02")
		exportRecord.ImDate=imlocalTime.Format("2006-01-02")
		exportRecords=append(exportRecords,exportRecord)
	}
	return
}


func SearchExportNumber()(length int,err error){
	rows,err:=db.Query("Select count(eid) from export_records")
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