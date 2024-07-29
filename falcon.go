package main

import (
    "fmt"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

const (

    // POLAR NIGHT
    NORD0 = 0x2e3440 // dark
    NORD1 = 0x3b4252 // brighter dark
    NORD2 = 0x434c5e // brigther brighter dark
    NORD3 = 0x4c566a // birghtest dark

    // SNOW STORM
    NORD4 = 0xd8dee9 // white
    NORD5 = 0xe5e9f0 // brighter white
    NORD6 = 0xeceff4 // brightest white

    // FROST
    NORD7 = 0x8fbcbb // turquoise
    NORD8 = 0x88c0d0 // lightblue
    NORD9 = 0x81a1c1 // semi darkblue
    NORD10 = 0x5e81ac // darkblue

    // AURORA
    NORD11 = 0xbf616a // red
    NORD12 = 0xd08770 // orange
    NORD13 = 0xebcb8b // yellow
    NORD14 = 0xa3be8c // green
    NORD15 = 0xb48ead // purple
)

func ternary_not_equal_int(val, comp, ifTrue, ifFalse int) int {

    if(val != comp) {return ifTrue} else {return ifFalse}
}

func loadDefaultStyle() {

    tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone
    tview.Styles.ContrastBackgroundColor = tcell.NewHexColor(NORD3)
    tview.Styles.GraphicsColor = tcell.NewHexColor(NORD3)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD6)
    tview.Styles.SecondaryTextColor = tcell.NewHexColor(NORD14)
    tview.Styles.MoreContrastBackgroundColor = tcell.NewHexColor(NORD6)
}

func allocateNewTreeNode(contactName string) *tview.TreeNode {

    return tview.NewTreeNode(contactName)
}

// anonymous function to toggle the contacts list
func toggleContactsList(node *tview.TreeNode) {

    if len(node.GetChildren()) != 0 {

	node.SetExpanded(!node.IsExpanded())
    }
}

func generateContactsArea(contactsList *tview.TreeView) *tview.Grid {
  
    return tview.NewGrid().SetBorders(false).
    AddItem(contactsList, 0, 0, 1, 1, 0, 0, false)
}

func generateLoginGrid(loginForm *tview.Form) *tview.Grid {

    return tview.NewGrid().SetSize(3,3,-3,-3).
    AddItem(loginForm, 1, 1, 1, 1, 0, 0, true).
    SetBorders(true)
}

func generateReceivedMessagesArea(s string) *tview.TextView {

    return tview.NewTextView().SetText(s).SetScrollable(true).
    SetDynamicColors(true)
}

func generateMessageArea(rma *tview.TextView, sm *tview.TextArea) *tview.Grid {

    return tview.NewGrid().
    SetBorders(false).
    SetRows(-43,-3).
    AddItem(rma, 0, 0, 1, 1, 0, 0, false).
    AddItem(sm, 1, 0, 1, 1, 0, 0, false)
}

func generateMainGrid(ca, ma, clia *tview.Grid) *tview.Grid {

    return tview.NewGrid().
    SetColumns(-1, -3).
    SetRows(-13, -1).
    SetGap(1,1).
    AddItem(ca, 0, 0, 1, 1, 0, 0, true).
    AddItem(ma, 0, 1, 1, 1, 0, 0, false).
    AddItem(clia, 1, 0, 1, 2, 0, 0, false).
    SetBorders(false)
}

func main() {

    var app *tview.Application
    var rootPrimitive *tview.Pages
    var contactsNode, contacts1, contacts2, contacts3 *tview.TreeNode
    var contactsList *tview.TreeView
    var contactsArea *tview.Grid
    var loginForm *tview.Form
    var loginGrid *tview.Grid
    var receivedMessagesArea *tview.TextView
    var sendingMessages *tview.TextArea
    var messageArea *tview.Grid
    var mainGrid *tview.Grid
    var consoleDebug *tview.TextArea

    const RECEIVED_MESSAGES_TEXT = `Here we receive messages from our contacts.

Use ^j and ^k to cycle through the UI elements.`

    const EMAIL string = "test"
    const PASSWORD string = "test"

    loadDefaultStyle()

    app = tview.NewApplication()
    rootPrimitive = tview.NewPages()

    // ! Check for more efficient way of setting colors for nodes
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD14)
    contacts1 = allocateNewTreeNode("2050@404.city")
    contacts2 = allocateNewTreeNode("2060@404.city")
    contacts3 = allocateNewTreeNode("2077@404.city")
    contactsNode = allocateNewTreeNode("Contacts").
    AddChild(contacts1).AddChild(contacts2).AddChild(contacts3)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(NORD6)

    contactsList = tview.NewTreeView().SetRoot(contactsNode).
    SetCurrentNode(contactsNode)
    contactsList.SetSelectedFunc(toggleContactsList) // UGLY, arg is a FUNCTION
    contactsList.SetBorder(true).SetBorderColor(tcell.NewHexColor(NORD3))

    contactsArea = generateContactsArea(contactsList)

    // Login Form
    /* When login button is pressed, we have choice between
    switching to the already existing main page UI, or we
    can also call AddAndSwitchToPage() to create the main
    page UI on the spot and switch to it when the contact
    informations are verified on the server
    */

    loginForm = tview.NewForm().
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
	    error_screen.SetBorder(false)
	    rootPrimitive.AddAndSwitchToPage("errorPage", error_screen, true)
	}

    }).
    AddButton("Quit", func() {app.Stop()})

    loginGrid = generateLoginGrid(loginForm)

    receivedMessagesArea = generateReceivedMessagesArea(RECEIVED_MESSAGES_TEXT)
    receivedMessagesArea.SetBorder(true)
    receivedMessagesArea.SetBorderColor(tcell.NewHexColor(NORD3))

    sendingMessages = tview.NewTextArea().SetLabel("> ").
    SetPlaceholder("Message wtv@404.city")
    sendingMessages.SetBorder(true)
    sendingMessages.SetBorderColor(tcell.NewHexColor(NORD3))

    messageArea = generateMessageArea(receivedMessagesArea, sendingMessages)

    consoleDebug = tview.NewTextArea().SetLabel("Console > ").SetWrap(false)
    consoleDebug.SetBorder(true).SetBorderColor(tcell.NewHexColor(NORD3))
    consoleArea := tview.NewGrid().
    SetBorders(false).
    AddItem(consoleDebug, 0, 0, 1, 1, 0, 0, false)

    uiElements := [4]*tview.Box{contactsList.Box,
    receivedMessagesArea.Box, sendingMessages.Box, consoleDebug.Box}
    var idx int = 0

    app.SetFocus(uiElements[idx])
    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

	switch(event.Key()) {

	case tcell.KeyCtrlJ:
	    idx += ternary_not_equal_int(idx, 3, 1, -3)
	    app.SetFocus(uiElements[idx])
	case tcell.KeyCtrlK:
	    idx += ternary_not_equal_int(idx, 0, -1, 3)
	    app.SetFocus(uiElements[idx])
	}
	return event
    })

    sendingMessages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
	    payload := sendingMessages.GetText()
	    sendingMessages.SetText("", true)
	    fmt.Fprintf(receivedMessagesArea, "\n\n[#5e81ac]%s:[-] %s", EMAIL, payload)
	    receivedMessagesArea.ScrollToEnd()
	    return nil
	}
	return event
    })

    mainGrid = generateMainGrid(contactsArea, messageArea, consoleArea)

    rootPrimitive.AddPage("loginPage", loginGrid, true, true).
    AddPage("mainPage", mainGrid, true, false)

    if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
    EnableMouse(false).Run(); err != nil {
	panic(err)
    }
}
