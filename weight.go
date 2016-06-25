package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Weight struct {
	Bits   []bool
	Value  float64
	Symbol string

	u, v *Weight
}

func (w *Weight) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s(%.2f) ", w.Symbol, w.Value)
	for _, b := range w.Bits {
		if b {
			fmt.Fprint(buf, "1")
		} else {
			fmt.Fprint(buf, "0")
		}
	}
	return buf.String()
}

func weight(s string) (*Weight, error) {
	i := strings.LastIndex(s, "+")
	if i < 0 {
		return nil, fmt.Errorf("%s: missing +", s)
	}
	f, err := strconv.ParseFloat(s[i:], 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", s, err)
	}
	return &Weight{Value: f, Symbol: s[:i]}, nil
}
