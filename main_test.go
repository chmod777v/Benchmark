package main

import "testing"

func BenchmarkCopy(b *testing.B) {
	input := []string{"1", "2", "3"}

	for i := 0; i < b.N; i++ {
		copy2(input)
	}
}

//go test -bench . -benchmem -cpuprofile=profile
//go tool pprof profile
//(pprof) top -топ 10 операций по времени/памяти
//(pprof) web - график

//(pprof) list copy1 - сколько времени код работал на каонкретной строке, copy1-функция
//(pprof) list runtime.nextFreeFast - функция из top

//go test -bench . -benchmem -memprofile=mem_profile

//go test -bench . -benchmem -count=5  > profile_old
