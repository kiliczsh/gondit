package handler

import (
	"fmt"
	"github.com/kiliczsh/gondit/database"
	"github.com/kiliczsh/gondit/model"
	"github.com/kiliczsh/gondit/services"

	"github.com/gofiber/fiber/v2"
)

// GetAllScans query all scans
func GetAllScans(c *fiber.Ctx) error {
	db := database.DB
	var scans []model.Scan
	db.Find(&scans)
	return c.JSON(fiber.Map{"status": "success", "message": "All scans", "data": scans })
}

// GetScan query scan
func GetScan(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var scan model.Scan
	db.Find(&scan, id)
	if scan.Url == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No scan found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Scan found", "data": scan})
}

func Test(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var scan model.Scan
	db.Find(&scan, id)
	res := services.Clone(scan)
	if res != nil {
		return c.JSON(fiber.Map{ "success": false })
	}
	go services.Execute(db, scan)
	return c.JSON(fiber.Map{  "success": true, "scan_id": id, "scan" : scan.Url})
}

// CreateScan new scan
func CreateScan(c *fiber.Ctx) error {
	db := database.DB
	scan := new(model.Scan)
	if err := c.BodyParser(scan); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create scan", "data": err})
	}
	db.Create(&scan)
	fmt.Println(scan.Url)
	res := services.Clone(*scan)
	if res != nil {
		return c.JSON(fiber.Map{ "success": false })
	}
	go services.Execute(db, *scan)
	return c.JSON(fiber.Map{ "status": "success", "message": "Scan successfully created!",
		"scan_id": scan.ID, "scan" : scan.Url })
}

// DeleteScan delete scan
func DeleteScan(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var scan model.Scan
	db.First(&scan, id)
	if scan.Url == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No scan found with ID", "data": nil})

	}
	db.Delete(&scan)
	return c.JSON(fiber.Map{"status": "success", "message": "Scan successfully deleted!", "data": nil})
}