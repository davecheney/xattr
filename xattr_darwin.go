package xattr

// Retrieve extended attribute data associated with path.
func Getxattr(path, name string) ([]byte, error) {
        // find size
	size, err := getxattr(path, name, nil, 0, 0, 0)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
        if size == 0 {
                return []byte{}, nil
        }

        // read into buffer of that size
	buf := make([]byte, size)
        size, err = getxattr(path, name, &buf[0], size, 0, 0)
	if err != nil {
		return nil, &XAttrError{"getxattr", path, name, err}
	}
        return buf[:size], nil
}

// Retrieves a list of names of extended attributes associated with the
// given path in the file system.
func Listxattr(path string) ([]string, error) {
        // find size
	size, err := listxattr(path, nil, 0, 0)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
        if size == 0 {
                return []string{}, nil
        }

        // read into buffer of that size
	buf := make([]byte, size)
        size, err = listxattr(path, &buf[0], size, 0)
	if err != nil {
		return nil, &XAttrError{"listxattr", path, "", err}
	}
        return nullTermToStrings(buf[:size]), nil
}

// Associates name and data together as an attribute of path.
func Setxattr(path, name string, data []byte) error {
        l := len(data)
        var p *byte
        if l != 0 {
                p = &data[0]
        }
        if err := setxattr(path, name, p, l, 0, 0); err != nil {
		return &XAttrError{"setxattr", path, name, err}
	}
	return nil
}

// Remove the attribute.
func Removexattr(path, name string) error {
	if err := removexattr(path, name, 0); err != nil {
		return &XAttrError{"removexattr", path, name, err}
	}
	return nil
}
