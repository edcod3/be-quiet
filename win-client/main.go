package main

import (
	"fmt"
	"net/http"
	"os"
	"syscall"
	"unsafe"
)

// Declare the MessageBox function from User32.dll
var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procMessageBoxW = user32.NewProc("MessageBoxW")
	procSendMessage = user32.NewProc("SendMessage")
	procFindWindow  = user32.NewProc("FindWindowW")
	WM_COMMAND      = 0x11
	MIN_ALL         = 419
	MIN_ALL_UNDO    = 416
)

// MessageBox function to call Windows API
func MessageBox(text string, caption string) {
	// Convert Go strings to wide strings (UTF-16)
	textPtr, _ := syscall.UTF16PtrFromString(text)
	captionPtr, _ := syscall.UTF16PtrFromString(caption)
	// Call MessageBoxW function from User32.dll
	procMessageBoxW.Call(0, uintptr(unsafe.Pointer(textPtr)), uintptr(unsafe.Pointer(captionPtr)), 0)
}

func SendMessage(hWnd uintptr, Msg int32, wParam uintptr, lParam uintptr) {
	procSendMessage.call(hWnd, Msg, wParam, lParam)
}

func FindWindow(className string, windowName string) uintptr {
	lpWindowName := uintptr(0)
	lpClassName, _ := syscall.UTF16PtrFromString(className)
	if windowName != "" {
		lpWindowName, _ = syscall.UTF16PtrFromString(windowName)
	}

	return procFindWindow.call(lpClassName, lpWindowName)
}

func minimizeWindows() {
	lHwnd := FindWindow("Shell_TrayWnd", "")
	SendMessage(lHwnd, int32(WM_COMMAND), uintptr(MIN_ALL), uintptr(0))
}

// Handler for the web API
func handler(w http.ResponseWriter, r *http.Request) {
	// Call the MessageBox API on every request
	MessageBox("You have received a BE QUIET alert. Please do not make anymore loud noises!", "BE QUIET!")
	// Respond with a simple message
	fmt.Fprintf(w, "BE QUIET alert received! You may rest peacefully now.")
}

func main() {
	// Get the port from CLI arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./win-client <port>")
		return
	}
	port := os.Args[1]

	// Start the web server
	http.HandleFunc("/", handler)
	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
