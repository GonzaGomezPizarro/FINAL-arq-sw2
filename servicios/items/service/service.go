package service

import (
	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/errors"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/model"
)

type itemService struct{}

type itemServiceInterface interface {
	GetItemById(id int) (dto.Item, e.ApiError)
	GetItems() (dto.Items, e.ApiError)
	NewItem(ItemDto dto.Item) (dto.Item, e.ApiError)
	NewItems(ItemsDto dto.Items) (dto.Items, e.ApiError)
	DeleteItem(id int) e.ApiError
}

var jwtKey = []byte("secret_key")

var (
	ItemService itemServiceInterface
)

func init() {
	ItemService = &itemService{}
}

func (s *itemService) GetItemById(id int) (dto.Item, e.ApiError) {

	var item model.Item = cliente.GetItemById(id)
	var itemDto dto.Item

	if item.Id == 0 {
		return itemDto, e.NewBadRequestApiError("item not found")
	}

	itemDto.Id = item.Id
	itemDto.Title = item.Title
	itemDto.Description = item.Description
	itemDto.UserId = item.UserId
	itemDto.Address = item.Address
	itemDto.Country = item.Country
	itemDto.State = item.State
	itemDto.Photos = item.Photos
	itemDto.City = item.City

	return itemDto, nil
}

func (s *itemService) GetItems() (dto.Items, e.ApiError) {

	var items model.Items = cliente.GetItems()
	var itemsDto dto.Items

	for _, item := range items {
		var itemDto dto.Item

		itemDto.Id = item.Id
		itemDto.Title = item.Title
		itemDto.Description = item.Description
		itemDto.UserId = item.UserId
		itemDto.Address = item.Address
		itemDto.Country = item.Country
		itemDto.State = item.State
		itemDto.Photos = item.Photos
		itemDto.City = item.City

		itemsDto = append(itemsDto, itemDto)
	}

	return itemsDto, nil
}

func (s *itemService) NewItem(itemDto dto.Item) (dto.Item, e.ApiError) {
	var item model.Item

	item.Title = itemDto.Title
	item.Description = itemDto.Description
	item.UserId = itemDto.UserId
	item.Address = itemDto.Address
	item.Country = itemDto.Country
	item.State = itemDto.State
	item.Photos = itemDto.Photos
	item.City = itemDto.City

	var err e.ApiError
	item, err = cliente.Newitem(item)

	itemDto.Id = item.Id

	return itemDto, err

}

func (s *itemService) NewItems(itemsDto dto.Items) (dto.Items, e.ApiError) {
	var items model.Items

	for _, itemDto := range itemsDto {
		var item model.Item
		item.Address = itemDto.Address
		item.City = itemDto.City
		item.Country = itemDto.Country
		item.Description = itemDto.Description
		item.Photos = itemDto.Photos
		item.State = itemDto.State
		item.Title = itemDto.Title
		item.UserId = itemDto.UserId

		items = append(items, item)
	}

	var err e.ApiError
	items, err = cliente.NewItems(items)

	var itemssDto dto.Items
	for _, item := range items {
		if item.Id == 0 {
			return itemsDto, e.NewBadRequestApiError("Something went wrong when creating, some items might have been created, but some not")
		}
		var itemDto dto.Item
		itemDto.Address = item.Address
		itemDto.City = item.City
		itemDto.Country = item.Country
		itemDto.State = item.State
		itemDto.Id = item.Id
		itemDto.Description = item.Description
		itemDto.Title = item.Title
		itemDto.UserId = item.UserId
		itemDto.Photos = item.Photos

		itemssDto = append(itemssDto, itemDto)
	}
	return itemssDto, err

}

func (s *itemService) Deleteitem(id int) e.ApiError {
	err := cliente.Deleteitem(id)

	return err
}
