
Fast hash table for counting short strings in Go. Seems to be almost twice as fast for counting unique words in a text file:

```
$ ./test.sh 
Testing map output against baseline
Testing counter output against baseline
Benchmarking map version

real    0m0.785s
user    0m0.835s
sys 0m0.037s

real    0m0.776s
user    0m0.842s
sys 0m0.021s
...
real    0m0.778s
user    0m0.839s
sys 0m0.025s

real    0m0.770s
user    0m0.837s
sys 0m0.017s

real    0m0.772s
user    0m0.815s
sys 0m0.046s
Benchmarking counter version

real    0m0.424s
user    0m0.404s
sys 0m0.028s

real    0m0.443s
user    0m0.435s
sys 0m0.016s

real    0m0.434s
user    0m0.406s
sys 0m0.036s

real    0m0.432s
user    0m0.400s
sys 0m0.040s

real    0m0.439s
user    0m0.415s
sys 0m0.032s
```

There are also these micro-benchmark results against three main variants:

1) Counter: use the `counter.Counter` type, whose `Inc` method takes a `[]byte`.
2) MapBytes: have the words come in as `[]byte`, which is how you'd get them from a read operation. This is slower because they have to be converted to string before inserting into the map (you can't have a map with a slice key type).
3) MapString: have the words come in as `string`. This is much faster, but also not realistic, because when reading a file you'll be getting `[]byte`.

For each variant, there's a "mostly-unique" and a "non-unique" version. The mostly-unique one inserts 10,000 random strings, which will mean approximately 10,000 insert operations. The non-unique version uses 1000 random strings but does it 10 times, so it will be 1000 insert and 9000 update operations.

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/benhoyt/counter
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkMostlyUniqueCounter-8          3072        340033 ns/op
BenchmarkNonUniqueCounter-8             5718        200205 ns/op
BenchmarkMostlyUniqueMapBytes-8         1335        789635 ns/op
BenchmarkNonUniqueMapBytes-8            2317        495804 ns/op
BenchmarkMostlyUniqueMapString-8        2834        429449 ns/op
BenchmarkNonUniqueMapString-8           7146        176765 ns/op
PASS
ok      github.com/benhoyt/counter  7.173s
```

