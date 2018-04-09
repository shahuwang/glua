package base

//GlobalState 表示全局的state
type GlobalState struct {
}

//LuaState 表示lua的state
type LuaState struct {
	Top       int     //栈中第一个空的位置
	Base      int     //函数参数的位置
	StackLast int     //栈中最后一个空的位置
	Stack     []StkID //栈底
	Stacksize int
	lG        *GlobalState
}

func (L *LuaState) luaGetTop() int {
	return L.Top - L.Base
}

func (L *LuaState) G() *GlobalState {
	return L.lG
}
