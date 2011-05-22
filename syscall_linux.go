package xattr

import (
	"syscall"
	"unsafe"
)

func getxattr(path string, name string, value *byte, size int) (n int, errno int) {
	r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0)
	n = int(r0)
	errno = int(e1)
	return
}

func listxattr(path string, namebuf *byte, size int) (n int, errno int) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_LISTXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(namebuf)), uintptr(size))
	n = int(r0)
	errno = int(e1)
	return
}

func setxattr(path string, name string, value *byte, size int) (errno int) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), uintptr(unsafe.Pointer(value)), uintptr(size), 0, 0)
	errno = int(e1)
	return
}

func removexattr(path string, name string) (errno int) {
	_, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), 0)
	errno = int(e1)
	return
}

