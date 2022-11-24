package resource

import (
	"conversation-topic/config"
	"conversation-topic/model"
	"conversation-topic/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func Auth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if !strings.Contains(authHeader, " ") {
		return ctx.SendStatus(401)
	}
	tokenString := strings.Split(authHeader, " ")[1]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Container.Secret), nil
	})

	if err != nil {
		return ctx.SendStatus(401)
	}

	user := model.User{
		ID:       int(claims["id"].(float64)),
		Name:     claims["name"].(string),
		Username: claims["username"].(string),
	}
	ctx.Locals("user", user)

	return ctx.Next()
}

type AuthResource struct{}

func (resource AuthResource) Init(app *fiber.App) {

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	repo := repository.InitUserRepository()

	app.Post("/login", func(ctx *fiber.Ctx) error {
		loginRequest := &LoginRequest{}
		err := ctx.BodyParser(loginRequest)

		if err != nil {
			return ctx.SendStatus(500)
		}

		//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
		user := repo.FindByUsername(loginRequest.Username)

		if user == nil {
			return ctx.SendStatus(404)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
			return ctx.SendStatus(401)
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"name":     user.Name,
			"username": user.Username,
			"exp":      jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*24).Unix(), 0)),
		})

		token, err := claims.SignedString([]byte(config.Container.Secret))

		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(fiber.Map{
				"error": "could not login",
			})
		}

		return ctx.JSON(fiber.Map{
			"token": token,
		})
	})

	app.Get("/auth", Auth, func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	})

}
