package service

import (
	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/errors"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/model"
)

type itemService struct{}

type itemServiceInterface interface {
	GetItemById(id string) (dto.Item, e.ApiError)
	GetItems() (dto.Items, e.ApiError)
	NewItem(ItemDto dto.Item) (dto.Item, e.ApiError)
	NewItems(ItemsDto dto.Items) (dto.Items, e.ApiError)
	DeleteItem(id string) e.ApiError
}

var jwtKey = []byte("secret_key")

var (
	ItemService itemServiceInterface
)

func init() {
	ItemService = &itemService{}
}

func (s *itemService) GetItemById(id string) (dto.Item, e.ApiError) {

	item, error := cliente.GetItemById(id)
	var itemDto dto.Item

	if error != nil {
		return itemDto, e.NewNotFoundApiError("item not found")
	}

	itemDto.Id = item.Id.Hex()
	itemDto.Title = item.Title
	itemDto.Description = item.Description
	itemDto.UserId = item.UserId
	itemDto.Address = item.Address
	itemDto.Country = item.Country
	itemDto.State = item.State
	itemDto.Photos = item.Photos
	itemDto.City = item.City
	itemDto.Price = item.Price
	itemDto.Bedrooms = item.Bedrooms
	itemDto.Bathrooms = item.Bathrooms
	itemDto.Mts2 = item.Mts2

	return itemDto, nil
}

func (s *itemService) GetItems() (dto.Items, e.ApiError) {

	items, error := cliente.GetItems()
	var itemsDto dto.Items

	if error != nil {
		return itemsDto, e.NewNotFoundApiError("item not found")
	}

	for _, item := range items {
		var itemDto dto.Item

		itemDto.Id = item.Id.Hex()
		itemDto.Title = item.Title
		itemDto.Description = item.Description
		itemDto.UserId = item.UserId
		itemDto.Address = item.Address
		itemDto.Country = item.Country
		itemDto.State = item.State
		itemDto.Photos = item.Photos
		itemDto.City = item.City
		itemDto.Price = item.Price
		itemDto.Bedrooms = item.Bedrooms
		itemDto.Bathrooms = item.Bathrooms
		itemDto.Mts2 = item.Mts2

		itemsDto = append(itemsDto, itemDto)
	}

	return itemsDto, nil
}

func (s *itemService) NewItem(itemDto dto.Item) (dto.Item, e.ApiError) {
	var item model.Item
	var newItem dto.Item

	item.Title = itemDto.Title
	item.Description = itemDto.Description
	item.UserId = itemDto.UserId
	item.Address = itemDto.Address
	item.Country = itemDto.Country
	item.State = itemDto.State
	item.Photos = itemDto.Photos
	item.City = itemDto.City
	item.Price = itemDto.Price
	item.Bedrooms = itemDto.Bedrooms
	item.Bathrooms = itemDto.Bathrooms
	item.Mts2 = itemDto.Mts2

	var err e.ApiError
	item, err = cliente.NewItem(item)

	newItem.Id = item.Id.Hex()
	newItem.Title = item.Title
	newItem.Description = item.Description
	newItem.UserId = item.UserId
	newItem.Address = item.Address
	newItem.Country = item.Country
	newItem.State = item.State
	newItem.Photos = item.Photos
	newItem.City = item.City
	newItem.Price = item.Price
	newItem.Bedrooms = item.Bedrooms
	newItem.Bathrooms = item.Bathrooms
	newItem.Mts2 = item.Mts2

	return newItem, err

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
		item.Price = itemDto.Price
		item.Bedrooms = itemDto.Bedrooms
		item.Bathrooms = itemDto.Bathrooms
		item.Mts2 = itemDto.Mts2

		items = append(items, item)
	}

	var err e.ApiError
	items, err = cliente.NewItems(items)

	if err != nil {
		return itemsDto, err
	}

	var itemssDto dto.Items
	for _, item := range items {
		var itemDto dto.Item
		itemDto.Address = item.Address
		itemDto.City = item.City
		itemDto.Country = item.Country
		itemDto.State = item.State
		itemDto.Id = item.Id.Hex()
		itemDto.Description = item.Description
		itemDto.Title = item.Title
		itemDto.UserId = item.UserId
		itemDto.Photos = item.Photos
		itemDto.Price = item.Price
		itemDto.Bedrooms = item.Bedrooms
		itemDto.Bathrooms = item.Bathrooms
		itemDto.Mts2 = item.Mts2

		itemssDto = append(itemssDto, itemDto)
	}
	return itemssDto, err

}

func (s *itemService) DeleteItem(id string) e.ApiError {
	err := cliente.DeleteItem(id)
	return err
}
