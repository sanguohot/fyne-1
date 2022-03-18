// Package desktop provides desktop specific driver functionality.
package desktop

import (
	"fyne.io/fyne/v2"
)

// Driver represents the extended capabilities of a desktop driver
type DriverSgh interface {
	// Create a new borderless window that is centered on screen
	CreateWindowSgh(title string, decorate, transparent, centered bool) fyne.Window
}
