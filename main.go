package main

import (
	// "crudbymap/cmd/model"
	"crudbymap/model"
	"fmt"
	"time"
)

var myMap = map[string]model.Article{}

func Delete(id int) {
	for key, val := range myMap {
		if val.ID == id {
			val.Content.Body = "updatedBody"
			delete(myMap, key)
			// fmt.Println(val)
			fmt.Println(myMap)
		}
	}
}
func Update(id int) {
	for _, val := range myMap {
		if val.ID == id {
			val.Content.Body = "updatedBody"
			fmt.Println(val)
			// fmt.Println(myMap)
		}
	}
}

func Create() {

	myMap["item1"] = model.Article{

		ID:        1,
		Content:   model.Content{Title: "heading", Body: "main"},
		Author:    model.Person{Firstname: "bob", Lastname: "Mr"},
		CreatedAt: &time.Time{},
	}
	myMap["item2"] = model.Article{
		ID:        2,
		Content:   model.Content{Title: "heading2", Body: "main2"},
		Author:    model.Person{Firstname: "bob2", Lastname: "Mr2"},
		CreatedAt: &time.Time{},
	}

	myMap["item3"] = model.Article{
		ID:        3,
		Content:   model.Content{Title: "heading3", Body: "main3"},
		Author:    model.Person{Firstname: "bob3", Lastname: "Mr3"},
		CreatedAt: &time.Time{},
	}
	// fmt.Println(myMap)

}
func main() {

	Create()
	// Update(1)
	Delete(1)
}
