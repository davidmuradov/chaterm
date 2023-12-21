package main

import (
	"fmt"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
/*
type Theme struct {
	//PrimitiveBackgroundColor    tcell.Color // Main background color for primitives.
	//ContrastBackgroundColor     tcell.Color // Background color for contrasting elements.
	//MoreContrastBackgroundColor tcell.Color // Background color for even more contrasting elements.
	//BorderColor                 tcell.Color // Box borders.
	//TitleColor                  tcell.Color // Box titles.
	//GraphicsColor               tcell.Color // Graphics.
	//PrimaryTextColor            tcell.Color // Primary text.
	//SecondaryTextColor          tcell.Color // Secondary text (e.g. labels).
	TertiaryTextColor           tcell.Color // Tertiary text (e.g. subtitles, notes).
	InverseTextColor            tcell.Color // Text on primary-colored backgrounds.
	ContrastSecondaryTextColor  tcell.Color // Secondary text on ContrastBackgroundColor-colored backgrounds.
}
*/

// Login information
const email string = "test@test.com"
const password string = "test"

// Default Style
func loadDefaultStyle () {
	tview.Styles.PrimitiveBackgroundColor = tcell.NewRGBColor(46, 52, 64)
	tview.Styles.ContrastBackgroundColor = tcell.NewRGBColor(76, 86, 106)
	tview.Styles.GraphicsColor = tcell.NewRGBColor(76, 86, 106)
	tview.Styles.SecondaryTextColor = tcell.NewRGBColor(163, 190, 140)
	tview.Styles.MoreContrastBackgroundColor = tcell.NewRGBColor(46, 52, 64)
}

func main() {

	// Load the default style (Nord-ish)
	loadDefaultStyle()

	app := tview.NewApplication()
	// Main Primitive to display
	rootPrimitive := tview.NewPages()

	// Contacts area
	// Contacts infos
	// ! Check for more efficient way of setting colors for nodes
	tview.Styles.PrimaryTextColor = tcell.NewRGBColor(163, 190, 140)
	contacts1 := tview.NewTreeNode("2050@404.city")
	contacts2 := tview.NewTreeNode("2060@404.city")
	contacts3 := tview.NewTreeNode("2077@404.city")
	contactsNode := tview.NewTreeNode("Contacts").
	AddChild(contacts1).
	AddChild(contacts2).
	AddChild(contacts3)
	tview.Styles.PrimaryTextColor = tcell.ColorWhite

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
	AddInputField("XMPP email adress", "", 0, nil, nil).
	AddPasswordField("Password", "", 0, 0, nil)

	loginForm.AddButton("Login", func() {
		email_check_form := loginForm.GetFormItemByLabel("XMPP email adress").(*tview.InputField)
		email_check := email_check_form.GetText()
		pass_check_form := loginForm.GetFormItemByLabel("Password").(*tview.InputField)
		pass_check := pass_check_form.GetText()

		if email_check == email && pass_check == password {
			rootPrimitive.SwitchToPage("mainPage")
			app.SetFocus(contactsList)
			return
		} else {
			error_screen := tview.NewModal().
			SetText("Error: wrong email or password").
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(ButtonIndex int, ButtonLabel string) {
				if ButtonLabel == "OK" {
					rootPrimitive.SwitchToPage("loginPage")
				}
			})
		rootPrimitive.AddAndSwitchToPage("errorPage", error_screen, true)
		}
		
	}).
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
	SetScrollable(true).
	SetDynamicColors(true)

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

	// Focus operations are hardcoded for now. Focus begins on the contactsList widget,
	// and by pressing escape, we change focus to the TextArea widget to type messages,
	// and pressing escape once again changes focus to the console. A final press on
	// escape brings focus back to the contactsList widget
	// Escape key should setfocus to the TextView widget when into the contacts widget
	contactsList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEscape {
		app.SetFocus(receivedMessages)
		return nil
	}
	return event
	})
	// Escape key should setfocus to the TextArea when into TextView widget
	receivedMessages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
	} else if event.Key() == tcell.KeyEnter {
		payload := sendingMessages.GetText()
		//payload_length := sendingMessages.GetTextLength()
		sendingMessages.SetText("", true)
		fmt.Fprintf(receivedMessages, "\n\n%s: %s", email, payload)
		receivedMessages.ScrollToEnd()
		time.Sleep(30 * time.Millisecond)
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
	SetGap(1,1).
	AddItem(contactsArea, 0, 0, 1, 1, 0, 0, true).
	AddItem(messageArea, 0, 1, 1, 1, 0, 0, false).
	AddItem(consoleArea, 1, 0, 1, 2, 0, 0, false).
	SetBorders(false)

	// Add Pages
	rootPrimitive.AddPage("loginPage", loginGrid, true, true).
	AddPage("mainPage", mainGrid, true, false)
	//fmt.Println(email_valid)

	// Run app
	if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
	EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
