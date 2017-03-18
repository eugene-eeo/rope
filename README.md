```golang
import "github.com/eugene-eeo/rope"
```

**rope** implements the [Rope data structure](https://en.wikipedia.org/wiki/Rope_(data_structure))
for Go. A Rope is a heavyweight string that can make O(n)
operations such as concatenation cheaper. This package supports
concatenation, splitting, indexing, and rebalancing.
Example usage:

```
package main

import "fmt"
import "github.com/eugene-eeo/rope"

func main() {
    t := rope.NewLeaf("Hello").Concat(rope.NewLeaf("John!"))
    l, r := t.SplitAt(5)
    t = rope.NewNode(l, rope.NewLeaf(" there, ") r)
    fmt.Println(t.Value()) // => "Hello there, John!"
}
```
