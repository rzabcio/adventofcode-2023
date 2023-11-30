# Advent of Code 2023
Just my solutions for [Advent of Code 2023](https://adventofcode.com/2023).

Uses [Cobra](https://github.com/spf13/cobra) as CLI framework and [Go-Funk](https://github.com/thoas/go-funk) for some array processing tools.

Requirements installation:
~~~~
> go mod tidy
~~~~

How to run specific puzzle (input files included in /input-files):
~~~~
> go run main.go day*.go day <day:1-25> <part:1/2> <input-file>
~~~~

Because of TDD approach, tests are also included:
~~~~
> go test
~~~~
