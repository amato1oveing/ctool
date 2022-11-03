package ctool

import (
	"golang.org/x/sync/singleflight"
)

// SingleFlight 减少重复
type SingleFlight struct {
	requestGroup singleflight.Group
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		requestGroup: singleflight.Group{},
	}
}

func (s *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	do, err, _ := s.requestGroup.Do(key, fn)
	return do, err
}
