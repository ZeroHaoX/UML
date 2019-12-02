package common

import (
	"regexp"
	"github.com/astaxie/beego/logs"
	// "strconv"
	// "fmt"
	// "reflect"
)

//检验特殊字符
func HasSpecialCharacter(s string)(isTrue bool){
	isTrue,err:=regexp.MatchString(`\.|\=|\/|\~|\-|\\|\'|\"|\(|\)|\{|\}|\[|\]|\*|\+|\>|\<|\#`,s)
	if err!=nil{
		logs.Error(err)
		return false
	}
	return
}

//检查年月
func CheckDate(year int,month int)(isTrue bool){
	if year<1000{
		return
	}
	if month<1||month>12{
		return
	}
	return true
}


//检查结构体
// func CheckStructHasNil(data interface{})(ok bool){
// 	value:=reflect.ValueOf(data)
// 	for i:=0;i<value.NumField();i++{
// 		vtype:=value.Field(i).Type()
// 		switch vtype.String(){
// 		case "string":{
// 			if value.Field(i).String()==""{
// 				return
// 			}
// 		}
// 		case "float32":{
// 			if value.Field(i).String()=="0"{
// 				return
// 			}
// 		}
// 		}
// 	}
// 	return true
// }