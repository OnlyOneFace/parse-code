package runtime

import "unsafe"

/*
* @
* @Author:
* @Date: 2020/3/18 17:38
 */


// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.
type typeAlg struct {
	// function for hashing objects of this type 用于散列此类型对象的函数
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type 用于比较此类型对象的函数
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}