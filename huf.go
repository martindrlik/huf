// huf uses Huffman method for the construction of minimum-redundancy codes.
package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	ws, err := weights(os.Args[1:])
	check(err)
	if len(ws) == 0 {
		usage()
	}
	right(ws)
	for _, w := range ws {
		fmt.Println(w)
	}
}

func right(ws Weights) {
	sort.Sort(ws)
	if len(ws) == 2 {
		ws[0].Bits = append(ws[0].Bits, false)
		ws[1].Bits = append(ws[1].Bits, true)
		left(ws[0])
		left(ws[1])
		return
	}
	w := &Weight{}
	w.u = ws[len(ws)-2]
	w.v = ws[len(ws)-1]
	w.Value = w.u.Value + w.v.Value
	nw := make(Weights, len(ws)-2, len(ws)-1)
	copy(nw, ws[:len(ws)-2])
	nw = append(nw, w)
	right(nw)
}

func left(w *Weight) {
	if w.u == nil {
		return
	}
	w.u.Bits = append(w.u.Bits, w.Bits...)
	w.v.Bits = append(w.v.Bits, w.Bits...)
	w.u.Bits = append(w.u.Bits, false)
	w.v.Bits = append(w.v.Bits, true)
	left(w.u)
	left(w.v)
}

func check(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "huf: %v\n", err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: huf s1+w1 [s2+w2...]\n")
	fmt.Fprintf(os.Stderr, "  s1, s2, etc.: symbol of alphabet\n")
	fmt.Fprintf(os.Stderr, "  w1, w2, etc.: weight of the symbol\n")
	os.Exit(2)
}
