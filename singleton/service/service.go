package service

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
