package main

import (
	"flag"
	"fmt"
	"github.com/blademainer/merger/processor"
)



func main() {
	p := &processor.Processor{}
	sourceDir := flag.String("s", "", "source dir")
	targetDir := flag.String("t", "", "target dir")
	removeSource := flag.Bool("d", true, "remove source dir")
	flag.Parse()
	fmt.Println("sourceDir: ", *sourceDir)
	fmt.Println("targetDir: ", *targetDir)
	fmt.Println("deleteSource: ", *removeSource)

}
