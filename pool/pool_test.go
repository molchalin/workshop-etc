package main

import (
	"io"
	"log"
	"sync"
	"testing"
)

type zero struct{}

func (r zero) Read(p []byte) (n int, err error) {
	return len(p), nil
}

type discard struct{}

func (d discard) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func copyV1() {
	r := zero{}
	w := discard{}

	_, err := io.CopyBuffer(w, io.LimitReader(r, 1024), make([]byte, 128))
	if err != nil {
		log.Fatal(err)
	}
}

func copyV2() {
	r := zero{}
	w := discard{}

	buf := dataPool.Get().(*[]byte)
	_, err := io.CopyBuffer(w, io.LimitReader(r, 1024), *buf)
	if err != nil {
		log.Fatal(err)
	}
	dataPool.Put(buf)
}

func BenchmarkV1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			copyV1()
		}
	})
}

func BenchmarkV2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			copyV2()
		}
	})
}

var dataPool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 128)
		return &buf
	},
}
