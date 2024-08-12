package gui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ScreenSize struct {
    width int
    height int
}

const (
    RECEIVED_MESSAGES_TEXT string = `Connected.

Use ^n and ^p to cycle through the UI elements.`
    EMAIL string = "test"
    PASSWORD string = "test"
)

// loadDefaultStyle loads default colors to use for the application. This might
// change later as it would need to support different colorschemes. We also want
// to be able to dynamically change the colors from within the application.
func loadDefaultStyle() {
    tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone
    tview.Styles.ContrastBackgroundColor = tcell.NewHexColor(NORD3)
    tview.Styles.GraphicsColor = tcell.NewHexColor(NORD3)
    tview.Styles.BorderColor = tcell.NewHexColor(NORD3)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD6)
    tview.Styles.SecondaryTextColor = tcell.NewHexColor(NORD14)
    tview.Styles.MoreContrastBackgroundColor = tcell.NewHexColor(NORD6)
}

// anonymous function to toggle the contacts list when pressing the ENTER key
func toggleContactsList(node *tview.TreeNode) {
    if len(node.GetChildren()) != 0 {
	node.SetExpanded(!node.IsExpanded())
    }
}

func generateContactsArea() (g *tview.Grid, cl *tview.TreeView) {
    var contactsNode, contacts1, contacts2, contacts3 *tview.TreeNode
    var contactsList *tview.TreeView
  
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD14)
    contacts1 = tview.NewTreeNode("2050@404.city")
    contacts2 = tview.NewTreeNode("2060@404.city")
    contacts3 = tview.NewTreeNode("2077@404.city")
    contactsNode = tview.NewTreeNode("Contacts").
    AddChild(contacts1).AddChild(contacts2).AddChild(contacts3)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD6)

    contactsList = tview.NewTreeView().SetRoot(contactsNode).
    SetCurrentNode(contactsNode)
    contactsList.SetSelectedFunc(toggleContactsList) // UGLY, arg is a FUNCTION
    contactsList.SetBorder(true)

    return tview.NewGrid().SetBorders(false).
    AddItem(contactsList, 0, 0, 1, 1, 0, 0, false), contactsList
}

func generateLogin(app *tview.Application, rp *tview.Pages, cl *tview.TreeView) *tview.Grid {
    loginForm := tview.NewForm().
    AddInputField("XMPP email adress", "", 0, nil, nil).
    AddPasswordField("Password", "", 0, 0, nil)

    loginForm.AddButton("Login", func() {
	email_check_form := loginForm.GetFormItemByLabel("XMPP email adress").
	(*tview.InputField)
	email_check := email_check_form.GetText()
	pass_check_form := loginForm.GetFormItemByLabel("Password").
	(*tview.InputField)
	pass_check := pass_check_form.GetText()

	if email_check == EMAIL && pass_check == PASSWORD {
	    rp.SwitchToPage("mainPage")
	    app.SetFocus(cl)
	    return
	} else {
	    error_screen := tview.NewModal().
	    SetText("Error: wrong email or password").
	    AddButtons([]string{"OK"}).
	    SetDoneFunc(func(ButtonIndex int, ButtonLabel string) {
		if ButtonLabel == "OK" {
		    rp.SwitchToPage("loginPage")
		}
	    })
	    error_screen.SetBorder(false)
	    rp.AddAndSwitchToPage("errorPage", error_screen, true)
	}
    }).
    AddButton("Quit", func() {app.Stop()})

    loginGrid := tview.NewGrid().SetSize(3,3,-3,-3).
    AddItem(loginForm, 1, 1, 1, 1, 0, 0, true).
    SetBorders(true)

    return loginGrid
}

func generateMessageArea() (g *tview.Grid, tv *tview.TextView, ta *tview.TextArea) {
    receivedMessagesArea := tview.NewTextView().SetText(RECEIVED_MESSAGES_TEXT).
    SetScrollable(true).SetDynamicColors(true)
    receivedMessagesArea.SetBorder(true)

    plchStyle := tcell.StyleDefault
    plchStyle = plchStyle.Foreground(tcell.NewHexColor(NORD14))
    plchStyle = plchStyle.Dim(true)

    sendMessages := tview.NewTextArea().SetLabel("> ").
    SetPlaceholder("Message wtv@404.city").
    SetPlaceholderStyle(plchStyle).SetWordWrap(true).
    SetWrap(true)
    sendMessages.SetBorder(true)


    return tview.NewGrid().
    SetBorders(false).
    SetRows(0, 3).
    AddItem(receivedMessagesArea, 0, 0, 1, 1, 0, 0, false).
    AddItem(sendMessages, 1, 0, 1, 1, 0, 0, false), receivedMessagesArea,
    sendMessages
}

