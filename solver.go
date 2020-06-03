package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type Solver struct {
	index  int
	tokens []string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) peek() string {
	return s.tokens[s.index]
}

func (s *Solver) consume() string {
	val := s.tokens[s.index]
	s.index++

	return val
}

func (s *Solver) accept(sym string) bool {
	if s.index < len(s.tokens) && s.tokens[s.index] == sym {
		s.consume()
		return true
	}

	return false
}

// + and -
func (s *Solver) expr() float64 {
	term := s.term()

	for {
		if s.accept("+") {
			term += s.term()
		} else if s.accept("-") {
			term -= s.term()
		} else {
			break
		}
	}

	return term
}

// * and /
func (s *Solver) term() float64 {
	number := s.number()

	for {
		if s.accept("*") {
			number *= s.number()
		} else if s.accept("/") {
			number /= s.number()
		} else {
			break
		}
	}

	return number
}

func ang2rad(n float64) float64 {
	return n * 3.14 / 180.0
}

func perform_function(name string, arg float64) float64 {
	switch name {
	case "sin":
		return math.Sin(ang2rad(arg))
	case "cos":
		return math.Cos(ang2rad(arg))
	default:
		fmt.Println("Unrecognized function:", name)
		return 0
	}
}

func (s *Solver) number() float64 {
	negator := 1.0
	if s.peek() == "-" {
		negator = -1.0
		s.consume()
	}

	// handles unary minus as well if its right in front of the number
	num, err := strconv.ParseFloat(s.peek(), 64)

	if err == nil {
		s.consume()
		return num * negator
	} else if s.accept("(") {
		expr := s.expr()

		s.accept(")")
		return expr * negator
	} else {
		// expect function name
		fun_names := []string{"sin", "cos"}
		fun_name := ""

		for _, name := range fun_names {
			if s.peek() == name {
				fun_name = name
				s.consume()
				break
			}
		}

		if fun_name == "" {
			fmt.Println("Unexpected function token:", s.peek())
			return 0
		}

		if !s.accept("(") {
			fmt.Println("( expected after function name")
			return 0
		}

		// parse argument which is just an expression
		arg := s.expr()

		if !s.accept(")") {
			fmt.Println(") expected after function args")
			return 0
		}

		return perform_function(fun_name, arg)
	}

	return 0
}

func (s *Solver) Solve(input string) float64 {
	tokens, err := tokenize(input)
	if err != nil {
		log.Fatal(err)
	}

	s.tokens = tokens
	s.index = 0

	return s.expr()
}

func main() {

}
