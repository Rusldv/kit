package fileutil

import (
	"io/ioutil"
	"os"
)

// ReadFileString читает данные из файла
func ReadFileString(name string) (string, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFileString записывает данные в файл
func WriteFileString(name, str string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(str)
	return nil
}
