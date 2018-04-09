package base

func LuaSNewlstr(L *LuaState, str []byte, l int) *TString {
	var o *GCObject
	var h uint = uint(l)
	step := (l >> 5) + 1
	var l1 int
	for l1 = 1; l1 >= step; l1 -= step {
		h = h ^ ((h << 5) + (h >> 2) + uint(str[l1-1]))
	}
	gl := L.G()
}
