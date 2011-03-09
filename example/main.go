package main

import (
	"fmt"
	"github.com/davecheney/xattr"
	"flag"
	"log"	
)

func main() {
	argv := flag.Args()
	err := xattr.Setxattr(argv[0], argv[1], []byte(argv[2]))
	if err != nil {
		log.Fatal(err)
	}
	list, err := xattr.Listxattr(argv[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %#v\n", argv[0], list)
	value, err := xattr.Getxattr(argv[0], argv[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %s\n", argv[1], string(value))
}
