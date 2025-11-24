package service

import "github.com/AugustSerenity/go-contest-L3/l3.7_WarehouseControl/internal/model"

type Storage interface {
	CreateItem(string, int, string) error
	ListItems() ([]model.Item, error)
}
