//go:build !windows && !darwin

package vips

// #cgo pkg-config: vips
import "C"
