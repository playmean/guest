package debug

import (
	"log"

	"github.com/go-git/go-billy/v5"
)

func PrintFs(virtualFs billy.Filesystem) {
	e, err := virtualFs.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range e {
		log.Println(f.Name())
	}
}
