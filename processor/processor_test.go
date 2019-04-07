package processor

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestProcessor_Process(t *testing.T) {
	sourceDir := "./test_source_dir"
	targetDir := "./test_source_dir"
	e := os.MkdirAll(sourceDir, os.ModePerm)
	if e != nil {
		panic(e)
	}
	if e := ioutil.WriteFile(fmt.Sprintf("%v/%v", sourceDir, "test"), []byte("hello"), os.ModePerm); e != nil {
		panic(e)
	}

	p, e := InitProcessor(sourceDir, targetDir, true, 10)
	if e != nil {
		panic(e)
	}

	p.Process()
}
