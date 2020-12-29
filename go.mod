module github.com/dhinojosac/to-api

go 1.15

require (
	database/database v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber v1.14.6
	github.com/jinzhu/gorm v1.9.16
	patient/patient v0.0.0-00010101000000-000000000000
	user/user v0.0.0-00010101000000-000000000000
)

replace patient/patient => C:\Users\dhinojosac\go\src\github.com\dhinojosac\to-api\patient

replace database/database => C:\Users\dhinojosac\go\src\github.com\dhinojosac\to-api\database

replace user/user => C:\Users\dhinojosac\go\src\github.com\dhinojosac\to-api\user
