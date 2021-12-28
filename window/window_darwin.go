//go:build darwin
// +build darwin

package window

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#import <AppKit/NSApplication.h>
#include <stdlib.h>

int Run(void) {
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

    [NSApp activateIgnoringOtherApps:YES];
    [NSApp run];
    return 0;
}

void SetWindowTitle(void * w, char * title) {
	[[(NSView*)w window] setTitle:[NSString stringWithUTF8String:title]];
}

void ShowWindow(void * w) {
	[[(NSView*)w window] makeKeyAndOrderFront:nil];
}
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/ClarkGuan/go-sciter"
)

func New(creationFlags sciter.WindowCreationFlag, rect *sciter.Rect) (*Window, error) {
	w := new(Window)
	w.creationFlags = creationFlags
	// create window
	hwnd := sciter.CreateWindow(
		creationFlags,
		rect,
		0,
		0,
		sciter.BAD_HWINDOW)

	if hwnd == sciter.BAD_HWINDOW {
		return nil, fmt.Errorf("Sciter CreateWindow failed")
	}

	w.Sciter = sciter.Wrap(hwnd)
	return w, nil
}
func (s *Window) SetTitle(title string) {
	t := C.CString(title)
	C.SetWindowTitle(unsafe.Pointer(s.GetHwnd()), t)
	C.free(unsafe.Pointer(t))
}

func (s *Window) Show() {
	C.ShowWindow(unsafe.Pointer(s.GetHwnd()))
}

func (s *Window) Run() {
	s.run()
	C.Run()
}
