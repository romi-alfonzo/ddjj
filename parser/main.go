package main

import (
	"fmt"
	"github.com/InstIDEA/ddjj/parser/extract"
	"github.com/InstIDEA/ddjj/parser/server"
	"os"
)

func handleSingleFile(filePath string) extract.ParserData {
	dat, err := os.Open(filePath)

	if err != nil {
		return extract.CreateError(fmt.Sprint("File ", filePath, " not found. ", err))
	}

	return extract.ParsePDF(dat)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: parser file.pdf")
		os.Exit(1)
		return
	}
	if os.Args[1] == "serve" {
		server.InitServer()
		return
	}
	parsed := handleSingleFile(os.Args[1])
	parsed.Print()
}
