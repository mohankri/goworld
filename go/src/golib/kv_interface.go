package golib 

// #include "cppwrap.h"
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"
import "unsafe"

type kv_interface interface {
	kv_get(key string, key_size int) int;
	kv_put(key string, key_size int) int;
}

type kv_struct struct {
	path string
	ssize, msize int
	kvptr C.Kvstore
}

func (kv kv_struct) kv_get(key string, key_size int) int {
	key = "KM"
	key_size = 0
	cs := C.CString(key)
	len := C.int(key_size)
	C.Cget_key(kv.kvptr, cs, len)
	fmt.Println("Key Get ", C.GoString(cs), " Size ", len, kv.kvptr)
	C.free(unsafe.Pointer(cs))
	return 0
}

func (kv kv_struct) kv_put(key string, key_size int) int {
	fmt.Println("Key ", key, " Size ", key_size, kv.kvptr);
	C.Cset_key(kv.kvptr, C.CString(key), C.int(key_size))
	return 0
}

func kv_open(path string, ssize int, msize int) kv_struct {
	kv := kv_struct{path:path, ssize:ssize, msize:msize}	
	kv.kvptr = C.Copen(C.CString(path), C.int(ssize), C.int(msize))
	fmt.Println("  kv open ", path, kv.kvptr)
	return kv
}
