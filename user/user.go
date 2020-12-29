package user

import (
	"fmt"

	"database/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//User ...
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Check ...
func Check(c *fiber.Ctx) {
	/*
		type DataAdminUser struct {
			FirstName string `json:"firstname"`
			LastName  string `json:"lastname"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}
	*/
	var dataUU User
	if err := c.BodyParser(&dataUU); err != nil {
		c.Status(503).Send(err)
		return
	}
	var user User
	db := database.UsersDB
	db.First(&user, dataUU.Email)
	c.JSON(user)
}

//GetUsers ...
func GetUsers(c *fiber.Ctx) {
	db := database.UsersDB
	var users []User
	db.Find(&users)
	c.JSON(users)
}

//GetUser ...
func GetUser(c *fiber.Ctx) {
	id := c.Params("ID")
	db := database.UsersDB
	var user User
	db.Find(&user, id)
	c.JSON(user)
}

//NewUser ...
func NewUser(c *fiber.Ctx) {
	db := database.UsersDB
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}
	var user2 User
	email := c.Params("e,ail")
	db.First(&user2, email)
	if user2.Email == user.Email {
		fmt.Println("User already exists")
		c.Status(409).Send("User already exists")
		return
	} else {
		db.Create(&user)
		c.JSON(user)
	}

}

//UpdateUser ...
func UpdateUser(c *fiber.Ctx) {
	type DataUpdateUser struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
	}
	var dataUU DataUpdateUser
	if err := c.BodyParser(&dataUU); err != nil {
		c.Status(503).Send(err)
		return
	}
	var user User
	id := c.Params("ID")
	db := database.UsersDB
	db.First(&user, id)

	user = User{
		FirstName: dataUU.FirstName,
		LastName:  dataUU.LastName,
		Email:     dataUU.Email,
	}
	db.Save(&user)
	c.JSON(user)
}

//DeleteUser ...
func DeleteUser(c *fiber.Ctx) {
	id := c.Params("ID")
	db := database.UsersDB

	var user User
	db.First(&user, id)
	fmt.Println(&user)
	if user.FirstName == "" {
		c.Status(500).Send("No Patient Found with ID")
		return
	}
	db.Delete(&user)
}
