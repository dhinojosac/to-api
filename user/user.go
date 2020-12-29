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
	db := database.DB
	db.First(&user, dataUU.Email)
	c.JSON(user)
}

//GetUsers ...
func GetUsers(c *fiber.Ctx) {
	db := database.DB
	var users []User
	db.Find(&users)
	c.JSON(users)
}

//GetUser ...
func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var user User
	db.Find(&user, id)
	c.JSON(user)
}

//NewUser ...
func NewUser(c *fiber.Ctx) {
	db := database.DB
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&user)
	c.JSON(user)
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
	id := c.Params("id")
	db := database.DB
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
	id := c.Params("id")
	db := database.DB

	var user User
	db.First(&user, id)
	fmt.Println(&user)
	if user.FirstName == "" {
		c.Status(500).Send("No Patient Found with ID")
		return
	}
	db.Delete(&user)
}
