package glua

import (
	"bytes"
)

type StringTable struct {
	Hash []*GCObject
	Nuse uint32
	Size int
}

type TString struct {
	CommonHeader
	Reserved uint8
	Hash     uint
	Len      int
	Data     []byte
}

func NewTString(L *State, str []byte) *TString {
	l := len(str)
	var h uint = uint(l)
	var step int = (l >> 5) + 1
	var l1 int
	for l1 = l; l1 >= step; l1 -= step {
		h = h ^ ((h << 5) + (h >> 2) + uint(str[l1-1]))
	}
	for o := L.G.Strt.Hash[lmod(h, L.G.Strt.Size)]; o != nil; o = o.Gch.Next {
		var tsobj TString = o.Obj.(TString)
		ts := &tsobj
		if ts.Len == l && bytes.Compare(str, ts.Data[0:l]) == 0 {
			gc := L.G.GC
			if gc.isDead(o) {
				gc.changeWhite(o)
			}
		}
	}
	return newTString(L, str, l, h)
}

func newTString(L *State, str []byte, l int, h uint) *TString {
	if l+1 > MAXINT {
		tooBig(L)
	}
	gc := L.G.GC
	ts := new(TString)
	ts.Len = l
	ts.Hash = h
	ts.Marked = gc.white(L.G)
	ts.Reserved = 0
	ts.Tt = TSTRING
	ts.Data = str[:]

	tb := L.G.Strt
	h = uint(lmod(h, tb.Size))
	ts.Next = tb.Hash[h]
	tb.Hash[h] = &GCObject{
		Obj: ts,
	}
	tb.Nuse++
	//TODO 这里是否要实现resize
	return ts
}
