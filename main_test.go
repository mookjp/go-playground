package main

import "testing"

func BenchmarkArrayAndMap_Array(b *testing.B) {
	targets := make([]int, 0)
	for i := 0; i < 100000000; i++ {
		targets = append(targets, i)
	}
	b.ResetTimer()
	
	for _, v := range targets {
		if v == 100000000 {
			return
		}
	}
}

func BenchmarkArrayAndMap_Map(b *testing.B) {
	targets := make(map[string]string)
	for i := 0; i < 100000000; i++ {
		targets[string(i)] = string(i)
	}
	b.ResetTimer()
	if targets["100000000"] == "100000000" {
		return
	}
}
