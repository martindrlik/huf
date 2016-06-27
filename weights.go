package main

import (
	"math"
)

type Weights []*Weight

func (ws Weights) Len() int {
	return len(ws)
}

func (ws Weights) Less(i, j int) bool {
	return ws[i].Value > ws[j].Value
}

func (ws Weights) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

// H returns entropy.
func (ws Weights) H() float64 {
	h := 0.0
	for _, w := range ws {
		h += w.Value * math.Log2(w.Value)
	}
	return -h
}

// Avg returns code word average length.
func (ws Weights) Avg() float64 {
	avg := 0.0
	for _, w := range ws {
		avg += float64(len(w.Bits)) * w.Value
	}
	return avg
}

func weights(s []string) (Weights, error) {
	ws := make(Weights, 0, len(s))
	for _, s := range s {
		w, err := weight(s)
		if err != nil {
			return nil, err
		}
		ws = append(ws, w)
	}
	return ws, nil
}
