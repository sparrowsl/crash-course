---
title: Go (Golang)
---

# {title}

**Go** is a general-purpose programming language with clean syntax and advance features, including concurrency and a useful language
in a wide variety of software domains and used to build tools like CLI apps, Web API, etc.

This guide provides some support you need to get started with the language, with short, simple and easy to read sections
that build on each other _(hopefully)_.

> This section assumes you already have Go installed on your system.

Many companies have started using Go because of its simplicity, ease of use, performance, low barrier of entry, and powerful tooling.


## Hello World in Go

Create a `main.go` file and add the following contents below
```go
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

write `go run main.go` on the terminal, the below should be your output.
```bash
 Hello, World!
```
