// Linux xattrs have a manditory prefix of "user.". This is prepended
// transparently for Get/Set/Remove and hidden in List
package xattr

// XAttrError records an error and the operation, file path and attribute that caused it.
type XAttrError struct {
	Op   string
	Path string
	Name string
	Err  error
}

func (e *XAttrError) Error() string {
	return e.Op + " " + e.Path + " " + e.Name + ": " + e.Err.Error()
}

// Convert an array of NUL terminated UTF-8 strings
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

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, error) {
        name = prefix + name

        // find size
        size, err := getxattr(path, name, nil, 0)
        if err != nil {
                return nil, &XAttrError{"getxattr", path, name, err}
        }
        if size == 0 {
                return []byte{}, nil
        }

        // read into buffer of that size
        buf := make([]byte, size)
        size, err = getxattr(path, name, &buf[0], size)
        if err != nil {
                return nil, &XAttrError{"getxattr", path, name, err}
        }
        return buf[:size], nil
}

// Retrieves a list of names of extended attributes associated with the
// given path in the file system.
func Listxattr(path string) ([]string, error) {
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

// Associates name and data together as an attribute of path.
func Setxattr(path, name string, data []byte) error {
        name = prefix + name
        l := len(data)
        var p *byte
        if l != 0 {
                p = &data[0]
        }
        if err := setxattr(path, name, p, l); err != nil {
                return &XAttrError{"setxattr", path, name, err}
        }
        return nil
}

// Remove the attribute.
func Removexattr(path, name string) error {
        name = prefix + name
        if err := removexattr(path, name); err != nil {
                return &XAttrError{"removexattr", path, name, err}
        }
        return nil
}
