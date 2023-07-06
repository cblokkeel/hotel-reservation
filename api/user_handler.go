package api

import (
	"github.com/cblokkeel/hotel-reservation/constants"
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
	id := c.Params(constants.IdParam)
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

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params(constants.IdParam)
	var params types.UpdateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.Validate(); err != nil {
		return err
	}
	res, err := h.userStore.UpdateUser(c.Context(), id, params)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"updatedId": res.Hex()})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params(constants.IdParam)
	res, err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]int64{"deleted": res})
}
