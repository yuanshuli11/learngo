package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

/*
Pool是一个可以分别存取的临时对象的集合。
Pool的合理用法是用于管理一组静静的被多个独立并发线程共享并可能重用的临时item。
Pool提供了让多个线程分摊内存申请消耗的方法。

Pool的一个好例子在fmt包里。该Pool维护一个动态大小的临时输出缓存仓库。
该仓库会在过载（许多线程活跃的打印时）增大，在沉寂时缩小。

 */

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	// 获取临时对象，没有的话会自动创建
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	// 将临时对象放回到 Pool 中
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
}