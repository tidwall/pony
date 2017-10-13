# 🐴 `PONY` 🌈

Make rainbow text !!

<img src="https://i.imgur.com/pDbvLG0.gif" height=40>

## Getting Started

### Installing

To start using Pony, install Go and run `go get`:

```sh
$ go get -u github.com/tidwall/pony
```

This will retrieve the library.

### Usage

There's only the one function:

```go
func Text(phrase string, offset int) string
```

Here're two examples for the price of one.

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/tidwall/pony"
)

func main() {
	// Create a single line of dashes that are VERY wonderful to gaze upon.
	fmt.Printf("%s\n", pony.Text(strings.Repeat("-", 80), 0))
	
	// Create a rainbow effect that will straight up dazzle!
	for i := 0; ; i++ {
		fmt.Printf("\r%s ",
			pony.Text("Mute - thy coronation - Meek - my Vive le roi", i),
		)
		time.Sleep(time.Second / 15)
	}
}
```




## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

Pony source code is available under the MIT [License](/LICENSE).

