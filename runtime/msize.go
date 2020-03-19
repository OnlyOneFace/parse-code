package runtime

/*
* @
* @Author:
* @Date: 2020/3/19 20:41
 */
// Returns size of the memory block that mallocgc will allocate if you ask for the size.//返回mallocgc在请求大小时将分配的内存块大小
func roundupsize(size uintptr) uintptr {
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			return uintptr(class_to_size[size_to_class8[(size+smallSizeDiv-1)/smallSizeDiv]])
		} else {
			return uintptr(class_to_size[size_to_class128[(size-smallSizeMax+largeSizeDiv-1)/largeSizeDiv]])
		}
	}
	if size+_PageSize < size {
		return size
	}
	return round(size, _PageSize)
}