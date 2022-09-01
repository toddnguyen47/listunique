# listunique
Make sure a list has unique items. Go implementaton

## Available functions

* `Add(elem T)`
* `AddAll(elems ...T)`
* `GetList() []T`
* `GetSize() int`
* `Update(index int, elem T) bool`
* `Delete(elem T) bool`

## Examples

```go
package main

import (
	"fmt"

	"github.com/toddnguyen47/listunique/pkg/listunique"
)

func main() {
	uniqueList := listunique.NewUniqueList[string]()
	uniqueList.Add("hello")
	uniqueList.AddAll("world", "hello", "number", "is", "42")

	list1 := uniqueList.GetList()

	fmt.Println(list1)

	// Output:
	// [hello world number is 42]
}
```
