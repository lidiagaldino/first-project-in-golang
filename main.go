package main

import (
	"fmt"
	"pdfGenerator/htmlParser"
	"pdfGenerator/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {
	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWK("tmp")

	dataHTML := Data{
		Name: "Lídia",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("html gerador", htmlGenerated)

	filePdfName, err := wk.Create(htmlGenerated)

	if err != nil {
		fmt.Println(filePdfName)
		fmt.Println(err)
		return
	}

	fmt.Println("pdf gerador", filePdfName)

}
