package service

type AsyncService struct {}

func NewAsyncService() *AsyncService {
	return &AsyncService{}
}

func (a *AsyncService) RunAsync(task func()) {
	go task()
}