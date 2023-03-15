package htmlParser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type HTMLStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &HTMLStruct{rootPath: rootPath}
}

func (a *HTMLStruct) Create(templateName string, data interface{}) (string, error) {

	templateGenerator, err := template.ParseFiles(templateName)

	if err != nil {
		return "", err
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	err = templateGenerator.Execute(fileWriter, data)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
