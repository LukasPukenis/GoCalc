Simple equation parser(recursive descent parser) using Golang

Main rationale is to parse patterns in the prioritized way so multiplication and division would work
before addition and subtraction but after parenthesis and functions. To implement that functions are
evaluated from the bottom up so to speak.

Recursive descent parser means substituting the patterns with other patterns until no substitution can be made. After that
what we are left with is a final result(or AST tree if you wish to construct one).
