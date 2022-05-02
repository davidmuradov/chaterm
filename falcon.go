package main

import( // "fmt"
	"github.com/rivo/tview"
)

func main() {

	// function to define the title of grid elements
	newPrimitive := func(title string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(title)
	}

	contacts := newPrimitive("Contacts")
	messages := newPrimitive("Messages")

	grid := tview.NewGrid().
		SetRows(/* set height */0).
		SetColumns(0).
		SetBorders(true)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(contacts, 0, 0, 0, 0, 0, 0, false).
		AddItem(messages, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(contacts, 0, 0, 1, 1, 0, 100, false).
		AddItem(messages, 0, 1, 1, 1, 0, 100, false)

	app := tview.NewApplication().
		SetRoot(grid, true).
		Run()
	if err := app; err != nil {
		panic(err)
	}
}
