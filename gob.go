package goutil

import (
	"compress/gzip"
	"encoding/gob"
	"io"
	"os"
	"strings"
)

// ReadGob ...
func ReadGob(path string, data interface{}) error {
	var reader io.ReadCloser
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if strings.HasSuffix(path, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		defer gz.Close()
		reader = gz
	} else {
		reader = file
	}
	err = gob.NewDecoder(reader).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// WriteGob ...
func WriteGob(path string, data interface{}) error {
	var writer io.WriteCloser
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if strings.HasSuffix(path, ".gz") {
		gz, err := gzip.NewWriterLevel(file, CompressionLevel)
		if err != nil {
			return err
		}
		defer gz.Close()
		writer = gz
	} else {
		writer = file
	}
	err = gob.NewEncoder(writer).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
