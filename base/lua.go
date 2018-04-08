package base

const (
	//LuaVersion 为版本号
	LuaVersion = "Lua 5.1"
	//LuaRelease 为详细版本号
	LuaRelease = "Lua 5.1.5"
	//LuaVersionNum 为数字格式版本号
	LuaVersionNum = 501
	//LuaAuthors 为作者
	LuaAuthors = "R. Ierusalimschy, L. H. de Figueiredo & W. Celes"
	//LuaCopyright 为版权信息
	LuaCopyright = "Copyright (C) 1994-2012 Lua.org, PUC-Rio"
	//LuaSignature 为预编译文件的标志头
	LuaSignature = "\033Lua"
)

//LuaBasicType Lua的基本数据类型
type LuaBasicType int8

const (
	//LuaTNone 表示无类型
	LuaTNone LuaBasicType = iota - 1
	//LuaTNil 表示空类型
	LuaTNil
	//LuaTBoolean 表示布尔类型
	LuaTBoolean
	//LuaTLightUserData 表示指针类型数据
	LuaTLightUserData
	//LuaTNumber 表示数字类型
	LuaTNumber
	//LuaTString 表示字符串
	LuaTString
	//LuaTTable 表示表
	LuaTTable
	//LuaTFunction 表示函数
	LuaTFunction
	//LuaTUserData 指针
	LuaTUserData
	//LuaTThread Lua虚拟机，协程
	LuaTThread
)

//ThreadStatus 表示线程状态
type ThreadStatus int8

const (
	//LuaOK 表示ok
	LuaOK ThreadStatus = iota
	//LuaYield 表示yield状态
	LuaYield
	//LuaErrRun 表示运行错误
	LuaErrRun
	//LuaErrSyntax 表示语法错误
	LuaErrSyntax
	//LuaErrMem 表示内存错误
	LuaErrMem
	//LuaErrErr 表示错误
	LuaErrErr
)

//LuaNumber Lua的数值定义
type LuaNumber float64
