//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package xreflect

// https://github.com/golang/go/blob/go1.18rc1/src/runtime/symtab.go#L415-L457

// moduledata is almost the same as runtime.moduledata.
type moduledata struct {
	pcHeader     *uintptr // *pcHeader
	funcnametab  []byte
	cutab        []uint32
	filetab      []byte
	pctab        []byte
	pclntable    []byte
	ftab         []functab
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr
	types, etypes         uintptr
	rodata                uintptr
	gofunc                uintptr

	textsectmap []textsect
	typelinks   []int32
	itablinks   []*uintptr // []*itab

	ptab []ptabEntry

	pluginpath string
	pkghashes  []modulehash

	modulename   string
	modulehashes []modulehash

	hasmain uint8

	gcdatamask, gcbssmask bitvector

	typemap map[int32]*uintptr // map[typeOff]*_type

	bad bool

	next *moduledata
}

type functab struct {
	entryoff uint32
	funcoff  uint32
}

type textsect struct {
	vaddr    uintptr
	end      uintptr
	baseaddr uintptr
}

type ptabEntry struct {
	name int32 // nameOff
	typ  int32 // typeOff
}

type modulehash struct {
	modulename   string
	linktimehash string
	runtimehash  *string
}

type bitvector struct {
	n        int32
	bytedata *uint8
}
