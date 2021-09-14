package main

import "testing"

const word = "Hello, Go!"

func appendWords(s []string, n int) []string {
	for i := 0; i < n; i++ {
		s = append(s, word)
	}
	return s
}

func benchmarkEchoNaive(b *testing.B, size int) {
	var args []string
	args = appendWords(args, size)
	for i := 0; i < b.N; i++ {
		echoNaive(args)
	}
}

func BenchmarkEchoNaive10(b *testing.B)    { benchmarkEchoNaive(b, 10) }
func BenchmarkEchoNaive100(b *testing.B)   { benchmarkEchoNaive(b, 100) }
func BenchmarkEchoNaive1000(b *testing.B)  { benchmarkEchoNaive(b, 1000) }
func BenchmarkEchoNaive10000(b *testing.B) { benchmarkEchoNaive(b, 10000) }

func benchmarkEchoJoin(b *testing.B, size int) {
	var args []string
	args = appendWords(args, size)
	for i := 0; i < b.N; i++ {
		echoJoin(args)
	}
}

func BenchmarkEchoJoin10(b *testing.B)    { benchmarkEchoJoin(b, 10) }
func BenchmarkEchoJoin100(b *testing.B)   { benchmarkEchoJoin(b, 100) }
func BenchmarkEchoJoin1000(b *testing.B)  { benchmarkEchoJoin(b, 1000) }
func BenchmarkEchoJoin10000(b *testing.B) { benchmarkEchoJoin(b, 10000) }
