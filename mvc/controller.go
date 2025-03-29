package mvc

import (
	"dad-joke/rest"
	"fmt"
	"strings"
)

func NewController(view *View) {
	// Create a closure with a reference to the view
	search := func() { handleSearchAction(view) }

	view.TextField.OnSubmitted = func(_ string) { search() }
	view.Button.OnTapped = search
}

func handleSearchAction(view *View) {
	searchTerm := strings.TrimSpace(view.TextField.Text)

	// Call REST client in a goroutine to not block the GUI
	go func() {
		var joke *rest.Joke
		var err error

		if searchTerm == "" {
			joke, err = rest.GetRandomJoke()
		} else {
			joke, err = rest.GetRandomJokeBySearchTerm(searchTerm)
		}
		if err != nil {
			view.TextArea.SetText(fmt.Sprintf("Failed to get joke: %v", err))
			return
		}

		// Apparently widgets can be updated from any goroutine
		view.TextArea.SetText(joke.Joke)
	}()
}
