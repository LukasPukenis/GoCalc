package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	type Result struct {
		input  string
		result float64
	}

	testcases := []Result{
		{"10--1", 11},
		{"10 --1", 11},
		{"2 + 5", 7},
		{"3 - 3 * 2", -3},
		{"3 * (3 + 2)", 15},
		{"sin(0)", 0},
		{"sin(90)", 0.99},
		{"sin(90-180)", -0.99},
	}

	for _, testcase := range testcases {
		solver := NewSolver()
		assert.InDelta(t, testcase.result, solver.Solve(testcase.input), 0.01)
	}
}
