package goutil

import (
	"compress/gzip"
	"encoding/json"
	"os"
)

// ReadJSON ...
func ReadJSON(path string, data interface{}) error {
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
	err = json.NewDecoder(gz).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// WriteJSON ...
func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	gz := gzip.NewWriter(file)
	defer gz.Close()
	err = json.NewEncoder(gz).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
