package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ZachEddy/helmzzy/command"
)

func main() {
	flag.Usage = Usage
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}
	if val, ok := command.CommandMap[os.Args[1]]; ok {
		if err := val.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("no such command %q\n", os.Args[1])
	}
}
