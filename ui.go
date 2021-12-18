package main

import (
	"fmt"
	"os"

	"github.com/webview/webview"
)

func initUI(path string, bd bindData) error {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Blue Pencil - " + path)
	w.SetSize(800, 600, webview.HintNone)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	w.Bind("keyDownBind", bd.keyDown)

	w.Navigate(fmt.Sprintf("file://%v/ui/index.html", wd))
	w.Run()
	return nil
}
