package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	mod "main/app/models/auth"
	ent "main/internals/interfaces"
	cm "main/pkg/commons"
)

var (
	Validate *validator.Validate = validator.New()
)

func Register(c *fiber.Ctx) error {
	data := ent.RegisterReq{}

	if errParser := c.BodyParser(&data); errParser != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errParser.Error(),
			"data":    nil,
		})
	}

	errValidate := Validate.Struct(data)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidate.Error(),
			"data":    nil,
		})
	}
	Record := mod.Register(data)
	if Record == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error on DB",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   true,
		"message": "Succcess",
		"data":    nil,
	})
}

func Login(c *fiber.Ctx) error {
	data := ent.LoginReq{}

	if errParser := c.BodyParser(&data); errParser != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errParser.Error(),
			"data":    nil,
		})
	}

	errValidate := Validate.Struct(data)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidate.Error(),
			"data":    nil,
		})
	}

	Record, authData := mod.Login(data)

	if Record == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":   false,
			"message": "Kayıt Bulunamadı",
			"data":    nil,
		})
	}
	claims := ent.TokenClaim{
		Id:        authData.Id,
		UserTitle: authData.UserTitle,
		UserName:  authData.UserName,
		Phone:     authData.Phone,
		City:      authData.City,
		Town:      authData.Town,
		Exp:       0,
		Iat:       0,
	}

	token, errGenToken := cm.GenerateToken(claims)
	if errGenToken != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   false,
			"message": "Başarısiz",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Başarılı",
		"token":   token,
	})
}

func ForgetPassword(c *fiber.Ctx) error {
	data := ent.ForgetPasswordReq{}
	if errParser := c.BodyParser(&data); errParser != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errParser.Error(),
			"data":    nil,
		})
	}

	errValidate := Validate.Struct(data)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidate.Error(),
			"data":    nil,
		})
	}

	//TODO Function from Model

	return nil
}

func Verification(c *fiber.Ctx) error {
	data := ent.Verify{}
	if errParser := c.BodyParser(&data); errParser != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errParser.Error(),
			"data":    nil,
		})
	}

	errValidate := Validate.Struct(data)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errValidate.Error(),
			"data":    nil,
		})
	}

	//TODO Function from Model

	return nil
}
