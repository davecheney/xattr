package xattr

import (
	"os"
)

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
	data, err := Getxattr0(path, name)
	if err != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(err)}
	}
	return data, nil
}

// Retrieves a list of names of extended attributes associated with the 
// given path in the file system.
func Listxattr(path string) ([]string, os.Error) {
	buf, err := Listxattr0(path)
	if err != 0 {
		return nil, &XAttrError{"listxattr", path, "", os.Errno(err)}
	}
	return nullTermToStrings(buf), nil
}

// Associates name and data together as an attribute of path. 
func Setxattr(path, name string, data []byte) os.Error {
	err := Setxattr0(path, name, data)
	if err != 0 {
		return &XAttrError{"setxattr", path, name, os.Errno(err)}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) os.Error {
	err := Removexattr0(path, name)
	if err != 0 {
		return &XAttrError{"removexattr", path, name, os.Errno(err)}
	}
	return nil
}
