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

func setxattr(path string, name string, value []byte) (errno int) {
	var _p0 unsafe.Pointer
	if len(value) > 0 {
		_p0 = unsafe.Pointer(&value[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(name))), uintptr(_p0), uintptr(len(value)), 0, 0)
	errno = int(e1)
	return
}

