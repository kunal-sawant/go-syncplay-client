package main

import (
	"fmt"

	"github.com/rivo/tview"
)

var serverAddress string = "ws://localhost:8080/connect"
var vlcInstallLocation string = "c:\\program files\\videolan\\vlc\\vlc.exe"

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)

	form := tview.NewForm().
		AddInputField("Server Address", serverAddress, 50, nil, nil).
		AddInputField("VLC Location", vlcInstallLocation, 50, nil, nil).
		AddButton("Start", func() {
			StartServer(textView)
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(textView, 0, 1, false)

	form.SetBorder(true).SetTitle("Go Syncplay Client").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}

func StartServer(textView *tview.TextView) {
	textView.Clear()
	fmt.Fprintf(textView, "[green]Server Connected[-]")
}
