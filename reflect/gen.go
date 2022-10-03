package main

import (
	"encoding/csv"
)

type CSVScanner interface {
	Scan([]string, map[string]int) error
}

type Decoder struct {
	names map[string]int
	r     *csv.Reader
}

func NewDecoder(r *csv.Reader) (*Decoder, error) {
	fieldNames, err := r.Read()
	if err != nil {
		return nil, err
	}
	names := make(map[string]int, len(fieldNames))
	for i, n := range fieldNames {
		names[n] = i
	}
	return &Decoder{names, r}, nil
}

func (d *Decoder) Scan(v any) error {
	records, err := d.r.Read()
	if err != nil {
		return err
	}
	if scanner, ok := v.(CSVScanner); ok {
		return scanner.Scan(records, d.names)
	}
	return d.scan(v, records)
}
