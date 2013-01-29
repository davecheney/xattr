package xattr

import (
        "syscall"
        "unsafe"
)

// ssize_t getxattr(const char *path, const char *name, void *value, size_t size);
func getxattr(path string, name string, value *byte, size int) (int, error) {
        r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
                uintptr(unsafe.Pointer(value)),
                uintptr(size), 0, 0)
        return int(r0), e1
}

// ssize_t listxattr(const char *path, char *list, size_t size);
func listxattr(path string, list *byte, size int) (int, error) {
        r0, _, e1 := syscall.Syscall(syscall.SYS_LISTXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(list)),
                uintptr(size))
        return int(r0), e1
}

// int setxattr(const char *path, const char *name, const void *value, size_t size, int flags);
func setxattr(path string, name string, value *byte, size int) (err error) {
        _, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
                uintptr(unsafe.Pointer(value)),
                uintptr(size), 0, 0)
        return e1
}

// int removexattr(const char *path, const char *name);
func removexattr(path string, name string) (err error) {
        _, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR,
                uintptr(unsafe.Pointer(syscall.StringBytePtr(path))),
                uintptr(unsafe.Pointer(syscall.StringBytePtr(name))),
                0)
        return e1
}
