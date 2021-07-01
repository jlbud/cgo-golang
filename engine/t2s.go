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

type PostBaseHandle struct {
	b         unsafe.Pointer // base handle
	modelPath *C.char
}

func CreatePostBaseHandle(modelPath string) (*PostBaseHandle, error) {
	p := &PostBaseHandle{}
	p.modelPath = C.CString(modelPath)
	rec := C.LoadT2sModel(p.modelPath, &p.b)
	if rec != 0 {
		return nil, fmt.Errorf("T2SModel LoadT2sModel fail code:%d", rec)
	}
	return p, nil
}

type PostSession struct {
	h unsafe.Pointer // handle
}

func CreatePostSession(bashHandle *PostBaseHandle) (*PostSession, error) {
	ps := &PostSession{}
	rec := C.InitializeT2sInstance(bashHandle.b, &ps.h)
	if rec != 0 {
		return nil, fmt.Errorf("T2SModel InitializeT2sInstance fail code:%d", rec)
	}
	return ps, nil
}

func (ps *PostSession) Process(input string, endFlag int) (output string, err error) {
	var out *C.char
	in := C.CString(input)
	defer C.free(unsafe.Pointer(in))
	rec := C.T2sProcess(ps.h, in, (C.int)(endFlag), &out)
	if rec != 0 {
		return "", fmt.Errorf("T2SModel Process fail code:%d", rec)
	}
	return C.GoString(out), nil
}

func (pb *PostBaseHandle) Destroy() error {
	rec := C.UnloadT2sModel(&pb.b)
	if rec != 0 {
		return fmt.Errorf("T2SModel UnloadT2sModel fail code:%d", rec)
	}
	C.free(unsafe.Pointer(pb.modelPath))
	return nil
}

func (pb *PostSession) Reset() error {
	rec := C.T2sReset(&pb.h)
	if rec != 0 {
		return fmt.Errorf("T2SModel Reset fail code:%d", rec)
	}
	return nil
}
