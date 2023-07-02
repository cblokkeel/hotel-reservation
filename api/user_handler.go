package api

import (
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
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleInsertUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := types.ParamsToUserValidated(params)
	if err != nil {
		return err
	}
	retrievedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(retrievedUser)
}
