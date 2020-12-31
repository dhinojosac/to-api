package main

import (
	"fmt"
	"patient/patient"
	"user/user"

	"database/database"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	jwtware "github.com/gofiber/jwt"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", user.Login)
	v1.Post("/singup", user.NewUser)

	v1.Post("/user", user.NewUser)
	v1.Get("/user", user.GetUsers)
	v1.Get("/user/:id", user.GetUser)
	v1.Delete("/user/:id", user.DeleteUser)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted routes
	v1.Get("/patient", patient.GetPatients)
	v1.Get("/patient/:id", patient.GetPatient)
	v1.Post("/patient", patient.NewPatient)
	v1.Patch("/patient/:id", patient.UpdatePatient)
	v1.Delete("/patient/:id", patient.DeletePatient)

}

func initDatabase() {
	var err error
	database.PatientsDB, err = gorm.Open("sqlite3", "patients.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.PatientsDB.AutoMigrate(&patient.Patient{})
	fmt.Println("Database Migrated")

	database.UsersDB, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.UsersDB.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()

	setupRoutes(app)
	app.Listen(":3002")

	defer database.PatientsDB.Close()
	defer database.UsersDB.Close()
}
