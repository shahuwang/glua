package base

//LuaStatus 用于作为Lua执行结果的值类型
type LuaStatus uint8

const (
	//FailedStatus 为最终执行结果，表示失败
	FailedStatus LuaStatus = iota
	//SuccessStatus 为最终执行结果，表示成功
	SuccessStatus
)
