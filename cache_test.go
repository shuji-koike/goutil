package goutil

import (
	"log"
	"testing"
)

var cache = CacheSingleflight{
	UseMem:  true,
	Dir:     nil,
	Postfix: ".gob.gz",
}

func job(key string) (int, error) {
	return len(key), nil
}

func TestCacheSingleflight(t *testing.T) {
	data, err := cache.Get("test", job)
	log.Printf("%s, %s", data, err)
}
