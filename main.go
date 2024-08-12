package main

import "github.com/davidmuradov/falcon/gui"

func main() {

    app, rootPrimitive := gui.BuildApp()

    if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
    EnableMouse(false).Run(); err != nil {
	panic(err)
    }
}
