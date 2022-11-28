package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
)

type UserHandler struct {
	service service.UserService
}

func (uh UserHandler) GetAllUser(c *fiber.Ctx) error {
	res, err := uh.service.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (uh UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("idUser")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := uh.service.GetUser(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (uh UserHandler) PostUser(c *fiber.Ctx) error {

	req := new(dto.UserRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := uh.service.PostUser(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (uh UserHandler) PutUser(c *fiber.Ctx) error {

	id := c.Params("idUser")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.UserRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := uh.service.UpdateUser(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (uh UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("idUser")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := uh.service.DeleteUser(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
