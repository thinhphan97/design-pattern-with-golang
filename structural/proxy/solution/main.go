package main

import (
	"fmt"
	"time"
)

type DataStorage interface {
	GetValue() int
}

type RealDataStorage struct{}

func (RealDataStorage) GetValue() int {
	time.Sleep(time.Second * 5)
	return 100
}

type ProxyDataStorage struct {
	CachedValue *int
	realStorage DataStorage
}

func (s ProxyDataStorage) GetValue() int {
	if val := s.CachedValue; val != nil {
		return *val
	}
	val := s.realStorage.GetValue()
	s.CachedValue = &val
	return val
}

func NewProxyDataStorage(realStorage DataStorage) ProxyDataStorage {
	return ProxyDataStorage{realStorage: realStorage}
}

type ValueService struct {
	storage DataStorage
}

func (s ValueService) FetchValue() int {
	return s.storage.GetValue()
}

func main() {
	value := ValueService{
		storage: NewProxyDataStorage(RealDataStorage{}),
	}.FetchValue()
	// It's too low at the first time
	fmt.Println(value)
	// Now it return instantly because of caching (proxy layer)
	fmt.Println(value)

}
