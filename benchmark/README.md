# Benchmark Tests

Directory contains performance tests of various language components.

Run benchmarks with `go test -bench=.` inside each directory.

 - JSON vs. Gob - trying two different serializing techniques.
 - Rand vs. /dev/urandom - comparing internal Golang rand byte stream and Linux /dev/random stream reading.  
 Using `urandom`, because `random` is obviously too slow due to generating **completely** random bytes.