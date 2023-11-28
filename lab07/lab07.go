package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)
var idcanuse int

type Book struct {
	// TODO: Finish struct
	Id  int
	Name string
	Pages int
}

var bookshelf = []Book{
	// TODO: Init bookshelf
	Book{Id: 1, Name: "Blue Bird", Pages: 500},
}

func getBooks(c *gin.Context) {
	books := []gin.H{}
	for _, book := range bookshelf {
		books = append(books, gin.H{
			"id":    book.Id,
			"name":  book.Name,
			"pages": book.Pages,
		})
	}
	c.JSON(200, books)
}
func getBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return	
	}

	for _, book := range bookshelf {
		if book.Id == id {
			c.JSON(200, gin.H{
				"id":    book.Id,
				"name":  book.Name,
				"pages": book.Pages,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "book not found",
	})
}
func addBook(c *gin.Context) {
	var Bookadded Book
	err := c.BindJSON(&Bookadded)
	if err != nil{
		return
	}
	Name := Bookadded.Name
	Pages := Bookadded.Pages

	for _, book := range bookshelf {
		if book.Name == Name {
			c.JSON(409, gin.H{
				"message": "duplicate book name",
			})
			return
		}
	}
	
	bookshelf = append(bookshelf, Book{Id: idcanuse, Name: Name, Pages: Pages})
	c.JSON(201, gin.H{
		"id":    idcanuse,
		"name":  Name,
		"pages": Pages,
	})
	idcanuse++
}
func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return	
	}
	for i, book := range bookshelf {
		if book.Id == id{
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			break
		}
	}
	c.Status(204)
	c.Abort()
}
func updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return	
	}

	var updateBook Book
	err = c.BindJSON(&updateBook)
	if err != nil{
		return
	}

	var index = -1
	for i ,book := range bookshelf{
		if book.Id == id{
			index = i
		}
		if book.Name == updateBook.Name{
			c.JSON(409, gin.H{
				"message": "duplicate book name",
			})
			return
		}
	}

	if index == -1{
		c.JSON(404, gin.H{
			"message": "book not found",
		})
		return
	}
	bookshelf[index].Name = updateBook.Name
	bookshelf[index].Pages = updateBook.Pages
	c.JSON(200, gin.H{
		"id": bookshelf[index].Id,
		"name": bookshelf[index].Name,
		"pages": bookshelf[index].Pages,
	})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	idcanuse = 2
	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	err := r.Run(":8087")
	if err != nil {
		return
	}
}
