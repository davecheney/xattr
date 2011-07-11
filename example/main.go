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
	long := flag.Bool("l", false, "long")
	flag.Parse()

	switch {
	case *print:
		name := flag.Arg(0)
		for _, file := range flag.Args()[1:] {
			value, _ := xattr.Getxattr(file, name)
			if *long {
				fmt.Printf("%s: %s\n", name, value)
			} else {
				fmt.Println(string(value))
			}
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
			names, _ := xattr.Listxattr(file)
			for _, name := range names {
				if *long {
					value, _ := xattr.Getxattr(file, name)
					fmt.Printf("%s: %s\n", name, value)
				} else {
					fmt.Println(name)
				}
			}
		}
	}
}
