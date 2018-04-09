package base

import (
	"os"
)

//LuaStatus 用于作为Lua执行结果的值类型
type LuaStatus uint8

const (
	//FailedStatus 为最终执行结果，表示失败
	FailedStatus LuaStatus = iota
	//SuccessStatus 为最终执行结果，表示成功
	SuccessStatus
)

//LoadF 是加载到的数据的存储结构
type LoadF struct {
	extraline int
	f         *os.File
	buff      [LUAL_BUFFERSIZE]byte
}

// func luaLLoadFile(L *LuaState, filename string) LuaStatus {
// 	var lf LoadF
// 	fnameindex := L.luaGetTop() + 1
// 	lf.extraline = 0
// 	if filename == "" {

// 	}
// }
