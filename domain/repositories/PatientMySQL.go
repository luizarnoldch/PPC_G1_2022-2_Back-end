package repositories

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
)

type PatientRepository interface {
	FindAllPatients() ([]entities.Patient, error)
	FinPatientById(id int64) (*entities.Patient, error)
	SavePatient(req dto.PatientRequest) (*entities.Patient, error)
	UpdatePatient(id int64, req dto.PatientRequest) (*entities.Patient, error)
	DeletePatient(id int64) (*entities.Patient, error)
}

type PatientDatabaseMySQL struct {
	client *sqlx.DB
}

// var wg = sync.WaitGroup{}

func (db PatientDatabaseMySQL) FindAllPatients() ([]entities.Patient, error) {
	//start := time.Now()

	var err error
	patients := make([]entities.Patient, 0)
	patientsQuery := "SELECT * FROM posta_ppc.paciente"
	err = db.client.Select(&patients, patientsQuery)
	if err != nil {
		return nil, err
	}

	//En el caso de ser solo 1 consulta no hya optimización en el tiempo
	/*
		wg.Add(1)
		var err error
		patients := make([]entities.Patient, 0)

		go func(errIn error) error {
			patientsQuery := "SELECT * FROM posta_ppc.patients"
			errIn = db.client.Select(&patients, patientsQuery)
			wg.Done()
			return errIn
		}(err)

		if err != nil {
			return nil, err
		}
		wg.Wait()
	*/
	//elapsed := time.Since(start)
	//fmt.Printf("Processes took %s", elapsed)

	return patients, nil
}

func (db PatientDatabaseMySQL) FinPatientById(id int64) (*entities.Patient, error) {

	var err error
	var patient entities.Patient

	patientQuery := "SELECT * FROM posta_ppc.paciente WHERE ID_Paciente = ?"

	err = db.client.Get(&patient, patientQuery, id)

	if err == sql.ErrNoRows {
		return nil, errors.New("no patient found")
	}

	if err != nil {
		return nil, errors.New("unexpected database error")
	}
	return &patient, nil
}

func (db PatientDatabaseMySQL) SavePatient(req dto.PatientRequest) (*entities.Patient, error) {
	patientQuery := `INSERT INTO 
    					paciente (
							Nombre_Paciente,
							Apellido_Paciente,
							Nick_Paciente,
							Clave_Paciente,
							Foto_Paciente,
							Nacionalidad_Paciente,
							DATE_Nac_Paciente,
							Email_Paciente,
							Telefono_Paciente,
							Celular_Paciente,
							Cedula_Paciente,
							Desc_Paciente,
							Archivo_Paciente,
							Estado_Paciente,
							Tag_Paciente) 
						VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	res, err := db.client.Exec(
		patientQuery,
		req.PatientName,
		req.PatientLastName,
		req.PatientNick,
		req.PatientKey,
		req.PatientPhoto,
		req.PatientNationality,
		req.PatientBirth,
		req.PatientEmail,
		req.PatientTelf,
		req.PatientPhone,
		req.PatientCard,
		req.PatientDiscount,
		req.PatientFile,
		req.PatientState,
		req.PatientTag)
	if err != nil {
		return nil, errors.New("error while saving patient")
	}
	id, errId := res.LastInsertId()
	if errId != nil {
		return nil, errors.New("error while getting id from patient")
	}

	patient := entities.Patient{
		id,
		req.PatientName,
		req.PatientLastName,
		req.PatientNick,
		req.PatientKey,
		req.PatientPhoto,
		req.PatientNationality,
		req.PatientBirth,
		req.PatientEmail,
		req.PatientTelf,
		req.PatientPhone,
		req.PatientCard,
		req.PatientDiscount,
		req.PatientFile,
		req.PatientState,
		req.PatientTag}

	return &patient, nil
}

func (db PatientDatabaseMySQL) UpdatePatient(id int64, req dto.PatientRequest) (*entities.Patient, error) {
	patientQuery := `UPDATE paciente p
    					SET Nombre_Paciente = ?,
							Apellido_Paciente = ?,
							Nick_Paciente = ?,
							Clave_Paciente = ?,
							Foto_Paciente = ?,
							Nacionalidad_Paciente = ?,
							DATE_Nac_Paciente = ?,
							Email_Paciente = ?,
							Telefono_Paciente = ?,
							Celular_Paciente = ?,
							Cedula_Paciente = ?,
							Desc_Paciente = ?,
							Archivo_Paciente = ?,
							Estado_Paciente = ?,
							Tag_Paciente = ?
    					WHERE p.ID_Paciente = ?`
	patientUpdate, err := db.client.Exec(
		patientQuery,
		req.PatientName,
		req.PatientLastName,
		req.PatientNick,
		req.PatientKey,
		req.PatientPhoto,
		req.PatientNationality,
		req.PatientBirth,
		req.PatientEmail,
		req.PatientTelf,
		req.PatientPhone,
		req.PatientCard,
		req.PatientDiscount,
		req.PatientFile,
		req.PatientState,
		req.PatientTag,
		id)
	if err != nil {
		return nil, errors.New("error while updating patient")
	}
	rows, errNoRow := patientUpdate.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting updated rows")
	}
	if rows == 0 {
		return nil, errors.New("no patient updated")
	}
	patientID, errID := patientUpdate.LastInsertId()
	if errID != nil {
		return nil, errors.New("error while getting id from patient")
	}
	patient := entities.NewPatient(
		patientID,
		req.PatientName,
		req.PatientLastName,
		req.PatientNick,
		req.PatientKey,
		req.PatientPhoto,
		req.PatientNationality,
		req.PatientBirth,
		req.PatientEmail,
		req.PatientTelf,
		req.PatientPhone,
		req.PatientCard,
		req.PatientDiscount,
		req.PatientFile,
		req.PatientState,
		req.PatientTag)
	return &patient, nil
}

func (db PatientDatabaseMySQL) DeletePatient(id int64) (*entities.Patient, error) {
	var patient entities.Patient
	patientQuery := "DELETE FROM posta_ppc.paciente WHERE ID_Paciente = ?"

	patientDelete, errDelete := db.client.Exec(patientQuery, id)
	if errDelete != nil {
		return nil, errors.New("error while deleting the patient")
	}
	rows, errNoRow := patientDelete.RowsAffected()
	if errNoRow != nil {
		return nil, errors.New("error while getting deleted rows")
	}
	if rows == 0 {
		return nil, errors.New("no patient deleted")
	}
	patient.PatientId = id
	return &patient, nil
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
