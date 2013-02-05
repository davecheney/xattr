package xattr

import (
	"syscall"
)

const (
	prefix = ""
)

// No-op on Darwin (Mac).
func stripPrefix(s []string) []string {
	return s
}

func isNotExist(err *XAttrError) bool {
	return err.Err == syscall.ENOATTR
}
