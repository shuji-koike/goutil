package goutil

import (
	"log"
)

// LogError prints error
// defer goutil.LogError(&err)
func LogError(err *error) error {
	if err != nil && *err != nil {
		log.Output(2, (*err).Error())
		return *err
	}
	return nil
}
