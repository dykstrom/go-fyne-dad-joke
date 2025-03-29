package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Joke struct {
	Joke   string // Joke text
	Status int    // HTTP status
}

type JokeList struct {
	TotalJokes int    `json:"total_jokes"` // Total number of jokes found
	Results    []Joke // The jokes in this paged result
}

const baseUrl = "https://icanhazdadjoke.com/"

// GetRandomJoke returns a random joke about anything.
func GetRandomJoke() (*Joke, error) {
	joke := Joke{}
	err := getAndMapData(createRequest(), baseUrl, &joke)
	return &joke, err
}

// GetRandomJokeBySearchTerm returns a random joke about the search term.
// If there are no such jokes, this function returns an error.
func GetRandomJokeBySearchTerm(searchTerm string) (*Joke, error) {
	// Fetch a single joke in order to find out the total number of jokes
	jokeList, err := getJokeList(searchTerm, 0)
	if err != nil {
		return nil, err
	}
	if jokeList.TotalJokes == 0 {
		return nil, errors.New("No jokes about '" + searchTerm + "' found.")
	}

	// Choose a random joke
	index := rand.IntN(jokeList.TotalJokes)

	// Fetch the chosen joke
	jokeList, err = getJokeList(searchTerm, index)
	if err != nil {
		return nil, err
	}

	return &Joke{jokeList.Results[0].Joke, 200}, nil
}

// getJokeList fetches a list of jokes for the given search term.
func getJokeList(searchTerm string, index int) (*JokeList, error) {
	jokeList := JokeList{}
	request := createRequest().
		SetQueryParams(map[string]string{
			"term":  searchTerm,
			"page":  fmt.Sprintf("%d", index+1),
			"limit": "1",
		})
	err := getAndMapData(request, baseUrl+"search", &jokeList)
	return &jokeList, err
}

// createRequest creates a basic REST request.
func createRequest() *resty.Request {
	return resty.
		New().
		R().
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", "go-fyne-dad-joke (https://github.com/dykstrom/go-fyne-dad-joke)")
}

// getAndMapData executes the HTTP GET request and, if successful, unmarshals the data into the provided value v.
func getAndMapData(request *resty.Request, url string, v any) error {
	response, err := request.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode() != http.StatusOK {
		return errors.New(fmt.Sprintf("Server returned status %d", response.StatusCode()))
	}
	return json.Unmarshal(response.Body(), v)
}
