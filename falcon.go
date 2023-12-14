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

	// Contacts area
	// Contacts infos
	contacts1 := tview.NewTreeNode("2050@404.city")
	contacts2 := tview.NewTreeNode("2060@404.city")
	contacts3 := tview.NewTreeNode("2077@404.city")
	contactsNode := tview.NewTreeNode("Contacts").
	AddChild(contacts1).
	AddChild(contacts2).
	AddChild(contacts3)

	// The actual contacts list
	contactsList := tview.NewTreeView().SetRoot(contactsNode).
	SetCurrentNode(contactsNode)

	contactsArea := tview.NewGrid().
	SetBorders(true).
	AddItem(contactsList, 0, 0, 1, 1, 0, 0, false)

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
	AddButton("Login", func() {rootPrimitive.SwitchToPage("mainPage")
	app.SetFocus(contactsList)}).
	AddButton("Quit", func() {app.Stop()})

	// Login Grid
	loginGrid := tview.NewGrid().
	SetSize(3,3,-3,-3).
	AddItem(loginForm, 1, 1, 1, 1, 0, 0, true).
	SetBorders(true)

	// Main Grid elements
	// Messaging area
	receivedMessages := tview.NewTextView().
	SetText("Here we receive messages from our contacts ...\n\nPressing <esc> will change the focus on another part of the app.\n\nPress <esc> to move between contacts, the message area and the console.").
	SetTextColor(tcell.ColorYellow)

	sendingMessages := tview.NewTextArea().
	SetLabel("> ")

	messageArea := tview.NewGrid().
	SetBorders(true).
	SetRows(-45,-1).
	AddItem(receivedMessages, 0, 0, 1, 1, 0, 0, false).
	AddItem(sendingMessages, 1, 0, 1, 1, 0, 0, false)

	// The useless console box
	consoleDebug := tview.NewTextArea().SetLabel("Console > ")
	consoleArea := tview.NewGrid().
	SetBorders(true).
	AddItem(consoleDebug, 0, 0, 1, 1, 0, 0, false)

	/* Next objective is to setup Focus. Users will be able to
	choose which windows to focus on, giving them a color hint as to 
	which widget has focus and to use them 
	*/
	// Focus operations are hardcoded for now. Focus begins on the contactsList widget,
	// and by pressing escape, we change focus to the TextArea widget to type messages,
	// and pressing escape once again changes focus to the console. A final press on
	// escape brings focus back to the contactsList widget
	// Escape key should setfocus to the TextArea widget when into the contacts widget
	contactsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEscape {
		app.SetFocus(sendingMessages)
		return nil
	}
	return event
	})
	// Escape key should setfocus to the ConsoleBox when into the sendingMessages widget
	sendingMessages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEscape {
		app.SetFocus(consoleDebug)
		return nil
	}
	return event
	})
	// Escape key should setfocus to the contactsList when into the console widget
	consoleDebug.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEscape {
		app.SetFocus(contactsList)
		return nil
	}
	return event
	})

	// Main Grid
	mainGrid := tview.NewGrid().
	SetColumns(-1, -3).
	SetRows(-9, -1).
	SetGap(2,2).
	AddItem(contactsArea, 0, 0, 1, 1, 0, 0, true).
	AddItem(messageArea, 0, 1, 1, 1, 0, 0, false).
	AddItem(consoleArea, 1, 0, 1, 2, 0, 0, false).
	SetBorders(false)

	// Add Pages
	rootPrimitive.AddPage("loginPage", loginGrid, true, true).
	AddPage("mainPage", mainGrid, true, false)

	// app needs to be of type tview.Primitive inside app.SetRoot()
	// form item is a valid primitive (Pages is used in this case)
	if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
	EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
