```go
package main

import "fmt"

func main() {
    var a [5]int
    for i := 0; i < 10; i++ {
        a[i] = i
    }
    fmt.Println(a)
}
```