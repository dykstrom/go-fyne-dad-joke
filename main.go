package main

import (
	"dad-joke/mvc"
)

func main() {
	view := mvc.NewView()
	mvc.NewController(view)
	view.ShowAndRun()
}
