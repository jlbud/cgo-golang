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

type PostBase struct {
	b         unsafe.Pointer // base handle
	modelPath *C.char
}

func CreatePostBase(modelPath string) (*PostBase, error) {
	pb := &PostBase{
		modelPath: C.CString(modelPath),
	}
	rec := C.LoadT2sModel(pb.modelPath, &pb.b)
	if rec != 0 {
		return nil, fmt.Errorf("PostBaseHandle LoadT2sModel fail code:%d", rec)
	}
	return pb, nil
}

func (pb *PostBase) Destroy() error {
	rec := C.UnloadT2sModel(&pb.b)
	if rec != 0 {
		return fmt.Errorf("PostBaseHandle UnloadT2sModel fail code:%d", rec)
	}
	C.free(unsafe.Pointer(pb.modelPath))
	return nil
}

func (pb *PostBase) CreateSession() (*PostSession, error) {
	ps := new(PostSession)
	rec := C.InitializeT2sInstance(pb.b, &ps.h)
	if rec != 0 {
		return nil, fmt.Errorf("PostSession InitializeT2sInstance fail code:%d", rec)
	}
	return ps, nil
}

type PostSession struct {
	h unsafe.Pointer // handle
}

func (ps *PostSession) Process(input string, endFlag int) (output string, err error) {
	var out *C.char
	in := C.CString(input)
	defer C.free(unsafe.Pointer(in))
	rec := C.T2sProcess(ps.h, in, (C.int)(endFlag), &out)
	if rec != 0 {
		return "", fmt.Errorf("PostSession T2sProcess fail code:%d", rec)
	}
	return C.GoString(out), nil
}

func (ps *PostSession) Reset() error {
	rec := C.T2sReset(ps.h)
	if rec != 0 {
		return fmt.Errorf("PostSession T2sReset fail code:%d", rec)
	}
	return nil
}
