package base

// Value 用于存放lua的值的位置，不包含类型信息
type Value interface{}

//TValue 存储Lua里所有的数据类型
type TValue struct {
	value Value
	tt    LuaBasicType
}

// StkID 用于表示栈指针
type StkID *TValue

//TODO  后续再来完善
type GCObject interface{}

//LuByte 在lua中的定义是 typedef unsigned char lu_byte
type LuByte byte

type CommonHeader struct {
	next   *GCObject
	tt     LuByte
	marked LuByte
}

type TString struct {
	CommonHeader
	reserved LuByte
	hash     uint
	len      int
}

type Stringtable struct {
	hash [][]*GCObject
	nuse LuInt32
	size int
}

//TODO lua的代码有两种实现，根本不同的编译条件，需要研究下golang怎么设置编译条件
func check_exp(cmp bool, e int) int {
	return e
}

func lmod(s int, size int) int {
	return check_exp(
		(size&(size-1)) == 0, (s)&(size-1))
}
