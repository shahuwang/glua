package glua

type GCObject struct {
	Gch GCHeader
	Obj interface{}
}

type CommonHeader struct {
	Next   *GCObject
	Tt     uint8
	Marked uint8
}

type GCHeader struct {
	CommonHeader
}

type State struct {
	CommonHeader
	G *GlobalState
}

type GlobalState struct {
	Strt StringTable
	GC   *GC
}
