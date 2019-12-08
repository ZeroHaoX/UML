package models

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	// "regexp"
	// "reflect"
	"../common"
)

type User struct{
	// Uno string `json:"uno"`
	Name string	`json:"username"`
	Password string	`json:"password"`
	ActualName string `json:"actualname"`
	Phone string `json:"phone"`
	Role string	`json:"role"`
}

type UserView struct{
	User
	Permissions []Permission	`json:"permissions"`
}

//按名字搜索某用户
func SearchUserByUserName(userName string)(user *User,err error){
	if userName==""||common.HasSpecialCharacter(userName){
		err=fmt.Errorf("userName is error userName=%v",userName)
		logs.Error(err)
		return
	}
	user=new(User)
	rows,err:=db.Query("select * from users where username=$1",userName)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		err=rows.Scan(&user.Name,&user.ActualName,user.Phone,&user.Role)
		if err!=nil{
			err=fmt.Errorf("scan user error:%v",err)
			logs.Error(err)
			return
		}
	}

	return
}


func SearchUserByName(name string)(users []User,err error){
	if name==""||common.HasSpecialCharacter(name){
		err=fmt.Errorf("userName is error name=%v",name)
		logs.Error(err)
		return
	}
	// user=new(User)
	rows,err:=db.Query("select name,username,phone,role from users where name=$1",name)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		var user User
		err=rows.Scan(&user.Name,&user.ActualName,&user.Phone,&user.Role)
		if err!=nil{
			err=fmt.Errorf("scan user error:%v",err)
			logs.Error(err)
			return
		}
		users=append(users,user)
	}
	return
}

// //按员工号查询
// func SearchUserByUno(userName string)(user *User,err error){
// 	if uno==""||common.HasSpecialCharacter(uno){
// 		err=fmt.Errorf("userName is error uno=%v",uno)
// 		logs.Error(err)
// 		return
// 	}
// 	user=new(User)
// 	row:=db.QueryRow("select * from users where username=$1",uno)
// 	if row==nil{
// 		return
// 	}else{
// 		err=row.Scan(&user.Name,&user.ActualName,user.Phone,&user.Role)
// 		if err!=nil{
// 			err=fmt.Errorf("scan user error:%v",err)
// 			logs.Error(err)
// 			return
// 		}
// 	}
// 	return
// }

//展示用户列表
func ShowUserList(page int,pageSize int,orderBy string)(users []User,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf("Select username,name,phone,role,password from users order by username %v limit %v offset %v",orderBy,pageSize,(page-1)*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		err=fmt.Errorf("query users error:%v",err)
		logs.Error(err)
		return
	}
	for rows.Next(){
		var user User
		err=rows.Scan(&user.Name,&user.ActualName,&user.Phone,&user.Role,&user.Password)
		if err!=nil{
			err=fmt.Errorf("row scan error:%v",err)
			logs.Error(err)
			return
		}
		// logs.Debug(user)
		users=append(users,user)
	}
	return
}

//检验信息
func CheckUserMes(userName,password string)(user *User,ok bool,err error){
	if userName==""||common.HasSpecialCharacter(userName){
		err=fmt.Errorf("username is error:userName=%v",userName)
		logs.Error(err)
		return
	}
	if common.HasSpecialCharacter(password)||common.HasSpecialCharacter(password){
		err=fmt.Errorf("password is error:password=%v",password)
		logs.Error(err)
		return
	}
	user=new(User)
	rows,err:=db.Query("select * from users where username=$1 and password=$2",userName,password)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		err=rows.Scan(&user.Name,&user.ActualName,&user.Phone,&user.Role,&user.Password)
		if err!=nil{
			logs.Error(err)
			return
		}
		ok=true
	}
	return
}

//用户信息插入
func InsertUser(user *User)(err error){
	if user==nil{
		err=fmt.Errorf("InsertUser user is error:user=%v",user)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Insert into users(username,name,phone,password,role) values($1,$2,$3,$4,'普通用户')")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(user.Name,user.ActualName,user.Phone,user.Password)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func UpdateUser(user *User)(err error){
	if user==nil{
		err=fmt.Errorf("UpdateUser user is nil")
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Update users set name=$1,password=$2,phone=$3,role=$4 where username=$5")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(user.ActualName,user.Password,user.Phone,user.Role,user.Name)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func DelUser(username string)(err error){
	if username==""||common.HasSpecialCharacter(username){
		err=fmt.Errorf("DelUser username=%v",username)
		logs.Error(err)
		return
	}
	stmt,err:=db.Prepare("Delete from users_message where username=$1")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(username)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

//用户信息查询
func SearchUser(username string)(uv *UserView,err error){
	if username==""||common.HasSpecialCharacter(username){
		err=fmt.Errorf("username is error:username=%v",username)
		logs.Error(err)
		return
	}
	uv=new(UserView)
	rows,err:=db.Query("Select * from user_view where username=$1",username)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		// var userName string
		// var actualName string
		// var phone string
		// var role string
		var permission Permission
		err=rows.Scan(&uv.Name,&uv.ActualName,&uv.Phone,&uv.Role,&permission.Name,&permission.API)
		if err!=nil{
			logs.Error(err)
			return
		}
		uv.Permissions=append(uv.Permissions,permission)
	}
	return
}


func SearchUsers(actualName string,page int,pageSize int)(uv *UserView,err error){
	if actualName==""||common.HasSpecialCharacter(actualName){
		err=fmt.Errorf("actualName is error:username=%v",actualName)
		logs.Error(err)
		return
	}
	if page<0{
		err=fmt.Errorf("page is error:page=%v",page)
		logs.Error(err)
		return
	}
	if pageSize<0{
		err=fmt.Errorf("page is error:pageSize=%v",pageSize)
		logs.Error(err)
		return
	}
	rows,err:=db.Query("Select * from user_view where name=$1 limit $2 offset $3",actualName,pageSize,page*pageSize)
	if err!=nil{
		logs.Error(err)
		return
	}
	for rows.Next(){
		// var userName string
		// var actualName string
		// var phone string
		// var role string
		var permission Permission
		err=rows.Scan(&uv.Name,&uv.ActualName,&uv.Phone,&uv.Role,&permission.Name,&permission.API)
		if err!=nil{
			logs.Error(err)
			return
		}
		uv.Permissions=append(uv.Permissions,permission)
	}
	return
}

func SearchUserNumber()(length int,err error){
	rows,err:=db.Query("Select count(username) from users")
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