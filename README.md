# iftruefalse-go

```bash
$ go test -bench .
```

- Result1

```bash
$ date; go test -bench .   
```

```result1
2022年 7月21日 木曜日 17時17分17秒 JST
goos: darwin
goarch: amd64
pkg: iftruefalse-go
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIftrue-8    	23352720	        47.71 ns/op
BenchmarkIffalse-8   	23869148	        50.15 ns/op
PASS
ok  	iftruefalse-go	5.434s
```

- Result2

```bash
$ date; go test -bench .
```

```result
2022年 7月21日 木曜日 17時17分24秒 JST
goos: darwin
goarch: amd64
pkg: iftruefalse-go
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIftrue-8    	18603630	        54.89 ns/op
BenchmarkIffalse-8   	23748326	        52.56 ns/op
PASS
ok  	iftruefalse-go	4.532s
```

- Result3

```bash
date; go test -bench .
```

```result
2022年 7月21日 木曜日 17時21分19秒 JST
goos: darwin
goarch: amd64
pkg: iftruefalse-go
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIftrue-8    	21789853	        58.65 ns/op
BenchmarkIffalse-8   	23304178	        51.71 ns/op
PASS
ok  	iftruefalse-go	4.547s
```

- Result4

```bash
date; go test -bench .
```

```result
2022年 7月21日 木曜日 17時22分20秒 JST
goos: darwin
goarch: amd64
pkg: iftruefalse-go
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkIftrue-8    	14065443	        76.32 ns/op
BenchmarkIffalse-8   	24842272	        44.38 ns/op
PASS
ok  	iftruefalse-go	3.437s
```

