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

	// -- Repositorys
	patientDb := repositories.NewPatientDataMySQL(mySQLclient)
	areaDb := repositories.NewAreaDataMySQL(mySQLclient)
	userDb := repositories.NewUserDataMySQL(mySQLclient)
	profileDb := repositories.NewProfileDataMySQL(mySQLclient)
	citationDb := repositories.NewCitationDataMySQL(mySQLclient)

	// -- Service
	patientService := service.NewPatientService(patientDb)
	areaService := service.NewAreaService(areaDb)
	userService := service.NewUserService(userDb)
	profileService := service.NewProfileService(profileDb)
	citationService := service.NewCitationService(citationDb)

	// -- Handler
	patientHandler := PacientHandler{patientService}
	areaHandler := AreaHandler{areaService}
	userHandler := UserHandler{userService}
	profileHandler := ProfileHandler{profileService}
	citationHandler := CitationHandler{citationService}

	apiPacient := app.Group("/paciente")
	apiPacient.Get("/", patientHandler.GetAllPatient)
	apiPacient.Get("/:idPatient", patientHandler.GetPatient)
	apiPacient.Post("/", patientHandler.PostPatient)
	apiPacient.Put("/:idPatient", patientHandler.PutPatient)
	apiPacient.Delete("/:idPatient", patientHandler.DeletePatient)

	apiArea := app.Group("/area")
	apiArea.Get("/", areaHandler.GetAllAreas)
	apiArea.Get("/:idArea", areaHandler.GetArea)
	apiArea.Post("/", areaHandler.PostArea)
	apiArea.Put("/:idArea", areaHandler.PutArea)
	apiArea.Delete("/:idArea", areaHandler.DeleteArea)

	apiUser := app.Group("/user")
	apiUser.Get("/", userHandler.GetAllUser)
	apiUser.Get("/:idUser", userHandler.GetUser)
	apiUser.Post("/", userHandler.PostUser)
	apiUser.Put("/:idUser", userHandler.PutUser)
	apiUser.Delete("/:idUser", userHandler.DeleteUser)

	apiProfile := app.Group("/profile")
	apiProfile.Get("/", profileHandler.GetAllProfiles)
	apiProfile.Get("/:idProfile", profileHandler.GetProfile)
	apiProfile.Post("/", profileHandler.PostProfile)
	apiProfile.Put("/:idProfile", profileHandler.PutProfile)
	apiProfile.Delete("/:idProfile", profileHandler.DeleteProfile)

	apiCitation := app.Group("/citation")
	apiCitation.Get("/", citationHandler.GetAllCitations)
	apiCitation.Get("/:idCitation", citationHandler.GetCitation)
	apiCitation.Post("/", citationHandler.PostCitation)
	apiCitation.Put("/:idCitation", citationHandler.PutCitation)
	apiCitation.Delete("/:idCitation", citationHandler.DeleteCitation)

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
