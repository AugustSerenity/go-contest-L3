package handler

type Service interface {
	CreateItem(string, string, int) error
}
