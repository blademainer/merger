package main

import (
	"flag"
	"fmt"
	"github.com/blademainer/merger/processor"
)

func main() {
	sourceDir := flag.String("s", "", "source dir")
	targetDir := flag.String("t", "", "target dir")
	removeSource := flag.Bool("r", true, "remove source dir")
	concurrent := flag.Uint("c", 100, "remove source dir")
	debug := flag.Bool("d", true, "remove source dir")
	flag.Parse()
	fmt.Println("sourceDir: ", *sourceDir)
	fmt.Println("targetDir: ", *targetDir)
	fmt.Println("deleteSource: ", *removeSource)
	p, e := processor.InitProcessor(*sourceDir, *targetDir, *removeSource, int(*concurrent), *debug)
	if e != nil {
		panic(e)
	}
	p.Process()
}
