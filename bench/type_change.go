// Package bench
/*
 Author: hawrkchen
 Date: 2022/3/31 11:29
 Desc:
*/
package bench

import (
	"github.com/spf13/cast"
)

/*
BenchmarkChangeWithGO-16        1000000000               0.2208 ns/op          0 B/op          0 allocs/op
BenchmarkChangeWithCast-16      92593306                12.17 ns/op            4 B/op          1 allocs/op
-_-    直接用 cast 比原生的直接，效率 还是差了50倍左右，所以一些高频操作，还是直接用原生的类型转换好
*/

func ChangeWithGO() {
	var a int32
	var b uint64
	a = 454454
	b = uint64(a)
	_ = b
	//fmt.Println("b:",b)
}

func ChangeWithCast() {
	var a int32
	var b uint64
	a = 454454
	b = cast.ToUint64(a)
	_ = b
}
