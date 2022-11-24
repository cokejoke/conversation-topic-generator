package resource

import (
	"conversation-topic/model"
	"conversation-topic/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type PackResource struct{}

var validate *validator.Validate

func (resource PackResource) Init(app *fiber.App) {

	validate = validator.New()

	repo := repository.InitPackRepository()

	router := app.Group("/pack")

	router.Get("/", func(ctx *fiber.Ctx) error {
		search := ctx.Query("q")
		if search != "" {
			return ctx.JSON(repo.Search(search))
		}

		pageStr, limitStr := ctx.Query("page"), ctx.Query("limit")
		if pageStr != "" && limitStr != "" {
			page, err := strconv.ParseInt(pageStr, 10, 32)
			if err != nil {
				page = 1
			}
			limit, err := strconv.ParseInt(limitStr, 10, 32)
			if err != nil {
				limit = 16
			}
			return ctx.JSON(repo.ListPaged(int(page), int(limit)))
		}

		return ctx.JSON(repo.List())
	})

	router.Get("/count", Auth, func(ctx *fiber.Ctx) error {
		return ctx.JSON(repo.Count())
	})

	router.Get("/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.SendStatus(404)
		}

		result := repo.FindById(id)
		if result == nil {
			return ctx.SendStatus(404)
		}

		return ctx.JSON(result)
	})

	router.Post("/", Auth, func(ctx *fiber.Ctx) error {
		pack := &model.Pack{}
		err := ctx.BodyParser(pack)

		if err != nil {
			return ctx.SendStatus(500)
		}

		err = validate.Struct(pack)
		if err != nil {
			return ctx.SendStatus(500)
		}

		if repo.ExistsById(pack.ID) {
			return ctx.SendStatus(403)
		}

		repo.Create(pack)
		return ctx.JSON(pack)
	})

	router.Put("/", Auth, func(ctx *fiber.Ctx) error {
		pack := &model.Pack{}
		err := ctx.BodyParser(pack)

		if err != nil {
			return ctx.SendStatus(500)
		}

		err = validate.Struct(pack)
		if err != nil {
			return ctx.SendStatus(500)
		}

		if repo.ExistsById(pack.ID) {
			repo.Update(pack)
		} else {
			repo.Create(pack)
		}

		return ctx.JSON(pack)
	})

	router.Delete("/:id", Auth, func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.SendStatus(404)
		}

		pack := repo.FindById(id)

		if pack == nil {
			return ctx.SendStatus(404)
		}

		repo.DeleteById(pack.ID)
		return ctx.JSON(pack)
	})

}
