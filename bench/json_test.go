// Package main
/*
 Author: hawrkchen
 Date: 2022/3/8 9:41
 Desc:
*/
package bench

import "testing"

func BenchmarkJsoniter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JsoniterEncode()
	}
}

func BenchmarkJson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JsonEncode()
	}
}

