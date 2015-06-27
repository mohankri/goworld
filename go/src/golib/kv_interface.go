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

func (kv kv_struct) kv_get(key[] string, key_size int) int {
	var k string
	cs := C.CString(k)
	len := C.int(key_size)
	C.Cget_key(kv.kvptr, cs, &len)
	//copy(key[:], C.GoString(cs))	
	//key = C.GoString(cs)
	fmt.Println("Key Get ", key, " Size ", len, kv.kvptr)
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
