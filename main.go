package main

func main() {

    app, rootPrimitive := BuildApp()

    if err := app.SetRoot(rootPrimitive, true).SetFocus(rootPrimitive).
    EnableMouse(false).Run(); err != nil {
	panic(err)
    }
}
