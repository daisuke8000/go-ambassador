package controllers

import (
	"ambassador-backend/src/database"
	"ambassador-backend/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		IsAmbassador: false,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	// AND Patern
	//database.DB.Where( "email = ? AND first_name = ?", data["email"], data["first_name"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mesasge": "User not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

	payload := jwt.StandardClaims{
		Subject: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mesasge": "Invalid Credentials",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// memo
	//type Cookie struct {
	//	Name     string    `json:"name"`
	//	Value    string    `json:"value"`
	//	Path     string    `json:"path"`
	//	Domain   string    `json:"domain"`
	//	MaxAge   int       `json:"max_age"`
	//	Expires  time.Time `json:"expires"`
	//	Secure   bool      `json:"secure"`
	//	HTTPOnly bool      `json:"http_only"`
	//	SameSite string    `json:"same_site"`
	//}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success!",
	})
}