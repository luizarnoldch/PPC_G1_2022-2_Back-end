package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
)

// Esturctura de datos que instaciarĂ¡ los servicios /*IMPORT*/

type PacientHandler struct {
	service service.PacientService
}

func (ph PacientHandler) GetAllPatient(c *fiber.Ctx) error {
	res, err := ph.service.GetAllPatients()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (ph PacientHandler) GetPatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ph.service.GetPatient(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ph PacientHandler) PostPatient(c *fiber.Ctx) error {
	req := new(dto.PatientRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ph.service.PostPatient(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ph PacientHandler) PutPatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.PatientRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ph.service.UpdatePatient(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ph PacientHandler) DeletePatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ph.service.DeletePatient(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
