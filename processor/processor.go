package processor

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"sync"
)

type Processor struct {
	sourceDir    os.File
	targetDir    os.File
	deleteSource bool
	concurrent   int `form:"concurrent"`
	fileChan     chan os.File
	done         chan bool
	sync.WaitGroup
}

func InitProcessor(sourcePath string, targetPath string, deleteSource bool, concurrent int) (p *Processor, e error) {
	validate := validator.New()
	e = validate.Var(sourcePath, "required")
	if e != nil {
		return
	}
	e = validate.Var(targetPath, "required")
	if e != nil {
		return
	}

	var sourceDir, targetDir *os.File
	if sourceDir, e = os.Open(sourcePath); e != nil {
		return
	} else if targetDir, e = os.Open(targetPath); e != nil {
		return
	}

	p = &Processor{}
	p.sourceDir = *sourceDir
	p.targetDir = *targetDir
	p.deleteSource = deleteSource
	p.concurrent = concurrent
	p.fileChan = make(chan os.File, concurrent)
	p.done = make(chan bool, 1)
	return
}

func (p *Processor) Process() {
	p.Add(2)
	go func() {
		defer func() {
			p.done <- true
			p.Done()
		}()
		p.readSourceDir()
	}()
	go func() {
		defer p.Done()
		p.processTargetDir()
	}()

	p.Wait()
}

func (p *Processor) readSourceDir() {
	p.listChildren(p.sourceDir)
}

func (p *Processor) listChildren(file os.File) {
	if fi, e := file.Stat(); e != nil {
		fmt.Printf("Failed to get file info! file: %v error: %v \n", file, e.Error())
		return
	} else {
		name := fi.Name()
		fmt.Printf("Found file: %v\n", name)
		if fi.IsDir() {

		} else {

		}
	}

}

func (p *Processor) processTargetDir() {
	for {
		select {
		case file := <-p.fileChan:
			fmt.Printf("Processing file: %v\n", file)
		case <-p.done:
			fmt.Println("Exit processor by done channel.")
			return
		}
	}
}
