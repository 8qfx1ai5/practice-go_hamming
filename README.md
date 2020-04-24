
# This project is an exercise to develop a simple function

The task was created from [Tino](https://github.com/pandorasNox) according to his [go-workshop](https://github.com/pandorasNox/go-workshop).

## Task

I want to have a function, which gets two strings and makes a hamming string comparison and returns the value.

### Acceptance criteria

- valid hamming result
- in case of invalid string length, the function should return -1
- passes all test-cases
- also support Chinese characters

## Steps

- define function signiture
- write unit tests for the function
- implement the function
- implement cli wrapper with one input argument for the possible isbn
- implement integration tests for the CLI
- implement input argument to pass a csv file-path with a list of possible ISBNs
- expend integration tests for the CLI with CSV file

### Hints

- [What is Hamming Code](https://de.wikipedia.org/wiki/Hamming-Code)

## The final result

### How to run

    $ go run main.go

### How to test

    $ go test main.go main_test.go
