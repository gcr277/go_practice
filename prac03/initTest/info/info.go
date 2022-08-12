package info

import(
	"runtime"
)

// 获取正在运行的函数名
func CurrFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
