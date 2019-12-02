package logger

import (
	"github.com/astaxie/beego/logs"
	"log"
	"encoding/json"
)


// InitLogger 初始化日志
func InitLogger() error {

	logs.SetLogFuncCallDepth(3)    //调用层级
	logs.EnableFuncCallDepth(true) //输出文件名和行号
	//logs.Async()                   //提升性能, 可以设置异步输出

	config := make(map[string]interface{})
	config["filename"] = `./log.log`

	configStr, err := json.Marshal(config)
	if err != nil {
		log.Fatal("initLogger failed, marshal err:", err)
		return err
	}

	logs.SetLevel(logs.LevelDebug)

	err = logs.SetLogger(logs.AdapterConsole, "") //控制台输出
	if err != nil {
		log.Fatal("SetLogger failed, err:", err)
		return err
	}

	err = logs.SetLogger(logs.AdapterFile, string(configStr)) //文件输出
	if err != nil {
		log.Fatal("SetLogger failed, err:", err)
		return err
	}

	

	//err = logs.SetLogger(logs.AdapterEs, `{"dsn":"http://localhost:9200/","level":1}`)
	//if err != nil {
	//	log.Fatal("SetLogger failed, err:", err)
	//	return err
	//}
	return nil
}