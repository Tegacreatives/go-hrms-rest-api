package main

import (
	"example/hello/database"
	"example/hello/lead"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Setup api routes
func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// Creates lead table in database
func performMigration(db *gorm.DB) {
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

// Connects to database and passes the db to performMigration function
func initDatabase() *gorm.DB {
	db := database.ConnectToDB()
	performMigration(db)
	return db
}

// Creates a new fiber app, performs the db initializaton and set up the routes
func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(":3000")
}
