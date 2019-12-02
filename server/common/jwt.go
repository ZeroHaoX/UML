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
	// if w==nil||r==nil{
	// 	err=fmt.Errorf("Parse token w or r is nil")
	// 	logs.Error(err)
	// 	return
	// }
	// cookie,err:=r.Cookie("token")
	// if err!=nil{
	// 	err=fmt.Errorf("get token from cookie error:%v",err)
	// 	logs.Error(err)
	// 	return
	// }
	// if cookie==nil{
	// 	err=fmt.Errorf("get token from cookie nil")
	// 	logs.Error(err)
	// 	return
	// }
	// tokenString:=cookie.Value
	// if tokenString==""{
	// 	err=fmt.Errorf("parse tokenString=%v",tokenString)
	// 	logs.Error(err)
	// 	return
	// }
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
	// valid=token.Valid
	// if valid{
	// 	tokenString,isRefresh,err:=IsRefreshToken(token.Claims)
	// 	if err!=nil{
	// 		logs.Error(err)
	// 		return nil,false,err
	// 	}
	// 	if isRefresh==true{
	// 		cookie:=&http.Cookie{
	// 			Name:"token",
	// 			Value:tokenString,
	// 		}
	// 		http.SetCookie(w,cookie)
	// 	}
		// claims=token.Claims
	// }
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
		token,err=CreateToken(claims.(jwt.MapClaims)["uname"].(string),claims.(jwt.MapClaims)["role"].(string))
		if err!=nil{
			logs.Error(err)
			return
		}
	}else{
		isRefresh=false
	}
	return
}