package runtime

/*
* @
* @Author:
* @Date: 2020/3/19 16:28
 */

//go:nosplit
func throw(s string) {
	// Everything throw does should be recursively nosplit so it
	// can be called even when it's unsafe to grow the stack. //throw所做的一切都应该递归nosplit，这样即使在堆栈增长不安全的情况下也可以调用它
	systemstack(func() {
		print("fatal error: ", s, "\n")
	})
	gp := getg()
	if gp.m.throwing == 0 {
		gp.m.throwing = 1
	}
	fatalthrow()
	*(*int)(nil) = 0 // not reached
}
