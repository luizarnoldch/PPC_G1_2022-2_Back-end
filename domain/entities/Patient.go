package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Patient struct {
	PatientId       int64  `db:"id_patient"`
	PatientName     string `db:"name_patient"`
	PatientLastName string `db:"last_name_patient"`
	PatientAge      int64  `db:"age_patient"`
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
