package mvc

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	TextField *widget.Entry
	Button    *widget.Button
	TextArea  *widget.Entry
	fyne.Window
}

func NewView() *View {
	// Must create app before everything else
	app := app.New()

	// Create widgets
	label := widget.NewLabel("Search for dad joke")
	textField := widget.NewEntry()
	textField.SetPlaceHolder("Enter a search term or leave blank for a random joke")
	button := widget.NewButton("Get joke", nil)
	textArea := widget.NewMultiLineEntry()
	textArea.Wrapping = fyne.TextWrapWord

	// Layout containers
	topContainer := container.NewVBox(
		label,
		container.NewBorder(nil, nil, nil, button, textField),
	)
	mainContainer := container.NewBorder(topContainer, nil, nil, nil, textArea)

	// Create window
	window := app.NewWindow("icanhazdadjoke.com")
	window.SetContent(container.New(layout.NewCustomPaddedLayout(0, 10, 10, 10), mainContainer))
	window.Resize(fyne.NewSize(500, 350))

	return &View{textField, button, textArea, window}
}
