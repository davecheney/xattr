package xattr

import (
	"syscall"
	"unsafe"
)

var noError = syscall.Errno(0)

// ssize_t getxattr(const char *path, const char *name, void *value, size_t size);
func getxattr(path string, name string, value *byte, size int) (resSize int, e error) {
	r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
		uintptr(unsafe.Pointer(value)),
		uintptr(size), 0, 0)
	resSize = int(r0)
	if e1 != noError {
		e = e1
	}
	return
}

// ssize_t listxattr(const char *path, char *list, size_t size);
func listxattr(path string, list *byte, size int) (resSize int, e error) {
	r0, _, e1 := syscall.Syscall(syscall.SYS_LISTXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(list)),
		uintptr(size))
	resSize = int(r0)
	if e1 != noError {
		e = e1
	}
	return
}

// int setxattr(const char *path, const char *name, const void *value, size_t size, int flags);
func setxattr(path string, name string, value *byte, size int) (e error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
		uintptr(unsafe.Pointer(value)),
		uintptr(size), 0, 0)
	if e1 != noError {
		e = e1
	}
	return
}

// int removexattr(const char *path, const char *name);
func removexattr(path string, name string) (e error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR,
		uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
		0)
	if e1 != noError {
		e = e1
	}
	return
}
