package application

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"strconv"
	"sync"
	"time"
)

// Esturctura de datos que instaciará los servicios /*IMPORT*/

type PacientHandler struct {
	service service.PacientService
}

var wg = sync.WaitGroup{}

func (ch PacientHandler) GetAllPatient(c *fiber.Ctx) error {
	start := time.Now()
	res, err := ch.service.GetAllPatients()
	if err != nil {
		return err
	}
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
	return c.JSON(res)
}

func (ch PacientHandler) GetPatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ch.service.GetPatient(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch PacientHandler) PostPatient(c *fiber.Ctx) error {
	req := new(dto.PatientRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ch.service.PostPatient(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch PacientHandler) PutPatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	req := new(dto.PatientRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := ch.service.UpdatePatient(idInt, *req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}

func (ch PacientHandler) DeletePatient(c *fiber.Ctx) error {
	id := c.Params("idPatient")
	idInt, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}
	res, errService := ch.service.DeletePatient(idInt)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
