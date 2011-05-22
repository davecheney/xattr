package xattr

import (
	"strings"
	"syscall"
	"os"
)

const (
	userPrefix = "user."
)

// Linux xattrs have a manditory prefix of "user.". This is prepended 
// transparently for Get/Set/Remove and hidden in List

// XAttrError records an error and the operation, file path and attribute that caused it.
type XAttrError struct {
	Op    string
	Path  string
	Name  string
	Error os.Error
}

func (e *XAttrError) String() string {
	return e.Op + " " + e.Path + " " + e.Name + ": " + e.Error.String()
}

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, os.Error) {
	name = userPrefix + name
	data, err := syscall.Getxattr(path, name)
	if err != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(err)}
	}
	return data, nil
}

// Retrieves a list of names of extended attributes associated with the 
// given path in the file system.
func Listxattr(path string) ([]string, os.Error) {
	buf, err := syscall.Listxattr(path)
	if err != 0 {
		return nil, &XAttrError{"listxattr", path, "", os.Errno(err)}
	}
	return stripUserPrefix(nullTermToStrings(buf)), nil
}

// Associates name and data together as an attribute of path. 
func Setxattr(path, name string, data []byte) os.Error {
	name = userPrefix + name
	err := syscall.Setxattr(path, name, data)
	if err != 0 {
		return &XAttrError{"setxattr", path, name, os.Errno(err)}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) os.Error {
	name = userPrefix + name
	err := syscall.Removexattr(path, name)
	if err != 0 {
		return &XAttrError{"removexattr", path, name, os.Errno(err)}
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
