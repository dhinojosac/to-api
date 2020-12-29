package patient

import (
	"fmt"
	"time"

	"database/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Patient ...
type Patient struct {
	gorm.Model
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	DOB         time.Time `json:"dob"`
	City        string    `json:"city"`
	Weigth      string    `json:"weigth"`
	Height      string    `json:"height"`
	Description string    `json:"description"`
}

//GetPatients ...
func GetPatients(c *fiber.Ctx) {
	db := database.PatientsDB
	var patients []Patient
	db.Find(&patients)
	c.JSON(patients)
}

//GetPatient ...
func GetPatient(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.PatientsDB
	var patient Patient
	db.Find(&patient, id)
	c.JSON(patient)
}

//NewPatient ...
func NewPatient(c *fiber.Ctx) {
	db := database.PatientsDB
	patient := new(Patient)
	if err := c.BodyParser(patient); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&patient)
	c.JSON(patient)
}

//UpdatePatient ...
func UpdatePatient(c *fiber.Ctx) {
	type DataUpdatePatient struct {
		FirstName   string    `json:"firstname"`
		LastName    string    `json:"lastname"`
		DOB         time.Time `json:"dob"`
		Description string    `json:"description"`
	}
	var dataUP DataUpdatePatient
	if err := c.BodyParser(&dataUP); err != nil {
		c.Status(503).Send(err)
		return
	}
	var patient Patient
	id := c.Params("id")
	db := database.PatientsDB
	db.First(&patient, id)

	patient = Patient{
		FirstName:   dataUP.FirstName,
		LastName:    dataUP.LastName,
		DOB:         dataUP.DOB,
		Description: dataUP.Description,
	}
	db.Save(&patient)
	c.JSON(patient)
}

//DeletePatient ...
func DeletePatient(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.PatientsDB

	var patient Patient
	db.First(&patient, id)
	fmt.Println(&patient)
	if patient.FirstName == "" {
		c.Status(500).Send("No Patient Found with ID")
		return
	}
	db.Delete(&patient)
}
