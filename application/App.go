package application

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/repositories"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/service"
	"log"
	"os"
	"time"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar archivo .env ")
	}

	app := fiber.New()
	app.Use(cors.New())
	/*
		app.Use(cors.New(cors.config{
			AllowHeaders: "Content-Type, Origin, Accept"
		}))
	*/

	// Creación de instancia Handler
	// -- DB
	mySQLclient := GetMySQLClient()
	patientDb := repositories.NewPatientDataMySQL(mySQLclient)
	// -- Service
	pService := service.NewPatientService(patientDb)
	// -- Handler
	pHandler := PacientHandler{pService}

	app.Get("/", pHandler.GetAllPatient)
	app.Get("/:idPatient", pHandler.GetPatient)
	app.Post("/", pHandler.PostPatient)
	app.Put("/:idPatient", pHandler.PutPatient)
	app.Delete("/:idPatient", pHandler.DeletePatient)

	PORT := os.Getenv("PORT")

	app.Listen(":" + PORT)
}

func GetMySQLClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
