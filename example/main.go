package main

import (
	"fmt"
	"github.com/davecheney/xattr"
	"flag"
)

func main() {
	// top level flags
	print := flag.Bool("p", false, "print")
	write := flag.Bool("w", false, "write")
	delete := flag.Bool("d", false, "delete")
	flag.Parse()

	switch {
	case *print:
		name := flag.Arg(0)
		for _, file := range flag.Args()[1:] {
			value, _ := xattr.Getxattr(file, name)
			fmt.Println(value)
		}
	case *write:
		name, value := flag.Arg(0), flag.Arg(1)
		for _, file := range flag.Args()[2:] {
			xattr.Setxattr(file, name, []byte(value))
		}
	case *delete:
		name := flag.Arg(0)
		for _, file := range flag.Args()[1:] {
			xattr.Removexattr(file, name)
		}
	default:
		for _, file := range flag.Args() {
			values, _ := xattr.Listxattr(file)
			for _, value := range values {
				fmt.Println(value)
			}
		}
	}
}
