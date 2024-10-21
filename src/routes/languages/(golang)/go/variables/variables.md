---
title: Variables in Go
---

## {title}

A variable is a container that stores a value.

### Examples

Create a `main.go` file and add the following contents below

```go
// main.go
package main

import "fmt"

func main() {
	name := "Jack"
	fmt.Println(name)
}
```

write `go run main.go` on the terminal, the below should be your output.

```bash
 Jack
```

**name** is a variable _(container)_ that holds a value of `"Jack"` _(without the quotes)_
