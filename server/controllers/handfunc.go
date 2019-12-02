package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	// "context"
	"../common"
	"../models"
	"fmt"
	"errors"
)

//登录
func LoginHand(w http.ResponseWriter,r *http.Request){
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	username,ok:=data["userName"].(string)
	if !ok{
		err=fmt.Errorf("get userName error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if username==""||common.HasSpecialCharacter(username){
		err=fmt.Errorf("LoginHand get username=%v",username)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	password,ok:=data["password"].(string)
	if !ok{
		err=fmt.Errorf("get password error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if password==""||common.HasSpecialCharacter(password){
		err=fmt.Errorf("LoginHand get password=%v",password)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	logs.Info("能够运行到这里")
	user,ok,err:=models.CheckUserMes(username,password)
	logs.Info("user=",user)
	logs.Info("ok",ok)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		SuccessResponse(w,r,[]models.Permission{},"用户不存在",0)
		return
	}

	/*搜权限*/
	userMes,err:=models.SearchUser(username)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if userMes==nil{
		err=fmt.Errorf("search permission return nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}

	if user==nil{
		err=fmt.Errorf("check user return nil user:%v",user)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}	
	tokerString,err:=common.CreateToken(user.Name,user.Role)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if tokerString==""{
		err=fmt.Errorf("create tokerString error:tokerString=%v",tokerString)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	cookie:=&http.Cookie{
		Name:"token",
		Value:tokerString,
	}
	http.SetCookie(w,cookie)
	
	SuccessResponse(w,r,userMes,"登录成功",0)

}

//注册
func RegisteHand(w http.ResponseWriter,r *http.Request){
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("get data error:data is nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	var user models.User
	user.Name,ok=data["username"].(string)
	if !ok{
		err=fmt.Errorf("get username error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if user.Name==""||common.HasSpecialCharacter(user.Name){
		err=fmt.Errorf("get userName error:userName=%v",user.Name)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.Password,ok=data["password"].(string)
	if !ok{
		err=fmt.Errorf("get password error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.Password==""||common.HasSpecialCharacter(user.Password){
		err=fmt.Errorf("get password error:password=%v",user.Password)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.Phone,ok=data["phone"].(string)
	if !ok{
		err=fmt.Errorf("get phone error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.Phone==""||common.HasSpecialCharacter(user.Phone){
		err=fmt.Errorf("get phone error:phone=%v",user.Phone)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.ActualName,ok=data["actualname"].(string)
	if !ok{
		err=fmt.Errorf("get actualName error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.ActualName==""||common.HasSpecialCharacter(user.ActualName){
		err=fmt.Errorf("get actualName error:actualName=%v",user.ActualName)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	u,err:=models.SearchUserByName(user.Name)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if u!=nil{
		SuccessResponse(w,r,nil,"用户已存在!",0)
	}
	err=models.InsertUser(&user)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}

	SuccessResponse(w,r,nil,"注册成功!",0)
	
}

//出货
func ExportHand(w http.ResponseWriter,r *http.Request){
	ok,err:=CheckPermission(r,"出货管理")
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("ExportHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("ExportHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	var exportRecord models.ExportRecord
	exportRecord.ID,ok=data["id"].(string)
	if !ok{
		err=fmt.Errorf("get export id error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.ID==""||common.HasSpecialCharacter(exportRecord.ID){
		err=fmt.Errorf("get export id=%v",exportRecord.ID)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.GID,ok=data["gid"].(string)
	if !ok{
		err=fmt.Errorf("ExportHand get gid error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.GID==""||common.HasSpecialCharacter(exportRecord.GID){
		err=fmt.Errorf("ExportHand get gid=%v",exportRecord.GID)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.GoodName,ok=data["gname"].(string)
	if !ok{
		err=fmt.Errorf("ExportHand get gname error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.GoodName==""||common.HasSpecialCharacter(exportRecord.GoodName){
		err=fmt.Errorf("ExportHand get gname=%v",exportRecord.GoodName)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.Count,ok=data["count"].(int)
	if !ok{
		err=fmt.Errorf("ExportHand get count error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.Count<0{
		err=fmt.Errorf("ExportHand get count<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.Price,ok=data["price"].(float32)
	if !ok{
		err=fmt.Errorf("ExportHand get price error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.Price<0{
		err=fmt.Errorf("ExportHand get price<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.TotalPrice,ok=data["totalprice"].(float32)
	if !ok{
		err=fmt.Errorf("ExportHand get totalprice error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.TotalPrice<0{
		err=fmt.Errorf("ExportHand get totalprice<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.Buyer,ok=data["buyer"].(string)
	if !ok{
		err=fmt.Errorf("ExportHand get buyer error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.Buyer==""||common.HasSpecialCharacter(exportRecord.Buyer){
		err=fmt.Errorf("ExportHand get buyer=%v",exportRecord.Buyer)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.Phone,ok=data["phone"].(string)
	if !ok{
		err=fmt.Errorf("ExportHand get phone error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if exportRecord.Phone==""||common.HasSpecialCharacter(exportRecord.Phone){
		err=fmt.Errorf("ExportHand get phone=%v",exportRecord.Phone)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	exportRecord.Detial,ok=data["detial"].(string)
	if !ok{
		err=fmt.Errorf("ExportHand get detial error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if common.HasSpecialCharacter(exportRecord.Detial){
		err=fmt.Errorf("ExportHand get detial error:detial=%v",exportRecord.Detial)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	err=models.Exoprt(&exportRecord)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,nil,"出货成功！",0)
}

//进货
func ImportHand(w http.ResponseWriter,r *http.Request){
	ok,err:=CheckPermission(r,"进货管理")
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("ImportHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("ImportHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	var importRecord models.ImportRecord
	importRecord.ID,ok=data["id"].(string)
	if !ok{
		err=fmt.Errorf("ImportHand get import id error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.ID==""||common.HasSpecialCharacter(importRecord.ID){
		err=fmt.Errorf("ImportHand get import id=%v",importRecord.ID)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	importRecord.GoodName,ok=data["gname"].(string)
	if !ok{
		err=fmt.Errorf("ImportHand get import good name error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.GoodName==""||common.HasSpecialCharacter(importRecord.GoodName){
		err=fmt.Errorf("ImportHand get import good name=%v",importRecord.GoodName)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	importRecord.Price,ok=data["price"].(float32)
	if !ok{
		err=fmt.Errorf("ImportHand get import price error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.Price<0{
		err=fmt.Errorf("ImportHand get import price<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	importRecord.Count,ok=data["count"].(int)
	if !ok{
		err=fmt.Errorf("ImportHand get import count error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.Count<0{
		err=fmt.Errorf("ImportHand get import count<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	importRecord.TotalPrice,ok=data["totalprice"].(float32)
	if !ok{
		err=fmt.Errorf("ImportHand get import totalprice error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.TotalPrice<0{
		err=fmt.Errorf("ImportHand get import totalprice<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	importRecord.Shipper,ok=data["shipper"].(string)
	if !ok{
		err=fmt.Errorf("ImportHand get shipper error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if importRecord.Shipper==""||common.HasSpecialCharacter(importRecord.Shipper){
		// err=fmt.Errorf()
	}
}

//商品列表
func GoodsListHand(w http.ResponseWriter,r *http.Request){
	ok,err:=CheckPermission(r,"商品管理")
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("GoodsListHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("GoodsListHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	orderBy,ok:=body["orderby"].(string)
	if !ok{
		err=fmt.Errorf("GoodsListHand get orderby error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("GoodsListHand get orderby=%v",orderBy)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	page,ok:=body["page"].(int)
	if !ok{
		err=fmt.Errorf("GoodsListHand get page error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if page<0{
		err=fmt.Errorf("GoodsListHand get page<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	pageSize,ok:=body["pagesize"].(int)
	if !ok{
		err=fmt.Errorf("get pagesize error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if pageSize<0{
		err=fmt.Errorf("get pagesize<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	goodsList,err:=models.GoodsList(page,pageSize,orderBy)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if goodsList==nil{
		err=fmt.Errorf("GoodsListHand goodsList is nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,goodsList,"查询成功！",pageSize)
}

//更新商品信息
func UpdateGoodsMes(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("UpdateGoodsMes get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("UpdateGoodsMes can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"商品信息管理")
	if err!=nil{
		err=fmt.Errorf("UpdateGoodsMes check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	var good models.Good
	good.ID,ok=data["gid"].(int)
	if !ok{
		err=fmt.Errorf("get gid error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.ID<0{
		err=fmt.Errorf("get gid<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	good.Name,ok=data["gname"].(string)
	if !ok{
		err=fmt.Errorf("get gname error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.Name==""||common.HasSpecialCharacter(good.Name){
		err=fmt.Errorf("UpdateGoodsMes get gname=%v",good.Name)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	good.Shipper,ok=data["shipper"].(string)
	if !ok{
		err=fmt.Errorf("UpdateGoodsMes get shipper error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.Shipper==""||common.HasSpecialCharacter(good.Shipper){
		err=fmt.Errorf("UpdateGoodsMes get shipper=%v",good.Shipper)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	good.Count,ok=data["count"].(int)
	if !ok{
		err=fmt.Errorf("UpdateGoodsMes get count error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.Count<0{
		err=fmt.Errorf("UpdateGoodsMes get count<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	good.Price,ok=data["price"].(float32)
	if !ok{
		err=fmt.Errorf("UpdateGoodsMes get Price error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.Price<0{
		err=fmt.Errorf("UpdateGoodsMes get Price<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	good.ImPrice,ok=data["imprice"].(float32)
	if !ok{
		err=fmt.Errorf("UpdateGoodsMes get ImPrice error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if good.ImPrice<0{
		err=fmt.Errorf("UpdateGoodsMes get ImPrice<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	err=models.UpdateGoodMes(&good)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,nil,"修改成功",0)
}

//用户列表
func UsersListHand(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("UserListHand get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("UserListHand can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"员工信息管理")
	if err!=nil{
		err=fmt.Errorf("UserListHand check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("UserListHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("UserListHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	orderBy,ok:=body["orderby"].(string)
	if !ok{
		err=fmt.Errorf("UserListHand get orderby error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if orderBy==""||common.HasSpecialCharacter(orderBy){
		err=fmt.Errorf("UserListHand get orderby=%v",orderBy)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	page,ok:=body["page"].(int)
	if !ok{
		err=fmt.Errorf("UserListHand get page error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if page<0{
		err=fmt.Errorf("UserListHand get page<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	pageSize,ok:=body["pagesize"].(int)
	if !ok{
		err=fmt.Errorf("UserListHand get pagesize error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if pageSize<0{
		err=fmt.Errorf("UserListHand get pagesize<0")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	users,err:=models.ShowUserList(page,pageSize,orderBy)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,users,"查询成功！",pageSize)
}

//更新用户信息
func UpdateUserHand(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("UpdateUserHand get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("UpdateUserHand can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"员工信息更改")
	if err!=nil{
		err=fmt.Errorf("UpdateUserHand check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("UpdateUserHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("UpdateUserHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	var user models.User
	user.Name,ok=data["username"].(string)
	if !ok{
		err=fmt.Errorf("get username error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if user.Name==""||common.HasSpecialCharacter(user.Name){
		err=fmt.Errorf("get userName error:userName=%v",user.Name)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.Password,ok=data["password"].(string)
	if !ok{
		err=fmt.Errorf("get password error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.Password==""||common.HasSpecialCharacter(user.Password){
		err=fmt.Errorf("get password error:password=%v",user.Password)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.Phone,ok=data["phone"].(string)
	if !ok{
		err=fmt.Errorf("get phone error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.Phone==""||common.HasSpecialCharacter(user.Phone){
		err=fmt.Errorf("get phone error:phone=%v",user.Phone)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	user.ActualName,ok=data["actualname"].(string)
	if !ok{
		err=fmt.Errorf("get actualName error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
	}
	if user.ActualName==""||common.HasSpecialCharacter(user.ActualName){
		err=fmt.Errorf("get actualName error:actualName=%v",user.ActualName)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}

	err=models.UpdateUser(&user)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	
	SuccessResponse(w,r,nil,"修改成功！",0)
}

//月结
func MonthlyHand(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("MonthlyHand get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("MonthlyHand can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"月结管理")
	if err!=nil{
		err=fmt.Errorf("MonthlyHand check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("MonthlyHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("MonthlyHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	year,ok:=data["year"].(int)
	if !ok{
		err=fmt.Errorf("MonthlyHand get year error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	day,ok:=data["day"].(int)
	if !ok{
		err=fmt.Errorf("MonthlyHand get day error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !common.CheckDate(year,day){
		err=fmt.Errorf("MonthlyHand get year=%v,day=%v",year,day)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	account,err:=models.GetAccount(year,day)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}

	SuccessResponse(w,r,account,"查询成功！",1)
}

//删除商品信息
func DelGoodHand(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("DelGoodHand get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("DelGoodHand can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"商品管理")
	if err!=nil{
		err=fmt.Errorf("DelGoodHand check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("DelGoodHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("DelGoodHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	gid,ok:=data["gid"].(string)
	if !ok{
		err=fmt.Errorf("DelGoodHand get gid error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if gid==""||common.HasSpecialCharacter(gid){
		err=fmt.Errorf("DelGoodHand get gid=%v",gid)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	err=models.DelGoodMes(gid)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,nil,"删除成功！",0)
}

//删除用户信息
func DelUserHand(w http.ResponseWriter,r *http.Request){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err:=fmt.Errorf("DelUserHand get role form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if role==""{
		err:=fmt.Errorf("DelUserHand can't get role form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	ok,err:=models.CheckPermission(role,"员工信息管理")
	if err!=nil{
		err=fmt.Errorf("DelUserHand check permission error:%v",err)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if !ok{
		ErrorResponse(w,r,errors.New("权限不足！"),403)
		return
	}
	body,err:=ReadBodyData(r)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	data,ok:=body["data"].(map[string]interface{})
	if !ok{
		err=fmt.Errorf("DelUserHand get data error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if data==nil{
		err=fmt.Errorf("DelUserHand get data nil")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	userName,ok:=data["username"].(string)
	if !ok{
		err=fmt.Errorf("DelUserHand get username error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if userName==""||common.HasSpecialCharacter(userName){
		err=fmt.Errorf("DelGoodHand get userName=%v",userName)
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	err=models.DelUser(userName)
	if err!=nil{
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	SuccessResponse(w,r,nil,"删除成功！",0)
}

//查找商品信息
func SearchGoodsHand(w http.ResponseWriter,r *http.Request){

}

//查找出货信息
func SearchExportHand(w http.ResponseWriter,r *http.Request){

}

//查找进货信息
func SearchImportHand(w http.ResponseWriter,r *http.Request){

}

//获取用户信息接口
func GetUserMes(w http.ResponseWriter,r *http.Request){
	userName,ok:=r.Context().Value("username").(string)
	if !ok{
		err:=fmt.Errorf("DelUserHand get username form context error")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}
	if userName==""{
		err:=fmt.Errorf("DelUserHand can't get username form context")
		logs.Error(err)
		ErrorResponse(w,r,err,500)
		return
	}

	user,err:=models.SearchUser(userName)
	if err!=nil{
		ErrorResponse(w,r,err,500)
		return
	}
	
	SuccessResponse(w,r,user,"信息搜索成功！",0)
}