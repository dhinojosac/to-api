package user

import (
	"fmt"
	"time"

	"database/database"

	"github.com/dgrijalva/jwt-go"
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

//Login ...
func Login(c *fiber.Ctx) {
	var user User
	if err := c.BodyParser(&user); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Throws Unauthorized error
	if user.Email != "d.hinojosa.cordova@gmail.com" || user.Password != "admin" {
		c.Status(401).Send("Unauthorized")
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "admin"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.Status(409).Send("cannot create token")
		return
	}

	c.JSON(fiber.Map{"token": t, "user": user.Email})

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
