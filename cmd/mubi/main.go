package main

import (
	"fmt"
	"github.com/jdennes/mubi"
	"time"
)

func main() {
	// Just get a user's ratings as an example at this point
	ratings := mubi.GetRatings(7995037)
	for _, rating := range ratings {
		fmt.Printf("%s: %s\n", rating.Film.Title, rating.Film.CanonicalUrl)
		when := time.Unix(rating.Timestamp, 0)
		fmt.Printf("Rating: %d stars on %s\n", rating.Overall, when)
		fmt.Println("-----")
	}
}
