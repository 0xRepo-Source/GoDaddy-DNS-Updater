package main

import (
	"syscall"
)

// For Windows: build with -ldflags "-H windowsgui" to hide the console window
// This file provides Windows-specific build constraints

func init() {
	// This will hide the console window when built with the windowsgui flag
	_ = syscall.GetStdHandle
}
