package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-metaverse/db"
	mgorm "go-metaverse/models/gorm"
	"go-metaverse/router"
	"log"
)
//
//var (
//	db  *gorm.DB
//	err error
//)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db.DB, db.Err = gorm.Open("sqlite3", "./api.db")
	if db.Err != nil {
		log.Fatal(db.Err)
	}
	defer db.DB.Close()

	mgorm.AutoMigrate(db.DB)

	//db.AutoMigrate(&User{})

	r := router.InitRouter()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/save", store)
	r.PUT("/users/:id", updateUser)
	r.GET("/users/:id", showById)
	r.GET("/list", list)
}

func showById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	db.DB.Find(&user, id)
	if user.ID == 0 {
		c.JSON(400, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)
}

func updateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	db.DB.Find(&user, id)

	if user.ID == 0 {
		c.JSON(400, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		db.DB.Save(&user)
		c.JSON(200, user)
	}
}

func list(c *gin.Context) {
	var users []User
	fmt.Println(111, users)
	db.DB.Find(&users)

	c.JSON(200, users)
}

func store(c *gin.Context) {
	var user User
	_ = c.BindJSON(&user)
	fmt.Printf("%v", user)
	db.DB.Create(&user)
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	db.DB.First(&user, id)
	if user.ID == 0 {
		c.JSON(400, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		db.DB.Delete(&user)
		c.JSON(200, gin.H{"message": "删除成功"})
	}
}
