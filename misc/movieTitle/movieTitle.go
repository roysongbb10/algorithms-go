package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// Movie defines fields associated a data item in the response of Movies response
type Movie struct {
	Title string `json:"Title"`
	Year int `json:"Year"`
	ID string `json:"imdbID"`
}

// Movies defines fields associated with the Movies response
type Movies struct {
	Page string `json:"page"`
	PerPage int `json:"per_page"`
	Totle int `json:"total"`
	TotlePages int `json:"total_pages"`
	Data []Movie `json:"data"`
}

/*
 * Complete the function below.
 */
func getMovieTitles(substr string) []string {
	// Get first page
    totalPages, movies := getMovieTitlesByPage(substr, 1)
    if movies == nil {
        fmt.Println("Failed to fetch data")
        return nil
    }

	// Get subsquent pages
    for i:=1; i<totalPages; i++ {
        _, next := getMovieTitlesByPage(substr, i + 1)
        movies = append(movies, next...);
    }

    // Sort the title array
	sort.Strings(movies)
	
    return movies
}

func getMovieTitlesByPage(substr string, pageNumber int) (int, []string) {
    url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/movies/search/?Title=%s&page=%d", substr, pageNumber)
	
	// Retrieve movies page from web
	response, error := http.Get(url)
    if error != nil {
        panic(error)
	}
	
	// Close the response once we return from this function
	defer response.Body.Close()

	// Check the status code to make sure we have received a proper response
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch. Error code %d", response.StatusCode)
		return 0, nil
	}
	
	// Initialize an empty Movies struct
	movies := Movies{}

	// Decode response data to Movies struct
	err := json.NewDecoder(response.Body).Decode(&movies)
	if err != nil {
		panic(err)
	}
	
	// Append title of movies into an array
	var titles []string
	for _, movie := range movies.Data {
		titles = append(titles, movie.Title)
	}

	return movies.TotlePages, titles
}

func main() {
    var substr string
    fmt.Scanln(&substr)
    res := getMovieTitles(substr)
    
    for _, v := range res {
        fmt.Println(v)
    }
}
