package main

import (
	"log"
	"path"

	"github.com/playmean/guest/ui"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	converter := typescriptify.New()
	converter.CreateInterface = true
	converter.BackupDir = ""
	converter.Prefix = "Api"
	converter.Add(ui.GetVersionResponse{})
	converter.Add(ui.GetWorkspaceResponse{})
	converter.Add(ui.ServerError{})

	err := converter.ConvertToFile(path.Join("ui/frontend/src/types/api.ts"))
	if err != nil {
		panic(err)
	}

	log.Println("types generated successfully")
}
