package goutil

import (
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"golang.org/x/sync/singleflight"
)

// CacheSingleflight (work in progress)
type CacheSingleflight struct {
	UseMem  bool
	Dir     *string
	Postfix string
	group   singleflight.Group
	cache   sync.Map
}

// Get (work in progress)
func (s *CacheSingleflight) Get(key string, fn interface{}) (interface{}, error) {
	data, err, _ := s.group.Do(key, func() (interface{}, error) {
		fnV := reflect.ValueOf(fn)
		var data interface{}
		dataTP := reflect.New(fnV.Type().Out(0))
		var err error
		if m, ok := s.cache.Load(key); ok {
			logger.Printf("GobCache.Get: from cache")
			return m, nil
		}
		if s.Dir != nil {
			logger.Printf("GobCache.Get: GobRead")
			err = GobRead(filepath.Join(*s.Dir, key)+s.Postfix, dataTP.Interface())
			if err != nil && !os.IsNotExist(err) {
				return dataTP.Interface(), err
			}
		}
		if s.Dir == nil || os.IsNotExist(err) {
			logger.Printf("GobCache.Get: fnV.Call")
			ret := fnV.Call([]reflect.Value{reflect.ValueOf(key)})
			err = ret[0].Interface().(error)
			if s.Dir != nil && err == nil {
				logger.Printf("GobCache.Get: GobSave")
				err = GobSave(filepath.Join(*s.Dir, key)+s.Postfix, data)
			}
		}
		if s.UseMem && err == nil {
			logger.Printf("GobCache.Get: cache.Store")
			s.cache.Store(key, data)
		}
		return data, err
	})
	return data, err
}
