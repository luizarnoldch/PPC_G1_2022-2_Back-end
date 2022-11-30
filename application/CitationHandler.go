package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
)

type CitationHandler struct {
	service service.CitationService
}

func (ch CitationHandler) GetAllCitations(c *fiber.Ctx) error {
	res, err := ch.service.GetAllCitations()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (ch CitationHandler) GetCitation(c *fiber.Ctx) error {
	id := c.Params("idCitation")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ch.service.GetCitation(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch CitationHandler) PostCitation(c *fiber.Ctx) error {
	req := new(dto.CitationRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ch.service.PostCitation(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch CitationHandler) PutCitation(c *fiber.Ctx) error {
	id := c.Params("idCitation")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.CitationRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ch.service.UpdateCitation(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch CitationHandler) DeleteCitation(c *fiber.Ctx) error {
	id := c.Params("idCitation")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ch.service.DeleteCitation(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
