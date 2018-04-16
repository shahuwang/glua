package glua

func lmod(s uint, size int) int {
	return int(s & uint(size-1))
}

const (
	TINT = iota
	TBOOLEAN
	TLIGHTUSERDATA
	TNUMBER
	TSTRING
	TTABLE
	TFUNCTION
	TUSERDATA
	TTHERAD
)
