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
	time.Sleep(time.Second * 2)
	return 100
}

type ValueService struct {
	Storage DataStorage
}

func (s ValueService) FetchValue() int {
	return s.Storage.GetValue()
}
func main() {
	value := ValueService{Storage: RealDataStorage{}}.FetchValue()
	fmt.Println(value)
}
