package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// https://www.kaggle.com/datasets/darinhawley/forbes-high-paid-athletes-19902021

var exampleBytes []byte

func init() {
	f, err := os.Open("forbesathletesv2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	exampleBytes, err = io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
}

func newExampleReader() io.Reader {
	return bytes.NewReader(exampleBytes)
}

// csvgen
type SportsmanRecord struct {
	Name     string
	Year     int     `csv:"year"`
	Sport    string  `csv:"sport"`
	Earnings float64 `csv:"earnings"`
}

func main() {
	r := csv.NewReader(newExampleReader())
	m, err := MostValuable(r, 2005)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m = %+v\n", m)
}

func MostValuable(r *csv.Reader, year int) (result SportsmanRecord, err error) {
	d, err := NewDecoder(r)
	if err != nil {
		return SportsmanRecord{}, err
	}

	var v SportsmanRecord
	var found bool
	for {
		err = d.Scan(&v)
		if err == io.EOF {
			break
		} else if err != nil {
			return SportsmanRecord{}, err
		}

		if v.Year == year {
			found = true
			if v.Earnings > result.Earnings {
				result = v
			}
		}
	}
	if !found {
		return SportsmanRecord{}, fmt.Errorf("not found")
	}
	return result, nil
}
