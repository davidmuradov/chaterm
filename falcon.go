package main

import (
	//"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Theme struct {
	PrimitiveBackgroundColor    tcell.Color // Main background color for primitives.
	ContrastBackgroundColor     tcell.Color // Background color for contrasting elements.
	MoreContrastBackgroundColor tcell.Color // Background color for even more contrasting elements.
	BorderColor                 tcell.Color // Box borders.
	TitleColor                  tcell.Color // Box titles.
	GraphicsColor               tcell.Color // Graphics.
	PrimaryTextColor            tcell.Color // Primary text.
	SecondaryTextColor          tcell.Color // Secondary text (e.g. labels).
	TertiaryTextColor           tcell.Color // Tertiary text (e.g. subtitles, notes).
	InverseTextColor            tcell.Color // Text on primary-colored backgrounds.
	ContrastSecondaryTextColor  tcell.Color // Secondary text on ContrastBackgroundColor-colored backgrounds.
}

func main() {
	app := tview.NewApplication()

	// Main Primitive to display
	rootPrimitive := tview.NewPages()

	// Login Form
	/* When login button is pressed, we have choice between
	switching to the already existing main page UI, or we
	can also call AddAndSwitchToPage() to create the main
	page UI on the spot and switch to it when the contact
	informations are verified on the server
	*/
	loginForm := tview.NewForm().
	AddInputField("XMPP email adress", "", 15, nil, nil).
	AddPasswordField("Password", "", 15, 0, nil).
	AddButton("Login", func() {rootPrimitive.SwitchToPage("mainPage")}).
	AddButton("Quit", func() {app.Stop()})

	// Login Grid
	loginGrid := tview.NewGrid().
	SetSize(3,3,-3,-3).
	AddItem(loginForm, 1, 1, 1, 1, 0, 0, true).
	SetBorders(true)

	/* Next objective is to setup Focus. Users will be able to
	choose which windows to focus on, giving them a color hint as to 
	which widget has focus and to use them 
	*/

	// Main Grid elements
	contactsList := tview.NewBox().SetTitle("Contacts").SetBorder(true)
	messageArea := tview.NewBox().SetTitle("Messages").SetBorder(true)
	consoleDebug := tview.NewBox().SetTitle("Console").SetBorder(true)

	// Main Grid
	mainGrid := tview.NewGrid().
	SetColumns(-1, -3).
	SetRows(-9, -1).
	AddItem(contactsList, 0, 0, 1, 1, 0, 0, true).
	AddItem(messageArea, 0, 1, 1, 1, 0, 0, false).
	AddItem(consoleDebug, 1, 0, 1, 2, 0, 0, false).
	SetBorders(true)

	// Add Pages
	rootPrimitive.AddPage("loginPage", loginGrid, true, true).
	AddPage("mainPage", mainGrid, true, false)

	// app needs to be of type tview.Primitive inside app.SetRoot()
	// form item is a valid primitive (Pages is used in this case)
	if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).Run(); err != nil {
		panic(err)
	}
}
