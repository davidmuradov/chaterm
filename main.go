package main

/*
import (

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)
*/


func main() {

    app, rootPrimitive := BuildApp()

    if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
    EnableMouse(false).Run(); err != nil {
	panic(err)
    }
}
