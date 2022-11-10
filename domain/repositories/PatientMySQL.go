package repositories

import (
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type PatientRepository interface {
	FindAllPatients() ([]entities.Patient, error)
	//FinPatientById(id int64) (*entities.Patient, error)
	//SavePatient(patient entities.Patient) ([]entities.Patient, error)
	//UpdatePatient(patient entities.Patient) ([]entities.Patient, error)
	//DeletePatient(id int64) ([]entities.Patient, error)
}

type PatientDatabaseMySQL struct {
	//client *sqlx.DB
	patients []entities.Patient
}

func (db PatientDatabaseMySQL) FindAllPatients() ([]entities.Patient, error) {
	return db.patients, nil
}
func NewPatientDataMySQL() PatientDatabaseMySQL {
	patients := []entities.Patient{
		{PatientId: 1, PatientName: "Arnold", PatientLastName: "Chavez", PatientAge: 23},
		{PatientId: 2, PatientName: "Kevin", PatientLastName: "Burgos", PatientAge: 21},
	}
	return PatientDatabaseMySQL{patients}
}

/*

func (db PatientDatabaseMySQL) PatientById(id int64) (*entities.Patient, error) {
	patients := []entities.Patient{
		{1, "Arnold", "Chavez", 23},
		{2, "Kevin", "Burgos", 21},
	}

	var patientId entities.Patient

	for _, item := range patients {
		if item.PatientId == id {
			patientId = item
		}
	}
	return &patientId, nil
}

func (db PatientDatabaseMySQL) SavePatient(patient entities.Patient) ([]entities.Patient, error) {
	patients := []entities.Patient{
		{1, "Arnold", "Chavez", 23},
		{2, "Kevin", "Burgos", 21},
	}

	patients.append(patient)

	return patients, nil
}

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
