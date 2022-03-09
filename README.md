# flashdb-bench

* [BadgerDB](https://github.com/dgraph-io/badger) (master branch with default options)
* [BoltDB](https://github.com/boltdb/bolt) (master branch  with default options)
* [NutsDB](https://github.com/xujiajun/nutsdb) (master branch with default options or custom options)
* [FlashDB](https://github.com/arriqaaq/flashdb) (master branch with default options or custom options)

## Benchmark System:

* Go Version : go1.17.3 darwin/amd64
* OS: Mac OS X 10.13.6
* Architecture: x86_64
* 16 GB 2133 MHz LPDDR3
* CPU: 3.1 GHz Intel Core i7


## To run benchmark:

> go test -bench=.

##  Benchmark results:

```
badger 2022/03/09 14:04:44 INFO: All 0 tables opened in 0s
goos: darwin
goarch: amd64
pkg: github.com/arriqaaq/flashbench
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz

BenchmarkBadgerDBPutValue64B-16     	    9940	    141844 ns/op	    2208 B/op	      68 allocs/op
BenchmarkBadgerDBPutValue128B-16    	    7701	    192942 ns/op	    2337 B/op	      68 allocs/op
BenchmarkBadgerDBPutValue256B-16    	    7368	    142600 ns/op	    2637 B/op	      69 allocs/op
BenchmarkBadgerDBPutValue512B-16    	    6980	    148056 ns/op	    3149 B/op	      69 allocs/op
BenchmarkBadgerDBGet-16             	 1000000	      1388 ns/op	     408 B/op	       9 allocs/op

BenchmarkFlashDBPutValue64B-16      	  159992	      7208 ns/op	    1461 B/op	      19 allocs/op
BenchmarkFlashDBPutValue128B-16     	  175634	      9499 ns/op	    2003 B/op	      19 allocs/op
BenchmarkFlashDBPutValue256B-16     	  148362	      9278 ns/op	    3322 B/op	      19 allocs/op
BenchmarkFlashDBPutValue512B-16     	  120865	     16542 ns/op	    6037 B/op	      19 allocs/op
BenchmarkFlashDBGet-16              	 1881042	       643.9 ns/op	      32 B/op	       2 allocs/op

PASS
ok  	github.com/arriqaaq/flashbench	28.947s
```

#### With FSync every write

```
badger 2022/03/09 14:04:44 INFO: All 0 tables opened in 0s
goos: darwin
goarch: amd64
pkg: github.com/arriqaaq/flashbench
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz

BenchmarkNutsDBPutValue64B-16     	      52	  20301019 ns/op	    1315 B/op	      17 allocs/op
BenchmarkNutsDBPutValue128B-16    	      63	  23496536 ns/op	    1059 B/op	      15 allocs/op
BenchmarkNutsDBPutValue256B-16    	      62	  20037952 ns/op	    1343 B/op	      15 allocs/op
BenchmarkNutsDBPutValue512B-16    	      62	  20090731 ns/op	    1754 B/op	      15 allocs/op

BenchmarkFlashDBPutValue64B-16     	     105	  16154174 ns/op	     770 B/op	      18 allocs/op
BenchmarkFlashDBPutValue128B-16    	      62	  21666543 ns/op	    1119 B/op	      18 allocs/op
BenchmarkFlashDBPutValue256B-16    	      61	  18631201 ns/op	    1956 B/op	      18 allocs/op
BenchmarkFlashDBPutValue512B-16    	      62	  20118103 ns/op	    3071 B/op	      18 allocs/op

PASS
ok  	github.com/arriqaaq/flashbench	28.947s
```
