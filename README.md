# LRU

[![Build Status](https://travis-ci.org/andersnormal/lru.svg?branch=master)](https://travis-ci.org/andersnormal/lru)
[![Go Report Card](https://goreportcard.com/badge/github.com/andersnormal/lru)](https://goreportcard.com/report/github.com/andersnormal/lru)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A simple [LRU Cache](https://www.geeksforgeeks.org/lru-cache-implementation/) implementation which supports TTL of cached items and fetching of keys which do not exist or have expired. You can build upon different cache mechanism, or other caches.

## Docs

You can find the documentation hosted on [godoc.org](https://godoc.org/github.com/andersnormal/lru).

## Install

```
go get github.com/andersnormal/lru
```

## Benchmarks

Run the benchmarks with `task benchmark`. The sample benchmarks are run on a MacBook Pro (2.7 GHz Intel Core i7).

```
BenchmarkCache_Rand/with_size_of_4096_and_no_ttl_set-8         	 3000000	       544 ns/op
--- BENCH: BenchmarkCache_Rand/with_size_of_4096_and_no_ttl_set-8
    simple_test.go:240: hit: 0 miss: 1 ratio: 0.000000
    simple_test.go:240: hit: 0 miss: 100 ratio: 0.000000
    simple_test.go:240: hit: 861 miss: 9139 ratio: 0.094212
    simple_test.go:240: hit: 111123 miss: 888877 ratio: 0.125015
    simple_test.go:240: hit: 332972 miss: 2667028 ratio: 0.124848
BenchmarkCache_Rand/with_size_of_8092_items_and_no_ttl_set-8   	 3000000	       580 ns/op
--- BENCH: BenchmarkCache_Rand/with_size_of_8092_items_and_no_ttl_set-8
    simple_test.go:240: hit: 0 miss: 1 ratio: 0.000000
    simple_test.go:240: hit: 0 miss: 100 ratio: 0.000000
    simple_test.go:240: hit: 1333 miss: 8667 ratio: 0.153802
    simple_test.go:240: hit: 197704 miss: 802296 ratio: 0.246423
    simple_test.go:240: hit: 593771 miss: 2406229 ratio: 0.246764
BenchmarkCache_Freq/with_size_of_4096_and_no_ttl_set-8         	 3000000	       560 ns/op
--- BENCH: BenchmarkCache_Freq/with_size_of_4096_and_no_ttl_set-8
    simple_test.go:291: hit: 1 miss: 0 ratio: +Inf
    simple_test.go:291: hit: 100 miss: 0 ratio: +Inf
    simple_test.go:291: hit: 8206 miss: 1794 ratio: 4.574136
    simple_test.go:291: hit: 236672 miss: 763328 ratio: 0.310053
    simple_test.go:291: hit: 702382 miss: 2297618 ratio: 0.305700
BenchmarkCache_Freq/with_size_of_8092_items_and_no_ttl_set-8   	 3000000	       557 ns/op
--- BENCH: BenchmarkCache_Freq/with_size_of_8092_items_and_no_ttl_set-8
    simple_test.go:291: hit: 1 miss: 0 ratio: +Inf
    simple_test.go:291: hit: 100 miss: 0 ratio: +Inf
    simple_test.go:291: hit: 8200 miss: 1800 ratio: 4.555556
    simple_test.go:291: hit: 233189 miss: 766811 ratio: 0.304102
    simple_test.go:291: hit: 708765 miss: 2291235 ratio: 0.309338
PASS
ok  	lru	19.696s
```

## License
[Apache 2.0](/LICENSE)
