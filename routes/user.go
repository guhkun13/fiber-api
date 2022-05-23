package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/guhkun13/tutorial/fiber-api/database"
	"github.com/guhkun13/tutorial/fiber-api/models"
)

type UserDto struct {
	// DTO or serializer
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UpdateUserDto struct {
	FirstName string
	LastName string
}


func CreateResponseUser(userModel models.User) UserDto {
	return UserDto{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	ret := CreateResponseUser(user)
	
	return c.Status(http.StatusOK).JSON(ret)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []UserDto{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(http.StatusOK).JSON(responseUsers)
}


func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please send correct id")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	ret := CreateResponseUser(user)

	return c.Status(200).JSON(ret)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateUser UpdateUserDto

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateUser.FirstName
	user.LastName = updateUser.LastName

	database.Database.Db.Save(&user)

	ret := CreateResponseUser(user)

	return c.Status(200).JSON(ret)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please send correct id")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	
	// ret := CreateResponseUser(user)
	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Successfully delete user")
}