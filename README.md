# Golang starter

Golang basic syntax and examples.

## Run

    go run main.go
    go build main.go
    // => main.exe

    // Run files in the same package
    go run main.go <file_1>.go

## Init Go module

    go mod init <module_name>

Init go module could enable to run `go test` to test the code.

## Golang Syntax

### Pointers

- `&variable`: Get the memory address of the value this variable is pointing at.
- `*pointer`: Get the value this memory address is pointing at.

### Array and slice

- Array: Fixed length, cannot be changed, primitive data type
- Slice: Dynamic length

### Value Types and Reference Types

| Value Types                       | Reference Types                             |
| --------------------------------- | ------------------------------------------- |
| int, float, string, bool, structs | slices, maps, channels, pointers, functions |
