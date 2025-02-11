## Key Features
- **Statically typed & compiled**: Ensures type safety and performance.
- **Garbage collection**: Efficient memory management.
- **Concurrency support**: Goroutines and channels for lightweight threading.
- **Fast compilation**: Optimized for quick builds.
- **Standard library**: Comprehensive and well-structured.
- **Cross-platform**: Supports multiple operating systems.


## Package Management
- Go uses **modules** (`go mod init project_name`).
- Package installation via `go get`.

## Use Cases
- **Web Development** (Gin, Echo frameworks)
- **Networking & System Programming**
- **Cloud Services** (Docker, Kubernetes)
- **Microservices Architecture**

## Resources to Learn Go
- [Go Official Docs](https://golang.org/doc/)
- [Tour of Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)

# Packages and Modules in Go

## Understanding Packages

In Go, a **package** is a collection of multiple `.go` files that are grouped together to serve a specific purpose. Each Go file within a package can contain functions, variables, types, and other code components that can be reused within the package or accessed by other packages if exported.

- Every Go program is made up of at least one package.
- The package name is declared at the top of each Go file using the `package` keyword.
- The `main` package is a special package in Go, as it serves as the entry point for executable applications.
- Other packages are typically used for organizing code into reusable modules, such as `fmt`, `net/http`, and `os`.

### Example of a Package
```go
// File: mypackage/example.go
package mypackage

import "fmt"

func SayHello() {
    fmt.Println("Hello from mypackage!")
}
```

In this example, `mypackage` contains a function `SayHello` that can be used in other files that import `mypackage`.

## Understanding Modules

A **module** in Go is a collection of related packages that are managed together as a unit. Modules provide a way to organize Go projects and manage dependencies effectively. When we initialize a Go project, we are essentially initializing a new module.

### Key Features of Modules:
- A module is identified by a module path, which serves as its unique name.
- Modules are initialized using `go mod init` followed by the module name.
- Dependencies of a module are managed using `go.mod` and `go.sum` files.
- The `go get` command is used to fetch external dependencies and update module requirements.

### Initializing a Go Module
To start a new Go module, run:
```sh
go mod init example.com/mymodule
```
This command creates a `go.mod` file with the module name `example.com/mymodule`.

### Example of a `go.mod` File
```go
module example.com/mymodule

go 1.20

require (
    github.com/some/package v1.2.3
)
```

Here, `example.com/mymodule` is our module name, and it depends on `github.com/some/package` version `1.2.3`.

## Summary
- **Packages** group related `.go` files together and help organize code.
- **Modules** are a higher-level structure that groups related packages together.
- Initializing a project using `go mod init` creates a new module and sets up dependency management.



