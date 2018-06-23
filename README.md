```golang
import "github.com/eugene-eeo/rope"
```

Implements the [Rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure))
for Go. A Rope is a heavyweight string that can make O(n)
operations such as concatenation cheaper. This package supports
concatenation, splitting, indexing, and rebalancing. Mainly an
exercise in making the most of Go's tooling ecosystem, therefore
not Production Ready™. Example usage:

```
package main

import "fmt"
import "github.com/eugene-eeo/rope"

func main() {
    t := rope.L("Hello").Concat(rope.L("John!"))
    l, r := t.SplitAt(t.Index(byte("J")))
    t = rope.Concat(l, rope.L(" there, "), r)
    fmt.Println(t.Value()) // => "Hello there, John!"
}
```
