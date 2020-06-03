# Title
Simple equation parser(recursive descent parser) using Golang

## Description
Main rationale is to parse patterns in the prioritized way so multiplication and division would work
before addition and subtraction but after parenthesis and functions. To implement that functions are
evaluated from the bottom up so to speak.

Recursive descent parser means substituting the patterns with other patterns until no substitution can be made. After that
what we are left with is a final result(or AST tree if you wish to construct one).

# Testing
Simply `go test` and that's it

# Examples
Example inputs supported:
```
10--1           // 11
10 --1          // 11
2+ 5            // 7
3 - 3*2         // -3
3 *    (3 + 2)  // 15
sin(0)          // 0
sin(90)         // 0.99
sin(90-180)     // -0.99
```
