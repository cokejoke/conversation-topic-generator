package resource

import (
	"conversation-topic/model"
	"conversation-topic/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TopicResource struct{}

func (resource TopicResource) Init(app *fiber.App) {

	repo := repository.InitTopicRepository()
	packRepo := repository.InitPackRepository()

	router := app.Group("/topic")

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

	router.Get("/random", func(ctx *fiber.Ctx) error {
		result := repo.RandomTopic()
		return ctx.JSON(result)
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
		topic := &model.Topic{}
		err := ctx.BodyParser(topic)

		if err != nil {
			return ctx.SendStatus(500)
		}

		if repo.ExistsById(topic.ID) {
			return ctx.SendStatus(403)
		}

		if packRepo.ExistsById(topic.PackID) {
			repo.Create(topic)
			return ctx.JSON(topic)
		}

		return ctx.Status(404).JSON(fiber.Map{
			"message": "No pack with the id " + strconv.Itoa(topic.PackID) + " found",
		})
	})

	router.Put("/", Auth, func(ctx *fiber.Ctx) error {
		topic := &model.Topic{}
		err := ctx.BodyParser(topic)

		if err != nil {
			return ctx.SendStatus(500)
		}

		if repo.ExistsById(topic.ID) {
			repo.Update(topic)
		} else {
			if packRepo.ExistsById(topic.PackID) {
				repo.Create(topic)
			}
		}

		return ctx.JSON(topic)
	})

	router.Delete("/:id", Auth, func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		if err != nil {
			return ctx.SendStatus(404)
		}

		topic := repo.FindById(id)

		if topic == nil {
			return ctx.SendStatus(404)
		}

		repo.DeleteById(topic.ID)
		return ctx.JSON(topic)
	})

}
