package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
)

type ProfileHandler struct {
	service service.ProfileService
}

func (ah ProfileHandler) GetAllProfiles(c *fiber.Ctx) error {
	res, err := ah.service.GetAllProfiles()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (ah ProfileHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("idProfile")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ah.service.GetProfile(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah ProfileHandler) PostProfile(c *fiber.Ctx) error {
	req := new(dto.ProfileRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ah.service.PostProfile(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah ProfileHandler) PutProfile(c *fiber.Ctx) error {
	id := c.Params("idProfile")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.ProfileRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ah.service.UpdateProfile(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ah ProfileHandler) DeleteProfile(c *fiber.Ctx) error {
	id := c.Params("idProfile")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ah.service.DeleteProfile(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
