//go:build go1.13 && !go1.16
// +build go1.13,!go1.16

package xreflect

// https://github.com/golang/go/blob/go1.13/src/runtime/symtab.go#L264-L300
// https://github.com/golang/go/blob/go1.14/src/runtime/symtab.go#L267-L303
// https://github.com/golang/go/blob/go1.15/src/runtime/symtab.go#L342-L378

// moduledata is almost the same as runtime.moduledata.
type moduledata struct {
	pclntable    []byte
	ftab         []functab
	filetab      []uint32
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr
	types, etypes         uintptr

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
	entry   uintptr
	funcoff uintptr
}

type textsect struct {
	vaddr    uintptr
	length   uintptr
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