func generateConsole() (g *tview.Grid, ta *tview.TextArea) {

    console := tview.NewTextArea().SetLabel("Console > ").SetWrap(false)
    console.SetBorder(true)
    consoleGrid := tview.NewGrid().
    SetBorders(false).
    AddItem(console, 0, 0, 1, 1, 0, 0, false)


    return consoleGrid, console
}

func generateMainGrid(ca, ma, clia *tview.Grid) *tview.Grid {
    return tview.NewGrid().
    SetColumns(-1, -3).
    SetRows(0, 3).
    SetGap(0,0).
    AddItem(ca, 0, 0, 1, 1, 0, 0, true).
    AddItem(ma, 0, 1, 1, 1, 0, 0, false).
    AddItem(clia, 1, 0, 1, 2, 0, 0, false).
    SetBorders(false)
}

// Generates the application and the root primitive to return to main
func BuildApp() (a *tview.Application, rp *tview.Pages) {
    var app *tview.Application
    var rootPrimitive *tview.Pages
    var contactsList *tview.TreeView
    var contactsGrid *tview.Grid
    var loginGrid *tview.Grid
    var receivedMessagesArea *tview.TextView
    var sendMessages *tview.TextArea
    var messageGrid *tview.Grid
    var mainGrid *tview.Grid
    var console *tview.TextArea
    var consoleGrid *tview.Grid

    loadDefaultStyle()

    app = tview.NewApplication()
    rootPrimitive = tview.NewPages()

    contactsGrid, contactsList = generateContactsArea()

    loginGrid = generateLogin(app, rootPrimitive, contactsList)

    messageGrid, receivedMessagesArea, sendMessages = generateMessageArea()

    consoleGrid, console = generateConsole()

    uiElements := [4]*tview.Box{contactsList.Box,
    receivedMessagesArea.Box, sendMessages.Box, console.Box}
    var idx int = 0


    // Size of inner rectangle for sendMessages
    var wdt int = 0 // find a way to calculate this in realtime
    const BASE_HEIGHT int = 3
    const MAX_WANTED_LINES = 4 // Max number of lines that displays for sending messages
    const MAX_HEIGHT = BASE_HEIGHT + MAX_WANTED_LINES
    var txtLen int = 0
    var ratio int = 0

    sendMessages.SetChangedFunc(func() {
	txtLen = sendMessages.GetTextLength()

	if ratio != txtLen / wdt {
	    ratio = txtLen / wdt

	    if ratio + BASE_HEIGHT < MAX_HEIGHT {
		messageGrid.SetRows(0, BASE_HEIGHT + ratio)
	    }
	}
    })

    sendMessages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// trash listener ngl but it will do for now
	if wdt == 0 {
	    _,_,wdt,_ = sendMessages.GetRect()
	    wdt -= 4
	    return nil
	}

	if event.Key() == tcell.KeyEnter {
	    payload := sendMessages.GetText()
	    sendMessages.SetText("", true)
	    fmt.Fprintf(receivedMessagesArea, "\n\n[#5e81ac]%s:[-] %s", EMAIL, payload)
	    receivedMessagesArea.ScrollToEnd()
	    return nil
	}

	return event
    })

    console.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
	    console.SetText("", true)
	    return nil
	}

	return event
    })

    mainGrid = generateMainGrid(contactsGrid, messageGrid, consoleGrid)

    mainGrid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

	switch(event.Key()) {
	case tcell.KeyCtrlN:
	    if idx != 3 {
		idx++
	    } else {
		idx = 0
	    }
	    app.SetFocus(uiElements[idx])

	case tcell.KeyCtrlP:
	    if idx != 0 {
		idx--
	    } else {
		idx = 3
	    }
	    app.SetFocus(uiElements[idx])
	}

	return event
    })

    rootPrimitive.AddPage("loginPage", loginGrid, true, true).
    AddPage("mainPage", mainGrid, true, false)

    // This kind of works, needs some fixing and some code refactor
    test := &ScreenSize{0, 0}

    app.SetAfterDrawFunc(func(screen tcell.Screen) {
	newx, newy := screen.Size()
	if test.width != newx {
	    test.width = newx
	    test.height = newy
	    _,_,wdt,_ = sendMessages.Box.GetRect()
	    wdt -= 4
	}
    })

    return app, rootPrimitive
}
