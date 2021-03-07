
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
