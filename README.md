# mubi

Do things with https://mubi.com.

Example usage:

```go
package main

import (
	"fmt"
	"github.com/jdennes/mubi"
)

func main() {
	api := mubi.NewMubiAPI()
	ratings := api.GetRatings(7995037)

	for _, rating := range ratings {
		fmt.Printf("%s (%d) - Rated %d/5\n", rating.Film.Title, rating.Film.Year, rating.Overall)
	}
}
```
