//go:build darwin

package vips

// #cgo CFLAGS: -I ./include
// #cgo CFLAGS: -I ./include/glib-2.0
// #cgo LDFLAGS: -L/opt/homebrew/Cellar/vips/8.14.3/lib -L/opt/homebrew/Cellar/glib/2.76.4/lib -L/opt/homebrew/opt/gettext/lib -lvips -lgio-2.0 -lgobject-2.0 -lglib-2.0 -lintl
import "C"
