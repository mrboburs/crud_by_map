package main

import (
	"crudbymap/db"
	"crudbymap/model"
	"fmt"
	"time"
	// "github.com/gin-gonic/gin"
)

// var inMemory storage.ArticleStorage

func main() {

	// articleStorage = make(storage.ArticleStorage)
	articleStorage := db.NewArticleStorage()

	//add new article
	// r := gin.Default()
	//1 article
	var a1 model.Article
	a1.ID = 1
	a1.Content.Title = "Lorem"
	a1.Content.Body = "Lorem ipsum"
	var a1p model.Person = model.Person{
		Firstname: "John",
		Lastname:  "Doe",
	}
	a1.Author = a1p
	t := time.Now()
	a1.CreatedAt = &t
	// err := articleStorage.Add(a1)
	// if err != nil {
	// 	fmt.Println(err, 1)
	// }
	// r.POST("/add", articleStorage.Add(a1))

	//2 - article
	var a2 model.Article
	a2.ID = 2
	a2.Content.Title = "Hello Nova"
	a2.Content.Body = "NASA"
	var a2p model.Person = model.Person{
		Firstname: "John",
		Lastname:  "Carter",
	}
	a2.Author = a2p
	a2.CreatedAt = &t
	// err = articleStorage.Add(a2)
	// if err != nil {
	// 	fmt.Println(err, 2)
	// }

	//3 article
	var a3 model.Article
	a3.ID = 3
	a3.Content.Title = "Why 13 Reasons"
	a3.Content.Body = "Cinema"
	var a3p model.Person = model.Person{
		Firstname: "Hanna",
		Lastname:  "Backer",
	}
	a3.Author = a3p
	a3.CreatedAt = &t
	// err = articleStorage.Add(a3)
	// if err != nil {
	// 	fmt.Println(err, 3)
	// }
	// err = articleStorage.Add(a1)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//update 1 article update
	var updateA1 model.Article
	updateA1.ID = 1
	updateA1.Content.Title = "Go"
	updateA1.Content.Body = "Golang"
	var updateA1Person model.Person = model.Person{
		Firstname: "MRB",
		Lastname:  "Hero",
	}
	updateA1.Author = updateA1Person

	//delete article  by id
	// deleteArticle := articleStorage.Delete(2)

	//Update article by ID
	// updateArticle := articleStorage.Update(updateA1)

	//Search  Article via Article Title
	searchArticleList := articleStorage.Search("Why 13 Reasons")
	if len(searchArticleList) == 0 {
		fmt.Println(db.ErrorSearch)
	}
	//get all article list
	// articleList := articleStorage.GetList()

	//get by id
	// getArticleByID, err := articleStorage.GetByID(1)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// fmt.Println()
	// fmt.Println("Get List    : ", articleList)
	// fmt.Println()
	// fmt.Println("Create storage: ", articleStorage)
	// fmt.Println()
	// fmt.Println("Get by ID: ", getArticleByID)
	// fmt.Println()
	// fmt.Println("Delete by ID: ", deleteArticle)
	// fmt.Println()
	// fmt.Println("Update by ID: ", updateArticle)
	// fmt.Println()
	// fmt.Println("Search by string:  ", searchArticleList)
	// fmt.Println()
	// fmt.Println("Get List    : ", articleList)
	// fmt.Println()
}
