package main

import (
	"fmt"
	"patient/patient"
	"user/user"

	"database/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/patient", patient.GetPatients)
	v1.Get("/patient/:id", patient.GetPatient)
	v1.Post("/patient", patient.NewPatient)
	v1.Patch("/patient/:id", patient.UpdatePatient)
	v1.Delete("/patient/:id", patient.DeletePatient)

	v1.Post("/login", user.Check)
	v1.Post("/user", user.NewUser)
	v1.Get("/user", user.GetUsers)
}

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DB.AutoMigrate(&patient.Patient{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(":3002")

	defer database.DB.Close()
}
