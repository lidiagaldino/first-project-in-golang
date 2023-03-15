package pdfGenerator

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkhtml struct {
	rootPath string
}

func NewWK(rootPath string) PDFGeneratorInterface {
	return &wkhtml{rootPath: rootPath}
}

func (w *wkhtml) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "i", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Create()

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "a", err
	}

	return fileName, nil
}
