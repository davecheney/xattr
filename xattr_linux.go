package xattr

import (
	"strings"
	"os"
)

const (
	userPrefix = "user."
)

// Linux xattrs have a manditory prefix of "user.". This is prepended 
// transparently for Get/Set/Remove and hidden in List

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, os.Error) {
	name = userPrefix + name
	// find size.
	size, e := getxattr(path, name, nil, 0)
	if e != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(e)}
	}
	buf := make([]byte, size)
	// Read into buffer of that size.
	read, e := getxattr(path, name, &buf[0], size)
	if e != 0 {
		return nil, &XAttrError{"getxattr", path, name, os.Errno(e)}
	}
	return buf[:read], nil
}

// Retrieves a list of names of extended attributes associated with the 
// given path in the file system.
func Listxattr(path string) ([]string, os.Error) {
	buf, e := listxattr(path)
	if e!= 0 {
		return nil, &XAttrError{"listxattr", path, "", os.Errno(e)}
	}
	return stripUserPrefix(nullTermToStrings(buf)), nil
}

// Associates name and data together as an attribute of path. 
func Setxattr(path, name string, data []byte) os.Error {
	name = userPrefix + name
	e:= setxattr(path, name, &data[0], len(data))
	if e!= 0 {
		return &XAttrError{"setxattr", path, name, os.Errno(e)}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) os.Error {
	name = userPrefix + name
	e:= removexattr(path, name)
	if e!= 0 {
		return &XAttrError{"removexattr", path, name, os.Errno(e)}
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
