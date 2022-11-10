package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
)

// Esturctura de datos que instaciará los servicios /*IMPORT*/

type PacientHandler struct {
	service service.PacientService
}

func (ch PacientHandler) GetAllPatient(c *fiber.Ctx) error {
	res, err := ch.service.GetAllPatients()
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (ch PacientHandler) GetPatient(c *fiber.Ctx) error {
	return c.JSON("Paciente: " + c.Params("idClient"))
}

func (ch PacientHandler) PostPatient(c *fiber.Ctx) error {
	return c.JSON("Paciente Creado!")
}

func (ch PacientHandler) PutPatient(c *fiber.Ctx) error {
	return c.JSON("Paciente : " + c.Params("idClient") + " actualizado")
}

func (ch PacientHandler) DeletePatient(c *fiber.Ctx) error {
	return c.JSON("Paciente : " + c.Params("idClient") + " elimination")
}
