package main

import (
	"math"
)

type Function struct {
	name string
	op   func(args ...float64) float64
}

var functions []Function

func init() {
	functions = []Function{
		{"sin", func(args ...float64) float64 { return math.Sin(ang2rad(args[0])) }},
		{"cos", func(args ...float64) float64 { return math.Cos(ang2rad(args[0])) }},
	}
}

func findFunction(name string) (func(args ...float64) float64, bool) {
	for _, fun := range functions {
		if fun.name == name {
			return fun.op, true
		}
	}

	return nil, false
}

func ang2rad(n float64) float64 {
	return n * 3.14 / 180.0
}
