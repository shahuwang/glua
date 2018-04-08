package base

//TValue 存储Lua里所有的数据类型
type TValue struct {
	value interface{}
	tt    LuaBasicType
}

// StkID 用于表示栈指针
type StkID *TValue
