package cpp

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L ./lib -lgo4c

#include "go4c.h"
#include <stdlib.h>

void FnCallBackGO(char* data, int32_t len);
*/
import "C"
import (
	"log"
	"reflect"
	"unsafe"
)

var defaultCallback = func(data []byte) {
	log.Printf(">> %v\n", string(data))
}

//export FnCallBackGO
func FnCallBackGO(data *C.char, len C.int32_t) {
	//var s = C.GoStringN(data, len)
	var s []byte = C.GoBytes(unsafe.Pointer(data), len)

	defaultCallback(s)
}

// InitLibrary 初始化
func InitLibrary(fun func(data []byte)) int {
	defaultCallback = fun

	var rc C.int32_t = C.Go4CInit_C(C.FnCallback_C(C.FnCallBackGO))

	return int(rc)
}

// ReleaseLibrary 释放库
func ReleaseLibrary() {
	C.Go4CRelease_C()
}

// Command 调用动态库函数
func Command(data []byte) int {
	//p := (*reflect.StringHeader)(unsafe.Pointer(&data))
	//var rc C.int32_t = C.Go4CInitCommand_C((*C.char)(unsafe.Pointer(p.Data)), C.int32_t(len(data)))

	var rc C.int32_t = C.Go4CInitCommand_C(
		(*C.char)(unsafe.Pointer(
			(*reflect.StringHeader)(unsafe.Pointer(&data)).Data)),
		C.int32_t(len(data)))

	return int(rc)
}
