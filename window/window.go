package window

import (
	"github.com/ClarkGuan/go-sciter"
	// "runtime"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}
