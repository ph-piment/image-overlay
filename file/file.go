package file

import (
	"path/filepath"
)

const (
	SourceDir string = "source"
	OutputDir string = "output"
)

func GetSourceFilePath(fileName string) (path string, err error) {
	path, err = filepath.Abs("./" + SourceDir + "/" + fileName)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetOutputFilePath(fileName string) (path string, err error) {
	path, err = filepath.Abs("./" + OutputDir + "/" + fileName)
	if err != nil {
		return "", err
	}
	return path, nil
}
