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
	http.HandleFunc("/goods/query",controllers.CheckTokenPreHand(controllers.SearchGoodsHand))
	http.HandleFunc("/goods/update",controllers.CheckTokenPreHand(controllers.UpdateGoodsMes))
	http.HandleFunc("/goods/del",controllers.CheckTokenPreHand(controllers.DelGoodHand))
	http.HandleFunc("/users",controllers.CheckTokenPreHand(controllers.UsersListHand))
	http.HandleFunc("/users/query",controllers.CheckTokenPreHand(controllers.UsersListHand))
	http.HandleFunc("/users/update",controllers.CheckTokenPreHand(controllers.UpdateUserHand))
	http.HandleFunc("/users/del",controllers.CheckTokenPreHand(controllers.DelUserHand))
	http.HandleFunc("/monthly",controllers.CheckTokenPreHand(controllers.MonthlyHand))
	
	err=http.ListenAndServe(":8080",nil)
	if err!=nil{
		logs.Error(err)
		return
	}
}