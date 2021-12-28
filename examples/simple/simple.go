package main

import (
	"log"

	"github.com/ClarkGuan/go-sciter"
	"github.com/ClarkGuan/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("handle: %v", w.Handle)
	w.LoadFile("simple.html")
	w.SetTitle("Example")
	w.Show()
	w.Run()
}
