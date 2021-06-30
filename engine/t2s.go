package engine

/*
#cgo LDFLAGS: -L ./libs -lt2s
#include<stdlib.h>
#include<string.h>
#include <string.h>
#include "t2s.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type T2SModel struct {
	b         unsafe.Pointer // base handle
	h         unsafe.Pointer // handle
	modelPath *C.char
	UserRules
}

type UserRules struct {
	userDictPath *C.char // user dict path
	size         C.int
	inFile       *C.char // user themself make rules file
	outFile      *C.char // user themself make rules model file
}

func CreateInstance(modelPath string) *T2SModel {
	b := &T2SModel{}
	b.modelPath = C.CString(modelPath)
	rec := C.LoadT2sModel(b.modelPath, &b.b)
	if rec != 0 {
		return nil
	}
	rec = C.InitializeT2sInstance(b.b, &b.h)
	if rec != 0 {
		return nil
	}
	return b
}

func (t *T2SModel) Process(input string, endFlag int) (output string) {
	var out *C.char
	in := C.CString(input)
	defer C.free(unsafe.Pointer(in)) // TODO
	C.T2sProcess(t.h, in, (C.int)(endFlag), &out)
	return C.GoString(out)
}

func (t *T2SModel) Destroy() error {
	rec := C.TerminateT2sInstance(&t.h)
	if rec != 0 {
		return fmt.Errorf("T2SModel TerminateT2sInstance fail: code:%d", rec)
	}
	rec = C.UnloadT2sModel(&t.b)
	if rec != 0 {
		return fmt.Errorf("T2SModel UnloadT2sModel fail: code:%d", rec)
	}
	C.free(unsafe.Pointer(t.modelPath))
	return nil
}

func (t *T2SModel) LoadUserRules(userDictPath string, size int) error {
	t.userDictPath = C.CString(userDictPath)
	t.size = C.int(size)
	rec := C.LoadUserRules(t.userDictPath, t.size, t.h)
	if rec != 0 {
		return fmt.Errorf("T2SModel LoadUserRules fail: code:%d", rec)
	}
	return nil
}

func (t *T2SModel) UnLoadUserRules() error {
	rec := C.UnloadUserRules(t.h)
	if rec != 0 {
		return fmt.Errorf("T2SModel UnloadUserRules fail: code:%d", rec)
	}
	return nil
}

func (t *T2SModel) GetVersion() string {
	v := C.T2sGetVersion()
	return C.GoString(v)
}

//func (t *T2SModel) MakeUserModel(inFile, outFile string) error {
//	t.inFile = C.CString(inFile)
//	t.outFile = C.CString(outFile)
//	rec := C.MakeT2sModel(t.inFile, t.outFile)
//	if rec != 0 {
//		return fmt.Errorf("T2SModel MakeT2sModel fail: code:%d", rec)
//	}
//	return nil
//}
