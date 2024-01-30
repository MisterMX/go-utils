package fs

import (
	"io/fs"
	"os"
)

var _ fs.FS = &OSFS{}

// OSFS is a dead-simple implementation of [io/fs.FS] that just wraps around
// [os.Open].
//
// It is usefull when one needs to work with Go's abstract filesystem interface
// that should work on every platform since [os.DirFS] does not work
// universally (see [https://github.com/golang/go/issues/44279]).
type OSFS struct{}

// NewOSFS creates a new OSFS instance.
func NewOSFS() *OSFS {
	return &OSFS{}
}

func (OSFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}
