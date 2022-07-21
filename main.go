package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("hello there")
	a := app.New()
	w := a.NewWindow("My new title")
	// Re-Sizing fyne app window
	w.Resize(fyne.NewSize(1200, 700))
	w.SetContent(widget.NewLabel("Hello World!"))
	btn := widget.NewButton("Button", func() {
		fmt.Println("Hi there")
	})

	check_1 := widget.NewCheck("Title of check", func(b bool) {
		fmt.Println(fmt.Sprintf("my check box  value %t", b))
	})

	w.SetContent(btn)
	w.SetContent(check_1)
	w.ShowAndRun()
}
