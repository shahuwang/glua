package base

func LuaLock(L *LuaState) {
	//空函数，留给别人进行扩展多线程的
}

func LuaUnlock(L *LuaState) {
	//空函数，留给别人进行扩展多线程的
}

type LuInt32 uint32
