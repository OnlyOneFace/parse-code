package runtime

/*
* @
* @Author:
* @Date: 2020/3/19 16:31
 */
// getg returns the pointer to the current g.
// The compiler rewrites calls to this function into instructions
// that fetch the g directly (from TLS or from the dedicated register). //getg返回指向当前g的指针。编译器将对此函数的调用重写为直接（从TLS或专用寄存器）获取g的指令
func getg() *g

// systemstack runs fn on a system stack.
// If systemstack is called from the per-OS-thread (g0) stack, or
// if systemstack is called from the signal handling (gsignal) stack,
// systemstack calls fn directly and returns.
// Otherwise, systemstack is being called from the limited stack
// of an ordinary goroutine. In this case, systemstack switches
// to the per-OS-thread stack, calls fn, and switches back.
// It is common to use a func literal as the argument, in order
// to share inputs and outputs with the code around the call
// to system stack: // 系统堆栈在系统堆栈上运行fn。如果系统堆栈是从每个操作系统线程（g0）堆栈调用的，或者如果系统堆栈是从信号处理（gsignal）堆栈调用的，则系统堆栈直接调用fn并返回。否则，将从普通goroutine的有限堆栈调用systemstack。在这种情况下，systemstack切换到每个OS线程堆栈，调用fn，然后切换回。通常使用func文字作为参数，以便与系统堆栈调用周围的代码共享输入和输出：
//
//	... set up y ...
//	systemstack(func() {
//		x = bigcall(y)
//	})
//	... use x ...
//
//go:noescape

func systemstack(fn func()) // 查看汇编asm_amd64.s文件
