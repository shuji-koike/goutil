package goutil

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

// GobRead ...
func GobRead(path string, data interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gz.Close()
	err = gob.NewDecoder(gz).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// GobSave ...
func GobSave(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	gz := gzip.NewWriter(file)
	defer gz.Close()
	err = gob.NewEncoder(gz).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
