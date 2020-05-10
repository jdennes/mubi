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
	userId, page, perPage := int64(7995037), 1, 20
	ratings := api.GetRatings(userId, page, perPage)

	for _, rating := range ratings {
		fmt.Printf("%s (%d) - Rated %d/5\n", rating.Film.Title, rating.Film.Year, rating.Overall)
	}
}
```
