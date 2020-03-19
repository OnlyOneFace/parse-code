package runtime

/*
* @
* @Author:
* @Date: 2020/3/19 16:22
 */

// Garbage collector phase.
// Indicates to write barrier and synchronization task to perform. //垃圾收集器阶段。指示要执行的写入屏障和同步任务。
var gcphase uint32

const (
	_GCoff             = iota // GC not running; sweeping in background, write barrier disabled //GC未运行；在后台扫描，写屏障已禁用
	_GCmark                   // GC marking roots and workbufs: allocate black, write barrier ENABLED //GC标记根和workbufs：分配黑色，启用写屏障
	_GCmarktermination        // GC mark termination: allocate black, P's help GC, write barrier ENABLED //GC标记终止：allocate black，P's help GC，write barrier ENABLED
)
