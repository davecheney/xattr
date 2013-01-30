// Package xattr provides a simple interface to user extended attributes on Linux and OSX.
// Support for xattrs is filesystem dependant, so not a given even if you are running one of those operating systems.
//
// On Linux you have to edit /etc/fstab to include "user_xattr". Also, Linux extended attributes have a manditory
// prefix of "user.". This is prepended transparently for Get/Set/Remove and hidden in List.
package xattr

// XAttrError records an error and the operation, file path and attribute that caused it.
type XAttrError struct {
	Op   string
	Path string
	Attr string
	Err  error
}

func (e *XAttrError) Error() string {
	return e.Op + " " + e.Path + " " + e.Attr + ": " + e.Err.Error()
}

// Converts an array of NUL terminated UTF-8 strings
// to a []string.
func nullTermToStrings(buf []byte) (result []string) {
	offset := 0
	for index, b := range buf {
		if b == 0 {
			result = append(result, string(buf[offset:index]))
			offset = index + 1
		}
	}
	return
}

// Retrieves extended attribute data associated with path.
func Get(path, attr string) ([]byte, error) {
	attr = prefix + attr

	// find size
	size, err := getxattr(path, attr, nil, 0)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, attr, err}
	}
	if size == 0 {
		return []byte{}, nil
	}

	// read into buffer of that size
	buf := make([]byte, size)
	size, err = getxattr(path, attr, &buf[0], size)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, attr, err}
	}
	return buf[:size], nil
}

// Retrieves a list of names of extended attributes associated with path.
func List(path string) ([]string, error) {
	// find size
	size, err := listxattr(path, nil, 0)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	if size == 0 {
		return []string{}, nil
	}

	// read into buffer of that size
	buf := make([]byte, size)
	size, err = listxattr(path, &buf[0], size)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
	return stripPrefix(nullTermToStrings(buf[:size])), nil
}

// Associates data as an extended attribute of path.
func Set(path, attr string, data []byte) error {
	attr = prefix + attr
	l := len(data)
	var p *byte
	if l != 0 {
		p = &data[0]
	}
	if err := setxattr(path, attr, p, l); err != nil {
		return &XAttrError{"setxattr", path, attr, err}
	}
	return nil
}

// Removes the extended attribute.
func Remove(path, attr string) error {
	attr = prefix + attr
	if err := removexattr(path, attr); err != nil {
		return &XAttrError{"removexattr", path, attr, err}
	}
	return nil
}
