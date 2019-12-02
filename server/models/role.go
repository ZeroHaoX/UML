package models

import (
	"github.com/astaxie/beego/logs"
	"../common"
	"fmt"
)

type Permission struct{
	Name string `json:"name"`
	API string	`json:"api"`
}


func PermissionList(role string)(result []*Permission,err error){
	if role==""||common.HasSpecialCharacter(role){
		err=fmt.Errorf("PermissionList role is error:role=%v",role)
		logs.Error(err)
		return
	}
	rows,err:=db.Query("select * from role_permission where role=$1",role)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var p Permission
		err=rows.Scan(&p.Name,&p.API)
		if err!=nil{
			logs.Error(err)
			return
		}
		result=append(result,&p)
	}
	return
}

func CheckPermission(role string,permission string)(ok bool,err error){
	if role==""||common.HasSpecialCharacter(role){
		err=fmt.Errorf("CheckPermission role is error:role=%v",role)
		logs.Error(err)
		return
	}
	if permission==""||common.HasSpecialCharacter(permission){
		err=fmt.Errorf("CheckPermission role is error:permission=%v",permission)
		logs.Error(err)
		return
	}
	row:=db.QueryRow("select * from role_permission where role=$1 and pname=$2",role,permission)
	if row!=nil{
		ok=true
	}
	return
}