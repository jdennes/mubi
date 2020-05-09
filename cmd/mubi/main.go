package main

import (
	"flag"
	"fmt"
	"github.com/jdennes/mubi"
	"net/http"
	"os"
	"time"
)

func main() {
	ratingsCmd := flag.NewFlagSet("ratings", flag.ExitOnError)
	ratingsUserId := ratingsCmd.Int64("userid", 0, "Mubi.com user ID")
	watchlistCmd := flag.NewFlagSet("watchlist", flag.ExitOnError)
	watchlistUserId := watchlistCmd.Int64("userid", 0, "Mubi.com user ID")

	if len(os.Args) < 2 {
		fmt.Println("no subcommand provided")
		os.Exit(1)
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	switch os.Args[1] {
	case "ratings":
		ratingsCmd.Parse(os.Args[2:])
		ratingsApi := mubi.RatingsAPI{&client}
		ratings := ratingsApi.GetRatings(*ratingsUserId)
		for _, rating := range ratings {
			fmt.Printf("%s (%d) - %s\n", rating.Film.Title, rating.Film.Year, rating.Film.CanonicalUrl)
			when := time.Unix(rating.Timestamp, 0)
			fmt.Printf("Rated %d stars on %s\n", rating.Overall, when)
			fmt.Printf("-----\n")
		}
	case "watchlist":
		watchlistCmd.Parse(os.Args[2:])
		watchlistApi := mubi.WatchlistAPI{&client}
		watchlist := watchlistApi.GetWatchlist(*watchlistUserId)
		for _, item := range watchlist {
			fmt.Printf("%s (%d) - %s\n", item.Film.Title, item.Film.Year, item.Film.CanonicalUrl)
			when := time.Unix(item.Timestamp, 0)
			fmt.Printf("Added to watchlist on %s\n", when)
			fmt.Printf("-----\n")
		}

	default:
		fmt.Println("unexpected subcommand provided")
		os.Exit(1)
	}
}
