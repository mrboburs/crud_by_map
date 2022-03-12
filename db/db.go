package db

import (
	"crudbymap/model"
	"errors"
	// "fmt"
)

// type articleStorage map[int]models.Article
func NewArticleStorage() articleStorage {
	return articleStorage{
		data: map[int]model.Article{},
	}

}

type articleStorage struct {
	data map[int]model.Article
}

//Errors varableName
var ErrorAlreadyExists = errors.New("already exists id")
var ErrorNotFound = errors.New("not found id ")
var ErrorNotUpdated = errors.New("not update because not found id")
var ErrorSearch = errors.New("not found data")

//Successes varableName
var SuccessUpdated = errors.New("SuccessFull Updated")
var SuccessAdded = errors.New("SuccessFull Added")
var SuccessDeleted = errors.New("SuccessFull Deleted")
var SuccessGetByID = errors.New("SuccessFull Got by ID")

//add new article
func (storage *articleStorage) Add(entity model.Article) error {
	if _, ok := storage.data[entity.ID]; ok {
		// fmt.Println(ok, 2, x)
		return ErrorAlreadyExists
	}
	storage.data[entity.ID] = entity
	return SuccessAdded
}

//get article by id return map
func (storage *articleStorage) GetByID(ID int) (model.Article, error) {
	var resp model.Article
	resp, ok := storage.data[ID]
	if ok {
		return resp, SuccessGetByID
	}
	return resp, ErrorNotFound
}

// get all article return array
func (storage *articleStorage) GetList() []model.Article {
	var resp []model.Article
	for _, data := range storage.data {
		resp = append(resp, data)
	}
	return resp
}

//search article via title return array
func (storage *articleStorage) Search(str string) []model.Article {
	var resp []model.Article
	for _, data := range storage.data {
		if data.Content.Title == str {
			resp = append(resp, data)
		}
	}
	return resp
}

//update exist article
func (storage *articleStorage) Update(entity model.Article) error {
	if _, ok := storage.data[entity.ID]; ok {
		storage.data[entity.ID] = entity
		return SuccessUpdated
	}
	return ErrorNotUpdated
}

//delete article via id
func (storage *articleStorage) Delete(ID int) error {
	if _, ok := storage.data[ID]; ok {
		delete(storage.data, ID)
		return SuccessDeleted
	}
	return ErrorNotFound
}
