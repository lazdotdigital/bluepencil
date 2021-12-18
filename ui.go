package main

import (
	"fmt"
	"os"

	"github.com/webview/webview"
)

func initUI(bd bindData) error {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Blue Pencil - " + bd.path)
	w.SetSize(800, 600, webview.HintNone)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	w.Bind("keyDownBind", bd.keyDown)
	w.Bind("ctrlKeyDownBind", bd.ctrlKeyDown)
	w.Bind("getBufferStringBind", bd.getBufferString)

	w.Navigate(fmt.Sprintf("file://%v/ui/index.html", wd))
	w.Run()
	return nil
}
