package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type PatientRepository interface {
	FindAllPatients() ([]entities.Patient, error)
	FinPatientById(id int64) (*entities.Patient, error)
	SavePatient(patient entities.Patient) ([]entities.Patient, error)
	//UpdatePatient(patient entities.Patient) ([]entities.Patient, error)
	//DeletePatient(id int64) ([]entities.Patient, error)
}

type PatientDatabaseMySQL struct {
	client *sqlx.DB
	//patients []entities.Patient
}

func (db PatientDatabaseMySQL) FindAllPatients() ([]entities.Patient, error) {
	var err error
	patients := make([]entities.Patient, 0)

	patientsQuery := "SELECT * FROM posta_ppc.patients"

	err = db.client.Select(&patients, patientsQuery)

	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (db PatientDatabaseMySQL) FinPatientById(id int64) (*entities.Patient, error) {

	var err error
	var patient entities.Patient

	patientQuery := "SELECT * FROM posta_ppc.patients WHERE id_patient = ?"

	err = db.client.Get(&patient, patientQuery, id)

	if err == sql.ErrNoRows {
		return nil, errors.New("No patient found")
	}

	if err != nil {
		return nil, errors.New("Unexpected database error")
	}
	return &patient, nil
}

func (db PatientDatabaseMySQL) SavePatient(patient entities.Patient) ([]entities.Patient, error) {
	patients := []entities.Patient{
		{1, "Arnold", "Chavez", 23},
		{2, "Kevin", "Burgos", 21},
	}

	patients.append(patient)

	return patients, nil
}

func NewPatientDataMySQL(db *sqlx.DB) PatientDatabaseMySQL {
	return PatientDatabaseMySQL{db}
}

/*
func NewPatientDataMySQL() PatientDatabaseMySQL {
	patients := []entities.Patient{
		{PatientId: 1, PatientName: "Arnold", PatientLastName: "Chavez", PatientAge: 23},
		{PatientId: 2, PatientName: "Kevin", PatientLastName: "Burgos", PatientAge: 21},
	}
	return PatientDatabaseMySQL{patients}
}
*/

/*





func (db PatientDatabaseMySQL) UpdatePatient(patient entities.Patient) ([]entities.Patient, error) {
	patients := []entities.Patient{
		{1, "Arnold", "Chavez", 23},
		{2, "Kevin", "Burgos", 21},
	}
	return patients, nil
}

func (db PatientDatabaseMySQL) DeletePatient(id int64) (*entities.Patient, error) {
	patients := []entities.Patient{
		{1, "Arnold", "Chavez", 23},
		{2, "Kevin", "Burgos", 21},
	}
	return patients, nil
}
*/
