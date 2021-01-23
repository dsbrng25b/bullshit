# bullshit

Go port of the [bullshit](http://man.cat-v.org/9front/1/bullshit) command of [9front](http://9front.org/).

## Usage
```go
package main

import (
	"fmt"

	"github.com/dvob/bullshit"
)

func main() {
	fmt.Println(bullshit.Get())
}
```
