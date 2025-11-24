package service

type Service struct {
	storage Storage
}

func New(st Storage) *Service {
	return &Service{storage: st}
}

func (s *Service) CreateItem(username string, name string, qty int) error {
	return s.storage.CreateItem(name, qty, username)
}
