package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
)

type AreaHandler struct {
	service service.AreaService
}

func (ah AreaHandler) GetAllAreas(c *fiber.Ctx) error {
	res, err := ah.service.GetAllAreas()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (ah AreaHandler) GetArea(c *fiber.Ctx) error {
	id := c.Params("idArea")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ah.service.GetArea(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah AreaHandler) PostArea(c *fiber.Ctx) error {
	req := new(dto.AreaRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ah.service.PostArea(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah AreaHandler) PutArea(c *fiber.Ctx) error {
	id := c.Params("idArea")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.AreaRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ah.service.UpdateArea(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah AreaHandler) DeleteArea(c *fiber.Ctx) error {
	id := c.Params("idArea")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ah.service.DeleteArea(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
