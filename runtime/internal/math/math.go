package math

import "parse-code/runtime/internal/sys"

const MaxUintptr = ^uintptr(0)

// MulUintptr returns a * b and whether the multiplication overflowed. //MulUintptr返回a*b以及乘法是否溢出
// On supported platforms this is an intrinsic lowered by the compiler.//在受支持的平台上，这是编译器降低的内部函数
func MulUintptr(a, b uintptr) (uintptr, bool) {
	if a|b < 1<<(4*sys.PtrSize) || a == 0 { // MaxUintptr的开平方的值大于一切32位数的乘积，a|b只要是32位以内的数即可
		return a * b, false
	}
	overflow := b > MaxUintptr/a // 说明a*b超出数据范围
	return a * b, overflow
}
