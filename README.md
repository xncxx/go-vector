# Go Vector Implementation

This repository accompanies a blog post about implementing a generic Vector (dynamic array) datatype in Go. The goal is to demonstrate how to build a resizable, type-safe container similar to C++'s `std::vector` or Java's `ArrayList`, using Go generics.

## Features

- Generic vector implementation using Go 1.18+ generics
- Dynamic resizing (automatic capacity management)
- Methods for push, pop, insert, prepend, delete, and access by index
- Example usage in `main.go` with output showing vector operations

## File Structure

- `vector/vector.go`: The core implementation of the generic `Vector[T]` type
- `main.go`: Example usage and demonstration of the vector's API
- `go.mod`: Go module definition

## Example Usage

```go
s := vector.NewVector[int]()
s.Push(1)
s.Push(2)
s.Insert(1, 3)
s.Prepend(0)
s.Pop()
s.Delete(1)
fmt.Println(s.GetItems()) // Output: [0 2]
```

## Running the Example

To run the demonstration in `main.go`:

```sh
go run main.go
```

You will see output showing when elements are added or removed, along with the vector's size and capacity after each operation.

## Blog Post

For a detailed explanation of the implementation, design decisions, and Go generics, see [the accompanying blog post](https://www.thecodechameleon.io/articles/go_vector/).

## License

MIT
