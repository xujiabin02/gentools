package common

import "github.com/astaxie/beego/logs"

func LogConsole() *logs.BeeLogger {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	return log
}
