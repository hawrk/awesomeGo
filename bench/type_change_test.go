// Package bench
/*
 Author: hawrkchen
 Date: 2022/3/31 11:32
 Desc:
*/
package bench

import "testing"

func BenchmarkChangeWithGO(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChangeWithGO()
	}
}

func BenchmarkChangeWithCast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChangeWithCast()
	}
}
