package xattr

import (
	"syscall"
	"unsafe"
)

// http://www.opensource.apple.com/source/xnu/xnu-2050.18.24/bsd/kern/syscalls.master

var noError = syscall.Errno(0)

// user_ssize_t getxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options);
func getxattr(path string, attrname string, value *byte, size int, position uint32, options int) (resSize int, e error) {
        r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(attrname))),
                uintptr(unsafe.Pointer(value)),
                uintptr(size), uintptr(position), uintptr(options))
        resSize = int(r0)
        if e1 != noError {
                e = e1
        }
        return
}

// user_ssize_t listxattr(user_addr_t path, user_addr_t namebuf, size_t bufsize, int options);
func listxattr(path string, namebuf *byte, bufsize int, options int) (resSize int, e error) {
        r0, _, e1 := syscall.Syscall6(syscall.SYS_LISTXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(namebuf)),
                uintptr(bufsize), uintptr(options), 0, 0)
        resSize = int(r0)
        if e1 != noError {
                e = e1
        }
        return
}

// int setxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options);
func setxattr(path string, attrname string, value *byte, size int, position uint32, options int) (e error) {
        _, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(attrname))),
                uintptr(unsafe.Pointer(value)),
                uintptr(size), uintptr(position), uintptr(options))
        if e1 != noError {
                e = e1
        }
        return
}

// int removexattr(user_addr_t path, user_addr_t attrname, int options);
func removexattr(path string, attrname string, options int) (e error) {
        _, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(attrname))),
                uintptr(options))
        if e1 != noError {
                e = e1
        }
        return
}
