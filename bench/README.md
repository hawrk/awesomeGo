##测试说明
1. 测试文件名要与测试函数在同一个包下
2. 测试文件名要以_test结尾
3. 用例函数要以 Benchmark开头，并且入参为 b *testing.B，用途函数首字母大写
4. 运行 go test -bench=.   (win: go test -bench="."),表示测试所用例，如果要
   测某个用例，则用 go test -bench=".*XXX.*"
5. 一般运行benchmark测试前，需要忽略一些准备工作, b.ResetTimer() 开始计时
6. 结果：

   goos: windows
   goarch: amd64
   pkg: hjsoniter
   cpu: 11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz
   BenchmarkJsoniter-16             4750550               245.4 ns/op           192 B/op          3 allocs/op
   BenchmarkJson-16                 3595208               335.1 ns/op           192 B/op          3 allocs/op
   PASS
   ok      hjsoniter       3.628s
   说明： BenchmarkJsoniter-16 中的16表示16核，   4750550 表示单位时间内运行的次数，
   245.4 表示每次运行耗时，每次运行分配内存，和一共分配的内存
7. benchmark 可配置参数
   -benchtime=5s   指定单位测试时间为5s
   -benchmem 表示加上内存分配信息