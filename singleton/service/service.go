package service

import "log"

// Singleton Service

type IdServiceSingleton struct {
	idService *IdService
}

func (s *IdServiceSingleton) GetService() *IdService {
	if s.idService == nil {
		log.Print("no id service available, instantiation")
		s.idService = NewIdService()
	}
	return s.idService
}

// ID Service
type IdService struct {
	counter int
}

func NewIdService() *IdService {
	return &IdService{0}
}

func (i *IdService) Next() int {
	i.counter++
	return i.counter
}
