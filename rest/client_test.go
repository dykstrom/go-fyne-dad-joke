package rest

import (
	"fmt"
	"testing"
)

func TestGetRandomJoke(t *testing.T) {
	joke, err := GetRandomJoke()
	if err != nil {
		t.Errorf("Failed to get random joke: %v", err)
		return
	}
	if joke.Status != 200 {
		t.Errorf("Got HTTP Status %d; want 200", joke.Status)
		return
	}
	fmt.Println(joke.Joke)
}

func TestGetRandomJokeBySearchTerm(t *testing.T) {
	joke, err := GetRandomJokeBySearchTerm("dog")
	if err != nil {
		t.Errorf("Failed to search for joke: %v", err)
		return
	}
	if joke.Status != 200 {
		t.Errorf("Got HTTP Status %d; want 200", joke.Status)
		return
	}
	fmt.Println(joke.Joke)
}

func TestShouldNotFindJokeAboutQwerty(t *testing.T) {
	joke, err := GetRandomJokeBySearchTerm("qwerty")
	if err == nil {
		t.Errorf("Got '%s'; want error", joke.Joke)
		return
	}
}
