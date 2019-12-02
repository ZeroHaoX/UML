package main

import (
	"github.com/astaxie/beego/logs"
	"./logger"
	"log"
	"net/http"
	"./models"
	"./controllers"
)

func main(){
	err:=logger.InitLogger()
	if err!=nil{
		log.Panic("初始化日志失败")
	}

	if err:=models.InitDB();err!=nil{
		log.Panic("初始化数据库失败")
	}
	
	http.HandleFunc("/login",controllers.PreHand(controllers.LoginHand))
	http.HandleFunc("/registe",controllers.PreHand(controllers.RegisteHand))
	http.HandleFunc("/export",controllers.CheckTokenPreHand(controllers.ExportHand))
	http.HandleFunc("/import",controllers.CheckTokenPreHand(controllers.ImportHand))
	http.HandleFunc("/goods",controllers.CheckTokenPreHand(controllers.GoodsListHand))
	http.HandleFunc("/goods/update",controllers.CheckTokenPreHand(controllers.UpdateGoodsMes))
	http.HandleFunc("/goods/del",controllers.CheckTokenPreHand(controllers.UpdateGoodsMes))
	http.HandleFunc("/users",controllers.CheckTokenPreHand(controllers.UsersListHand))
	
	err=http.ListenAndServe(":8080",nil)
	if err!=nil{
		logs.Error(err)
		return
	}
}