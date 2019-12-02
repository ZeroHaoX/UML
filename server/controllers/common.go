package controllers

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	// "context"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"../common"
	"context"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"../models"
)

//后端响应数据通信协议
type ReplyProto struct {
	Status    int         `json:"status"` //状态 0正常，小于0出错，大于0可能有问题
	Msg       string      `json:"msg"`    //状态信息
	Data      interface{} `json:"data"`
	API       string      `json:"API"`        //api接口
	Method    string      `json:"method"`     //post,put,get,delete
	RowCount  int         `json:"rowCount"`   //Data若是数组，算其长度
	// Time      int64       `json:"time"`       //请求响应时间，毫秒
	// CheckTime int64       `json:"check_time"` //检测时间，毫秒
}

//前端请求数据通讯协议
type ReqProto struct {
	Action   string      `json:"action"` //请求类型GET/POST/PUT/DELETE
	Data     interface{} `json:"data"`   //请求数据
	Sets     []string    `json:"sets"`
	OrderBy  string      `json:"orderBy"`  //排序要求
	Filter   string      `json:"filter"`   //筛选条件
	Page     int         `json:"page"`     //分页
	PageSize int         `json:"pageSize"` //分页大小
}


// type Context struct{
// 	Url string
// 	r *http.Request
// 	w http.ResponseWriter
// }

//检查权限
func CheckPermission(r *http.Request,permission string)(ok bool,err error){
	role,ok:=r.Context().Value("role").(string)
	if !ok{
		err=fmt.Errorf("get role form context error")
		logs.Error(err)
		return
	}
	if role==""{
		err=fmt.Errorf("get role form context nil")
		logs.Error(err)
		return
	}
	ok,err=models.CheckPermission(role,permission)
	if err!=nil{
		err=fmt.Errorf("GoodsList check permission error:%v",err)
		logs.Error(err)
		return
	}
	return
}

//读取body post
func ReadBodyData(r *http.Request)(data map[string]interface{},err error){
	if r==nil{
		err=fmt.Errorf("ReadData r is erro:%v",err)
		logs.Error(err)
		return
	}
	body,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		logs.Error(err)
		return
	}
	if len(body)==0{
		err=fmt.Errorf("Read request body nil")
		logs.Error(err)
		return
	}
	err=json.Unmarshal(body,&data)
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

func ReadUrl(){
	
}

//token校验
func CheckTokenPreHand(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if err:=recover();err!=nil{
				logs.Error(err)
				http.Error(w,"发生未知错误",500)
			}
		}()
		if r.Referer()==""{
			err:=fmt.Errorf("Referer is nil someone attacking")
			ErrorResponse(w,r,err,500)
			return
		}
		cookie,err:=r.Cookie("token")
		if err!=nil{
			err=fmt.Errorf("get token from cookie error:%v",err)
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		if cookie==nil{
			err=fmt.Errorf("get token from cookie nil")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		tokenString:=cookie.Value
		if tokenString==""{
			err=fmt.Errorf("parse tokenString=%v",tokenString)
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		token,err:=common.ParseToken(tokenString)
		if err!=nil{
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		if token==nil{
			err=fmt.Errorf("parse token get nil")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		if !token.Valid{
			ErrorResponse(w,r,errors.New("权限不足！"),403)
			return
		}
		logs.Debug("刷新token")
		tokenString,isRefresh,err:=common.IsRefreshToken(token.Claims)
		if err!=nil{
			logs.Error(err)
		}
		if isRefresh==true{
			cookie:=&http.Cookie{
				Name:"token",
				Value:tokenString,
			}
		http.SetCookie(w,cookie)
		}

		userName,ok:=token.Claims.(jwt.MapClaims)["uname"].(string)
		if !ok{
			err=fmt.Errorf("get user name from token error")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		if userName==""{
			err=fmt.Errorf("user name in token is nil")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		role,ok:=token.Claims.(jwt.MapClaims)["role"].(string)
		if !ok{
			err=fmt.Errorf("get role from token error")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		if role==""{
			err=fmt.Errorf("role in token is nil")
			logs.Error(err)
			ErrorResponse(w,r,err,500)
			return
		}
		ctx:=context.WithValue(r.Context(),"username",userName)
		ctx=context.WithValue(ctx,"role",role)
		r=r.WithContext(ctx)
		h(w,r)
	}
}

//合法校验
func PreHand(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if err:=recover();err!=nil{
				logs.Error(err)
				http.Error(w,"发生未知错误",500)
			}
		}()
		
		if r.Referer()==""{
			err:=fmt.Errorf("Referer is nil someone attacking")
			ErrorResponse(w,r,err,403)
			return
		}
		if r.RequestURI!="getmes"{
			if r.Method!="POST"{
				err:=fmt.Errorf("Method is error")
				ErrorResponse(w,r,err,403)
				return
			}
		}


		// common.ParseToken()
		h(w,r)
	}
}

func ErrorResponse(w http.ResponseWriter,r *http.Request,err error,code int){
	if w==nil||r==nil{
		err=fmt.Errorf("ErrorResponse get w=%v,r=%v",w,r)
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
	if err==nil{
		err=fmt.Errorf("ErrorResponse get error=%v",err)
		logs.Error(err)
		http.Error(w,err.Error(),500)
	}
	if code!=500 && code!=403{
		err=fmt.Errorf("ErrorResponse get code=%v",code)
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
	resp := ReplyProto{}
	resp.Status = -1
	resp.Msg = err.Error()
	resp.Data = nil
	resp.API = r.RequestURI
	resp.Method = r.Method
	resp.RowCount = 0

	respData,err:=json.Marshal(resp)
	if err!=nil{
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}

	w.WriteHeader(code)
	_,err=w.Write(respData)
	if err!=nil{
		err=fmt.Errorf("w write error:%v",err)
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
}

func SuccessResponse(w http.ResponseWriter,r *http.Request,data interface{},msg string,count int){
	if w == nil || r == nil {
		err := fmt.Errorf("arguments can not be a nil value")
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
	if count < 0 {
		err := fmt.Errorf("rowCount must be positive")
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
	resp := ReplyProto{}
	resp.Status = 0
	resp.Msg = msg
	resp.Data = data
	resp.API = r.RequestURI
	resp.Method = r.Method
	resp.RowCount = count

	respData,err:=json.Marshal(resp)
	if err!=nil{
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}

	_,err=w.Write(respData)
	if err!=nil{
		err=fmt.Errorf("w write error:%v",err)
		logs.Error(err)
		http.Error(w,err.Error(),500)
		return
	}
}