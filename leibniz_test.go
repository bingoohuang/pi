package main

import "testing"

func benchmarkLeibnizFormulaforπ(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		leibnizFormulaforπ(float64(i))
	}
}

func BenchmarkLeibnizFormulaforπ1(b *testing.B)         { benchmarkLeibnizFormulaforπ(1, b) }
func BenchmarkLeibnizFormulaforπ10(b *testing.B)        { benchmarkLeibnizFormulaforπ(10, b) }
func BenchmarkLeibnizFormulaforπ100(b *testing.B)       { benchmarkLeibnizFormulaforπ(100, b) }
func BenchmarkLeibnizFormulaforπ1000(b *testing.B)      { benchmarkLeibnizFormulaforπ(1000, b) }
func BenchmarkLeibnizFormulaforπ10000(b *testing.B)     { benchmarkLeibnizFormulaforπ(10000, b) }
func BenchmarkLeibnizFormulaforπ100000(b *testing.B)    { benchmarkLeibnizFormulaforπ(100000, b) }
func BenchmarkLeibnizFormulaforπ1000000(b *testing.B)   { benchmarkLeibnizFormulaforπ(1000000, b) }
func BenchmarkLeibnizFormulaforπ10000000(b *testing.B)  { benchmarkLeibnizFormulaforπ(10000000, b) }
func BenchmarkLeibnizFormulaforπ100000000(b *testing.B) { benchmarkLeibnizFormulaforπ(100000000, b) }
