package main

import (
	"fmt"
	"testing"
)

func generateInput(size int) []string {
	input := make([]string, size)
	for i := range input {
		input[i] = fmt.Sprintf("element-%d", i)
	}
	return input
}
func BenchmarkCopy1(b *testing.B) {
	input := generateInput(1000)
	b.ReportAllocs() //можно не писать а использовать флаг -benchmem
	b.ResetTimer()   //желательно сбрасывать после подготовки

	for i := 0; i < b.N; i++ {
		copy1(input)
		//copy2(input) // → компилятор подставляет код copy2 прямо здесь
		//copy3(input) // → компилятор подставляет код copy3 прямо здесь поэтому здесь должны быть 1 функция для проверки

	}
}
func BenchmarkAllCopiesTable(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) []string
	}{
		{"Copy1", copy1},
		{"Copy2", copy2},
		{"Copy3", copy3},
	}
	input := generateInput(1000)
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) { //создает под-бенчмарк (sub-benchmark) внутри основного бенчмарка.
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
} /*
BenchmarkAllCopiesTable/Copy1-20                  251874              5376 ns/op           35184 B/op         11 allocs/op
BenchmarkAllCopiesTable/Copy2-20                  545947              2315 ns/op           16384 B/op          1 allocs/op
BenchmarkAllCopiesTable/Copy3-20                  666026              1920 ns/op           16384 B/op          1 allocs/op

10.09GB 40.15% 40.15%    10.09GB 40.15%  main%2ego.copy3 (inline)
8.10GB 32.22% 72.37%     8.10GB 32.22%  main%2ego.copy2 (inline)
6.94GB 27.61%   100%     6.94GB 27.61%  main%2ego.copy1 (inline)

copy1:
	251874(кол-во выполненых операций)
	5376 ns/op (ns/op - наносекунды на 1 операцию)
	35184 B/op (B/op - байт на 1 операцию)
	11 allocs/op (allocs/op - алокаций на 1 операцию)

copy3 использовал больше всего памяти 10.09GB т.к. быстрее всего: 666026 итераций × 16384 B/op = 10.09GB
Это не значит что он самый затратный
*/

func BenchmarkAllCopies(b *testing.B) {
	input := generateInput(1000)
	b.ResetTimer()

	b.Run("Copy1", func(b *testing.B) { //создает под-бенчмарк (sub-benchmark) внутри основного бенчмарка.
		for i := 0; i < b.N; i++ {
			copy1(input)
		}
	})
	b.Run("Copy2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy2(input)
		}
	})
	b.Run("Copy1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy3(input)
		}
	})

}

// в данном случае не целесообразно т.к. я хочу измерить алгоритм, а не конкурентность, добавляет шум синхронизации=>Результаты будут менее точными
func BenchmarkParallelCopy1(b *testing.B) {
	input := generateInput(1000)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Код выполняется в нескольких goroutine
			copy1(input)
		}
	})
}

// Подходит для параллельных бенчмарков
func BenchmarkChannel(b *testing.B) {
	ch := make(chan int, 100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch <- 1
			<-ch
		}
	})
}

//go test -bench . -benchmem -cpuprofile=profile
//go tool pprof profile
//(pprof) top -топ 10 операций по времени/памяти
//(pprof) web - график

//(pprof) list copy1 - сколько времени код работал на каонкретной строке, copy1-функция
//(pprof) list runtime.nextFreeFast - функция из top

//go test -bench . -benchmem -memprofile=mem_profile

//go test -bench . -benchmem -count=5  > profile_old (лучше запускать в wsl иначе проблемы с кодировкой)

//go tool pprof -http=:9191 mem_profile - тоже график но удобнее

//benchstat profile_new profile_old -сравнение

//go test -bench=BenchmarkAllCopies  -benchmem -count=5
//go test -bench=BenchmarkAllCopies -benchmem -count=5 -memprofile=mem_profileBenchmarkAllCopies
