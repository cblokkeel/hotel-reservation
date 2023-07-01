package api

import (
	"context"

	"github.com/cblokkeel/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(store db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: store,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	users, err := h.userStore.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	return c.JSON(users)
}
