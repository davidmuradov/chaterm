package main

import( 
	"github.com/rivo/tview"
    "fmt"
)

// Git test
func main() {
    fmt.Println("test")
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
