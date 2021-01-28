package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/jdennes/mubi"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	ratingsCmd := flag.NewFlagSet("ratings", flag.ExitOnError)
	ratingsUserId := ratingsCmd.Int64("userid", 0, "Mubi.com user ID")
	ratingsPage := ratingsCmd.Int("page", 1, "Results page number")
	ratingsPerPage := ratingsCmd.Int("per-page", 20, "Number of results per page")
	ratingsExportForLetterboxd := ratingsCmd.Bool("export-for-letterboxd", false, "If true, print output in CSV format for Letterboxd importer")

	watchlistCmd := flag.NewFlagSet("watchlist", flag.ExitOnError)
	watchlistUserId := watchlistCmd.Int64("userid", 0, "Mubi.com user ID")
	watchlistPage := watchlistCmd.Int("page", 1, "Results page number")
	watchlistPerPage := watchlistCmd.Int("per-page", 20, "Number of results per page")

	favouriteFilmsCmd := flag.NewFlagSet("favourite-films", flag.ExitOnError)
	favouriteFilmsUserId := favouriteFilmsCmd.Int64("userid", 0, "Mubi.com user ID")
	favouriteFilmsPage := favouriteFilmsCmd.Int("page", 1, "Results page number")
	favouriteFilmsPerPage := favouriteFilmsCmd.Int("per-page", 20, "Number of results per page")

	if len(os.Args) < 2 {
		fmt.Println("no subcommand provided")
		os.Exit(1)
	}

	api := mubi.NewMubiAPI()

	switch os.Args[1] {
	case "ratings":
		ratingsCmd.Parse(os.Args[2:])
		printRatings(*api, *ratingsUserId, *ratingsPage, *ratingsPerPage, *ratingsExportForLetterboxd)
	case "watchlist":
		watchlistCmd.Parse(os.Args[2:])
		printWatchlist(*api, *watchlistUserId, *watchlistPage, *watchlistPerPage)
	case "favourite-films":
		favouriteFilmsCmd.Parse(os.Args[2:])
		printFavouriteFilms(*api, *favouriteFilmsUserId, *favouriteFilmsPage, *favouriteFilmsPerPage)
	case "films-of-the-day":
		printFilmsOfTheDay(*api)
	default:
		fmt.Println("unexpected subcommand provided")
		os.Exit(1)
	}
}

func printRatings(api mubi.MubiAPI, userId int64, page int, perPage int, exportForLetterboxd bool) {
	ratings := api.GetRatings(userId, page, perPage)
	if exportForLetterboxd == true {
		printRatingsForLetterboxd(ratings)
	} else {
		printRatingsStandard(ratings)
	}
}

func printRatingsStandard(ratings []mubi.Rating) {
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

func printRatingsForLetterboxd(ratings []mubi.Rating) {
	// Print CSV output for Letterboxd importer as defined here:
	// https://letterboxd.com/about/importing-data/
	lines := [][]string{{"Title", "Year", "Directors", "Rating", "WatchedDate"}}
	for _, rating := range ratings {
		var directorNames []string
		for _, director := range rating.Film.Directors {
			directorNames = append(directorNames, director.Name)
		}
		when := time.Unix(rating.Timestamp, 0)

		line := []string{
			rating.Film.Title,
			strconv.Itoa(rating.Film.Year),
			strings.Join(directorNames, ", "),
			strconv.Itoa(rating.Overall),
			when.Format("2006-01-02"),
		}
		lines = append(lines, line)
	}

	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(lines)

	if err := writer.Error(); err != nil {
		log.Fatalln("Error writing CSV:", err)
	}
}

func printWatchlist(api mubi.MubiAPI, userId int64, page int, perPage int) {
	watchlist := api.GetWatchlist(userId, page, perPage)
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

func printFavouriteFilms(api mubi.MubiAPI, userId int64, page int, perPage int) {
	favourites := api.GetFavouriteFilms(userId, page, perPage)
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

func printFilmsOfTheDay(api mubi.MubiAPI) {
	filmsOfTheDay := api.GetFilmsOfTheDay()
	for _, fotd := range filmsOfTheDay {
		fmt.Printf("%s (%d) - %s\n", fotd.Film.Title, fotd.Film.Year, fotd.Film.WebUrl)
		fmt.Printf("Expires %s\n", fotd.ExpiresAt)
		fmt.Printf("-----\n")
	}
}
