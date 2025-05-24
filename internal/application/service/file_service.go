package application

import (
	"bytes"
	"os"
	"path/filepath"
)

type FileService struct {
	FilePath   string
	buffer     *bytes.Buffer
	OutputFile *os.File
}

func NewFileService() *FileService {
	return &FileService{
		buffer: &bytes.Buffer{},
	}
}

func (f *FileService) SetFile(fileName, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	f.FilePath = filepath.Join(path, fileName)
	filePath, err := os.Create(f.FilePath)

	if err != nil {
		return err
	}

	f.OutputFile = filePath
	return nil
}

func (f *FileService) Write(chunk []byte) error {

	if f.OutputFile == nil {
		return nil
	}

	_, err := f.OutputFile.Write(chunk)
	return err
}

func (f *FileService) Close() error {
	return f.OutputFile.Close()
}
