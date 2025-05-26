# seblang-interpreter

A simple interpreter for the Seblang programming language. Written for educational purposes following "Writing an Interpreter in Go" by Thorsten Ball.

Take a look at example programs in the [examples](./examples/) directory to see how Seblang works.

## Usage

To run the interpreter, you can use the following command:

```bash
go run main.go program.seb
```

You can also build the interpreter and run it directly:

```bash
go build -o seblang-interpreter main.go
./seblang-interpreter program.seb
```

To try Seblang our in the REPL, simply omit the program file:

```bash
go run main.go

# or

./seblang-interpreter
```

## Installation

### Prerequisites

- Go 1.24 or later

### Steps

The simplest way to install the interpreter is using `go install`:

```bash
go install github.com/sendelivery/seblang-interpreter@latest
```

This will install the `seblang-interpreter` binary in your `$GOPATH/bin` directory.
You can then run it from anywhere in your terminal.
