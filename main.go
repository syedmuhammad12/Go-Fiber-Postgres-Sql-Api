package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"mode/models"
	"mode/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type GetVariable struct {
	Variable string `json:"variable"`
	Value    uint   `json:"value"`
}

type Repository struct {
	DB *gorm.DB
}

var lis []GetVariable

func (r *Repository) resetVariableA(context *fiber.Ctx) error {

	err := r.DB.Exec(`Update test_vars SET "value" = 0 WHERE "variable" = 'a'`)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete GetVariable",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "variable a is resetted to 0 successfully",
	})
	return nil
}

func (r *Repository) resetVariableB(context *fiber.Ctx) error {

	err := r.DB.Exec(`Update test_vars SET "value" = 0 WHERE "variable" = 'b'`)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete GetVariable",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "variable b is resetted to 0 successfully",
	})
	return nil
}

func (r *Repository) GetVariableA(context *fiber.Ctx) error {

	id := "a"
	GetVariableModel := &models.TestVars{}

	err := r.DB.Where("variable = ?", id).First(GetVariableModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the GetVariable"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "variable a is fetched successfully",
		"data":    GetVariableModel,
	})

	return nil
}

func (r *Repository) GetVariableB(context *fiber.Ctx) error {

	id := "b"
	GetVariableModel := &models.TestVars{}

	err := r.DB.Where("variable = ?", id).First(GetVariableModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the GetVariable"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Variable b is fetched successfully",
		"data":    GetVariableModel,
	})
	return nil
}

func (r *Repository) GetSum(context *fiber.Ctx) error {

	r.DB.Raw(`select * from "test_vars"`).Scan(&lis)
	// fmt.Println(lis[0].Value + lis[1].Value)

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Sum Fetched Successfully",
		"sum":     lis[0].Value + lis[1].Value,
	})
	return nil
}

func (r *Repository) SettingUpA(context *fiber.Ctx) error {

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "value cannot be empty",
		})
		return nil
	}

	ind, error := strconv.Atoi(id)
	if error != nil {
		fmt.Println("error")
	}
	err := r.DB.Exec(`Update "test_vars" SET "value" = ? WHERE "variable" = 'a'`, ind)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not update variable a",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "variable a is updated successfully",
	})
	return nil
}

func (r *Repository) SettingUpB(context *fiber.Ctx) error {

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "value cannot be empty",
		})
		return nil
	}

	ind, error := strconv.Atoi(id)
	if error != nil {
		fmt.Println("error")
	}
	err := r.DB.Exec(`Update test_vars SET "value" = ? WHERE "variable" = 'b'`, ind)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not update variable b",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "variable b is updated successfully",
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api_a := app.Group("/a")
	api_a.Post("/setvariable/:id", r.SettingUpA)
	api_a.Delete("/deletevariable", r.resetVariableA)
	api_a.Get("/getvariable", r.GetVariableA)

	api_b := app.Group("/b")
	api_b.Post("/setvariable/:id", r.SettingUpB)
	api_b.Delete("/deletevariable", r.resetVariableB)
	api_b.Get("/getvariable", r.GetVariableB)

	api_c := app.Group("/sum")
	api_c.Get("/getsum", r.GetSum)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.NewConnection()

	if err != nil {
		log.Fatal("could not load the database")
	}
	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")
}
