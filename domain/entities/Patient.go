package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Patient struct {
	PatientId       int64  `db:"patient_id"`
	PatientName     string `db:"patient_name"`
	PatientLastName string `db:"patient_last_name"`
	PatientAge      int64  `db:"patient_age"`
}

func (p Patient) ToPatientResponse() *dto.PatientResponse {
	return &dto.PatientResponse{PatientId: p.PatientId}
}

func NewPatient(patientId int64, patientName string, patientLastName string, patientAge int64) Patient {
	return Patient{
		patientId,
		patientName,
		patientLastName,
		patientAge,
	}
}
