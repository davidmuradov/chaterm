package main

import (
    "fmt"
    "time"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

// Default Style
func loadDefaultStyle() {

    tview.Styles.PrimitiveBackgroundColor = tcell.ColorNone
    tview.Styles.ContrastBackgroundColor = tcell.NewHexColor(0x4c566a)
    tview.Styles.GraphicsColor = tcell.NewHexColor(0x4c566a)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(0xffffff)
    tview.Styles.SecondaryTextColor = tcell.NewHexColor(0xa3be8c)
    tview.Styles.MoreContrastBackgroundColor = tcell.NewHexColor(0xffffff)
}

// Perhaps useless function
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
  
    return tview.NewGrid().SetBorders(true).
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
    SetBorders(true).
    SetRows(-45,-1).
    AddItem(rma, 0, 0, 1, 1, 0, 0, false).
    AddItem(sm, 1, 0, 1, 1, 0, 0, false)
}

func generateMainGrid(ca, ma, clia *tview.Grid) *tview.Grid {

    return tview.NewGrid().
    SetColumns(-1, -3).
    SetRows(-9, -1).
    SetGap(1,1).
    AddItem(ca, 0, 0, 1, 1, 0, 0, true).
    AddItem(ma, 0, 1, 1, 1, 0, 0, false).
    AddItem(clia, 1, 0, 1, 2, 0, 0, false).
    SetBorders(false)
}

func main() {

    // App widgets
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

    const RECEIVED_MESSAGES_TEXT = `Here we receive messages from our contacts ...

Pressing <esc> will change the focus on another part of the app.

Press <esc> to move between contacts, the message area and the console.`

    // Login information
    const EMAIL string = "test@test.com"
    const PASSWORD string = "test"

    loadDefaultStyle()

    // Main Primitive to display
    app = tview.NewApplication()
    rootPrimitive = tview.NewPages()

    // Contacts infos
    // ! Check for more efficient way of setting colors for nodes
    tview.Styles.PrimaryTextColor = tcell.NewRGBColor(163, 190, 140)
    contacts1 = allocateNewTreeNode("2050@404.city")
    contacts2 = allocateNewTreeNode("2060@404.city")
    contacts3 = allocateNewTreeNode("2077@404.city")
    contactsNode = allocateNewTreeNode("Contacts").
    AddChild(contacts1).AddChild(contacts2).AddChild(contacts3)
    tview.Styles.PrimaryTextColor = tcell.NewHexColor(0xffffff)

    contactsList = tview.NewTreeView().SetRoot(contactsNode).
    SetCurrentNode(contactsNode)
    contactsList.SetSelectedFunc(toggleContactsList) // UGLY, arg is a FUNCTION

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

    sendingMessages = tview.NewTextArea().SetLabel("> ")

    messageArea = generateMessageArea(receivedMessagesArea, sendingMessages)

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
	    app.SetFocus(receivedMessagesArea)
	    return nil
	}
	return event
    })
    // Escape key should setfocus to the TextArea when into TextView widget
    receivedMessagesArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
	    fmt.Fprintf(receivedMessagesArea, "\n\n[#5e81ac]%s:[-] %s", EMAIL, payload)
	    receivedMessagesArea.ScrollToEnd()
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
    mainGrid = generateMainGrid(contactsArea, messageArea, consoleArea)

    // Add Pages
    rootPrimitive.AddPage("loginPage", loginGrid, true, true).
    AddPage("mainPage", mainGrid, true, false)

    // Run app
    if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
    EnableMouse(false).Run(); err != nil {
	panic(err)
    }
}
