package base

//GlobalState 表示全局的state
type GlobalState struct {
}

//LuaState 表示lua的state
type LuaState struct {
	top       StkID //栈中第一个空的位置
	base      StkID //函数参数的位置
	stackLast StkID //栈中最后一个空的位置
	stack     StkID //栈底
	lG        *GlobalState
}

func (L *LuaState) luaGetTop() int {
	return 1
}
