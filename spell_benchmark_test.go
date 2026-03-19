package main

import (
	"testing"
)

func BenchmarkCast_ExactMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cast("accio")
	}
}

func BenchmarkCast_Similar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cast("acccio")
	}
}

func BenchmarkCast_NoMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cast("abcdefgxyz")
	}
}

func BenchmarkLevenshtein_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levenshtein("accio", "acccio")
	}
}

func BenchmarkLevenshtein_Long(b *testing.B) {
	a := "expectopatronum"
	c := "expecto patronumm"
	for i := 0; i < b.N; i++ {
		levenshtein(a, c)
	}
}
