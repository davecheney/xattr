package xattr

import (
	"strings"
	"syscall"
)

const (
	userPrefix = "user."
)

// Linux xattrs have a manditory prefix of "user.". This is prepended
// transparently for Get/Set/Remove and hidden in List

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, error) {
	name = userPrefix + name
	// find size.
	size, err := syscall.Getxattr(path, name, nil)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, err := syscall.Getxattr(path, name, buf)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
	return buf[:read], nil
}

// Retrieves a list of names of extended attributes associated with the
// given path in the file system.
func Listxattr(path string) ([]string, error) {
	// find size.
	size, err := syscall.Listxattr(path, nil)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, err := syscall.Listxattr(path, buf)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	return stripUserPrefix(nullTermToStrings(buf[:read])), nil
}

// Associates name and data together as an attribute of path.
func Setxattr(path, name string, data []byte) error {
	name = userPrefix + name
	if err := syscall.Setxattr(path, name, data, 0); err != nil {
		return &XAttrError{"setxattr", path, name, err}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) error {
	name = userPrefix + name
	if err := syscall.Removexattr(path, name); err != nil {
		return &XAttrError{"removexattr", path, name, err}
	}
	return nil
}

// Strip off "user." prefixes from attribute names.
func stripUserPrefix(s []string) []string {
	for i, a := range s {
		if strings.HasPrefix(a, userPrefix) {
			s[i] = a[5:]
		}
	}
	return s
}
