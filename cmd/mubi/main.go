package main

import (
	"fmt"
  "github.com/jdennes/mubi"
)

func main() {
  // Just get a user's ratings as an example at this point
  ratings := mubi.GetRatings(7995037)
	for _, rating := range ratings {
		fmt.Printf("%s: %s\n", rating.Film.Title, rating.Film.CanonicalUrl)
		fmt.Println("Rating:", rating.Overall)
		fmt.Println("-----")
	}
}
