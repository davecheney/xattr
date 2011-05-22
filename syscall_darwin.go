package xattr

func Getxattr0(path, name string) (buf []byte, errno int) {
	// find size.
	size, err := getxattr(path, name, nil, 0, 0, 0)
	if err != 0 {
		return nil, err
	}
	buf = make([]byte, size)
	// Read into buffer of that size.
	read, err := getxattr(path, name, &buf[0], size, 0, 0)
	if err != 0 {
		return nil, err
	}
	return buf[:read], 0
}

func Listxattr0(path string) (buf []byte, errno int) {
	// find size.
	size, err := listxattr(path, nil, 0, 0)
	if err != 0 {
		return nil, err
	}
	buf = make([]byte, size)
	// Read into buffer of that size.
	read, err := listxattr(path, &buf[0], size, 0)
	if err != 0 {
		return nil, err
	}
	return buf[:read], 0
}

func Setxattr0(path string, name string, value []byte) (errno int) {
	return setxattr(path, name, &value[0], len(value), 0, 0)
}

func Removexattr0(path string, name string) (errno int) {
	return removexattr(path, name, 0)
}
