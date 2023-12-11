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

	// Login Form
	loginForm := tview.NewForm().
	AddInputField("XMPP email adress", "rat@404.city", 25, nil, nil).
	AddPasswordField("Password", "", 25, 0, nil).
	AddButton("Login", nil).
	AddButton("Quit", func() {app.Stop()})

	// Main application's structure
	rootPrimitive := tview.NewPages().
	AddPage("loginPage", loginForm, true, true)

	// app needs to be of type tview.Primitive inside app.SetRoot()
	// form item is a valid primitive
	if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).Run(); err != nil {
		panic(err)
	}
}
