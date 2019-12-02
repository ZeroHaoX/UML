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
	Name string	`json:"name"`
	Password string	`json:"password"`
	ActualName string `json:"actualname"`
	Phone string `json:"phone"`
	Role string	`json:"role"`
}
type User1 struct{
	// Uno string `json:"uno"`
	UserName string	`json:"userName"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Role string	`json:"role"`
	Password string	`json:"password"`
}

type UserView struct{
	User
	Permissions []Permission
}

//按名字搜索某用户
func SearchUserByName(userName string)(user *User,err error){
	if userName==""||common.HasSpecialCharacter(userName){
		err=fmt.Errorf("userName is error userName=%v",userName)
		logs.Error(err)
		return
	}
	user=new(User)
	row:=db.QueryRow("select * from users where username=$1",userName)
	if row==nil{
		return
	}else{
		err=row.Scan(&user.Name,&user.ActualName,user.Phone,&user.Role)
		if err!=nil{
			err=fmt.Errorf("scan user error:%v",err)
			logs.Error(err)
			return
		}
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
func ShowUserList(page int,pageSize int,orderBy string)(users []*User,err error){
	if page<0||pageSize<0||orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("query parameters has error:page=%v pageSize=%v orderby=%v",page,pageSize,orderBy)
		logs.Error(err)
		return
	}
	queryString:=fmt.Sprintf("Select * from users order by %v limit %v offset %v",orderBy,pageSize,page*pageSize)
	rows,err:=db.Query(queryString)
	if err!=nil{
		err=fmt.Errorf("query users error:%v",err)
		logs.Error(err)
		return
	}
	for rows.Next(){
		var user User
		err=rows.Scan(&user.Name,&user.ActualName,&user.Phone,&user.Role)
		if err!=nil{
			err=fmt.Errorf("row scan error:%v",err)
			logs.Error(err)
			return
		}
		users=append(users,&user)
	}
	return
}

//检验信息
func CheckUserMes(userName,password string)(user *User1,ok bool,err error){
	user = new(User1)
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
	row:=db.QueryRow("select username,name,phone,role,password from users where username=$1 and password=$2",userName,password)
	logs.Info("row",row)
	if row==nil{
		ok=false
	}else{
		fmt.Printf("%T\n",user.UserName)
		row.Scan(&user.UserName,&user.Name,&user.Phone,&user.Role,&user.Password)
		logs.Info("UserName=",user.UserName,user.Name)
		if(user.UserName==""){
			logs.Info("用户不存在")
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
	stmt,err:=db.Prepare("Insert into users values($1,$2,$3,$4,'普通用户')")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(user.Name,user.ActualName,user.Password,user.Phone)
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
	stmt,err:=db.Prepare("Update users set uname=$1,actualname=$2,password=$3,phone=$4,role=$5")
	if err!=nil{
		logs.Error(err)
		return
	}
	_,err=stmt.Exec(user.Name,user.ActualName,user.Password,user.Phone,user.Role)
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
	stmt,err:=db.Prepare("Delete from users where username=$1")
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
	rows,err:=db.Query("Select * from user_view where actualname=$1 limit $2 offset $3",actualName,pageSize,page*pageSize)
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