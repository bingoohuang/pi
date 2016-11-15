package main

import "testing"

func benchmarkLeibnizπ(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		leibnizπ(float64(i))
	}
}

func BenchmarkLeibnizπ1(b *testing.B)         { benchmarkLeibnizπ(1, b) }
func BenchmarkLeibnizπ10(b *testing.B)        { benchmarkLeibnizπ(10, b) }
func BenchmarkLeibnizπ100(b *testing.B)       { benchmarkLeibnizπ(100, b) }
func BenchmarkLeibnizπ1000(b *testing.B)      { benchmarkLeibnizπ(1000, b) }
func BenchmarkLeibnizπ10000(b *testing.B)     { benchmarkLeibnizπ(10000, b) }
func BenchmarkLeibnizπ100000(b *testing.B)    { benchmarkLeibnizπ(100000, b) }
func BenchmarkLeibnizπ1000000(b *testing.B)   { benchmarkLeibnizπ(1000000, b) }
func BenchmarkLeibnizπ10000000(b *testing.B)  { benchmarkLeibnizπ(10000000, b) }
func BenchmarkLeibnizπ100000000(b *testing.B) { benchmarkLeibnizπ(100000000, b) }
