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

### Map and Struct

| Map                                                                                                | Struct                                                        |
| -------------------------------------------------------------------------------------------------- | ------------------------------------------------------------- |
| **Reference** type                                                                                 | **Value** type                                                |
| Key-value collection, All **keys** must be the same type, All **values** **must be the same type** | Group different data, **Values** **can be different** types   |
| Use to represent a collection of related properties                                                | Use to represent a "thing" with a lot of different properties |
| Need to know all the keys at compile time                                                          | Don't need to know all the fields at compile time             |
| Key are **indexed**, **can be iterated over**                                                      | Fields are **not indexed**, cannot be iterated over           |
