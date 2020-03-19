package sys

/*
* @
* @Author:
* @Date: 2020/3/18 17:49
 */
// Declarations for runtime services implemented in C or assembly. 在C或程序集中实现的运行时服务的声明。

const PtrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const 结果为8
