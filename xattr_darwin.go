package xattr

import (
	"os"
)

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, os.Error) {
	// find size.
	size, e := getxattr(path, name, nil, 0, 0, 0)
	if e != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(e)}
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, e := getxattr(path, name, &buf[0], size, 0, 0)
	if e != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(e)}
	}
	return buf[:read], nil
}

// Retrieves a list of names of extended attributes associated with the 
// given path in the file system.
func Listxattr(path string) ([]string, os.Error) {
	// find size.
	size, e := listxattr(path, nil, 0, 0)
	if e != 0 {
		return nil, &XAttrError{"listxattr", path, "", os.Errno(e)}
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, e := listxattr(path, &buf[0], size, 0)
	if e != 0 {
		return nil, &XAttrError{"listxattr", path, "", os.Errno(e)}
	}
	return nullTermToStrings(buf[:read]), nil
}

// Associates name and data together as an attribute of path. 
func Setxattr(path, name string, data []byte) os.Error {
	e := setxattr(path, name, &data[0], len(data), 0, 0)
	if e != 0 {
		return &XAttrError{"setxattr", path, name, os.Errno(e)}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) os.Error {
	e := removexattr(path, name, 0)
	if e != 0 {
		return &XAttrError{"removexattr", path, name, os.Errno(e)}
	}
	return nil
}
