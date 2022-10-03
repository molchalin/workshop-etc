package main

import (
	"testing"
)

func BenchmarkCodegen(b *testing.B) {
	s := SportsmanRecord{}
	a := []string{"Stephen Curry", "74.5", "2021", "Basketball"}
	m := map[string]int{
		"Name":     0,
		"earnings": 1,
		"year":     2,
		"sport":    3,
	}
	for i := 0; i < b.N; i++ {
		err := s.Scan(a, m)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReflect(b *testing.B) {
	s := SportsmanRecord{}
	a := []string{"Stephen Curry", "74.5", "2021", "Basketball"}
	m := map[string]int{
		"Name":     0,
		"earnings": 1,
		"year":     2,
		"sport":    3,
	}
	d := Decoder{m, nil}
	for i := 0; i < b.N; i++ {
		err := d.scan(&s, a)
		if err != nil {
			b.Fatal(err)
		}
	}
}
