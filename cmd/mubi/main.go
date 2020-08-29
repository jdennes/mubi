package main

import (
	"flag"
	"fmt"
	"github.com/jdennes/mubi"
	"os"
	"strings"
	"time"
)

func main() {
	ratingsCmd := flag.NewFlagSet("ratings", flag.ExitOnError)
	ratingsUserId := ratingsCmd.Int64("userid", 0, "Mubi.com user ID")

	watchlistCmd := flag.NewFlagSet("watchlist", flag.ExitOnError)
	watchlistUserId := watchlistCmd.Int64("userid", 0, "Mubi.com user ID")

	favouriteFilmsCmd := flag.NewFlagSet("favourite-films", flag.ExitOnError)
	favouriteFilmsUserId := favouriteFilmsCmd.Int64("userid", 0, "Mubi.com user ID")

	if len(os.Args) < 2 {
		fmt.Println("no subcommand provided")
		os.Exit(1)
	}

	api := mubi.NewMubiAPI()

	switch os.Args[1] {
	case "ratings":
		ratingsCmd.Parse(os.Args[2:])
		printRatings(*api, *ratingsUserId)
	case "watchlist":
		watchlistCmd.Parse(os.Args[2:])
		printWatchlist(*api, *watchlistUserId)
	case "favourite-films":
		favouriteFilmsCmd.Parse(os.Args[2:])
		printFavouriteFilms(*api, *favouriteFilmsUserId)
	default:
		fmt.Println("unexpected subcommand provided")
		os.Exit(1)
	}
}

func printRatings(api mubi.MubiAPI, userId int64) {
	ratings := api.GetRatings(userId, 1, 20)
	for _, rating := range ratings {
		fmt.Printf("%s (%d) - %s\n", rating.Film.Title, rating.Film.Year, rating.Film.CanonicalUrl)

		var directorNames []string
		for _, director := range rating.Film.Directors {
			directorNames = append(directorNames, director.Name)
		}
		fmt.Printf("Directed by %s\n", strings.Join(directorNames, ", "))

		when := time.Unix(rating.Timestamp, 0)
		fmt.Printf("Rated %d/5 stars on %s\n", rating.Overall, when)
		fmt.Printf("-----\n")
	}
}

func printWatchlist(api mubi.MubiAPI, userId int64) {
	watchlist := api.GetWatchlist(userId, 1, 20)
	for _, item := range watchlist {
		fmt.Printf("%s (%d) - %s\n", item.Film.Title, item.Film.Year, item.Film.CanonicalUrl)

		var directorNames []string
		for _, director := range item.Film.Directors {
			directorNames = append(directorNames, director.Name)
		}
		fmt.Printf("Directed by %s\n", strings.Join(directorNames, ", "))

		when := time.Unix(item.Timestamp, 0)
		fmt.Printf("Added to watchlist on %s\n", when)
		fmt.Printf("-----\n")
	}
}

func printFavouriteFilms(api mubi.MubiAPI, userId int64) {
	favourites := api.GetFavouriteFilms(userId, 1, 20)
	for _, fav := range favourites {
		fmt.Printf("%s (%d) - %s\n", fav.Fannable.Film.Title, fav.Fannable.Film.Year, fav.Fannable.Film.CanonicalUrl)

		var directorNames []string
		for _, director := range fav.Fannable.Film.Directors {
			directorNames = append(directorNames, director.Name)
		}
		fmt.Printf("Directed by %s\n", strings.Join(directorNames, ", "))

		when := time.Unix(fav.Timestamp, 0)
		fmt.Printf("Added to favourites on %s\n", when)
		fmt.Printf("-----\n")
	}
}
