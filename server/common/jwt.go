package common

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	// "net/http"
)

var SecretKey="secretkey"

type MTokenClaim struct{
	jwt.StandardClaims
	UserName string	`json:"username"`
	Role string	`json:"role"`
}

//生成token
func CreateToken(uname string,role string)(tokenString string,err error){
	if uname==""||role==""{
		err=fmt.Errorf("create uname=%v",uname)
		logs.Error(err)
		return
	}
	claims:=&MTokenClaim{
		jwt.StandardClaims{
			ExpiresAt:int64(time.Now().Add(72*time.Hour).Unix()),
		},
		uname,
		role,
	}
	// logs.Debug(claims)
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err=token.SignedString([]byte(SecretKey))
	if err!=nil{
		logs.Error(err)
		return
	}
	return
}

//解析token
func ParseToken(tokenString string)(token *jwt.Token,err error){
	if tokenString==""{
		err=fmt.Errorf("ParseToken tokenString is nil")
		logs.Error(err)
		return
	}
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err!=nil{
		err=fmt.Errorf("token parse error:%v",err)
		logs.Error(err)
		return
	}
	return
}

//小于一天刷新token
func IsRefreshToken(claims jwt.Claims)(token string,isRefresh bool,err error){
	if claims.(jwt.MapClaims)["exp"]==nil{
		err = fmt.Errorf("读取token参数错误")
		logs.Error(err)
		return
	}
	if claims.(jwt.MapClaims)["exp"].(float64) < float64(time.Now().Add(time.Hour*24).Unix()) {
		isRefresh=true
		token,err=CreateToken(claims.(jwt.MapClaims)["username"].(string),claims.(jwt.MapClaims)["role"].(string))
		if err!=nil{
			logs.Error(err)
			return
		}
	}else{
		isRefresh=false
	}
	return
}