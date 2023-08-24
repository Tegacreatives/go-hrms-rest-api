package lead

import (
	"example/hello/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User Information
type Lead struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
}

// Get the leads from the database
func GetLeads(c *fiber.Ctx) error {
	db := database.DBconn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

// Get individual lead from the database
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

// Create a new lead
func NewLead(c *fiber.Ctx) error {
	db := database.DBconn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

// Delete lead from the database
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBconn

	var lead Lead
	db.First(&lead, id)
	if lead.FirstName == "" {
		c.Status(500).SendString("No Lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")

}
