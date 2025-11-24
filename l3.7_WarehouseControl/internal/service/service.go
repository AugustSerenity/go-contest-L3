package service

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{storage: st}
}
