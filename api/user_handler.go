package api

import (
	"context"

	"github.com/cblokkeel/hotel-reservation/db"
	"github.com/cblokkeel/hotel-reservation/types"
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

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "James",
		LastName:  "Foo",
	}
	return c.JSON(user)
}
