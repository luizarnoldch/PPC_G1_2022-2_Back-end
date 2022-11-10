package service

import (
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/repositories"
)

type PacientService interface {
	GetAllPatients() ([]dto.PatientResponse, error)
	//GetPatient(id int64) (*dto.PatientResponse, error)
}

type DefaultPacientService struct {
	db repositories.PatientRepository
}

func NewPatientService(db repositories.PatientRepository) DefaultPacientService {
	return DefaultPacientService{db}
}

func (s DefaultPacientService) GetAllPatients() ([]dto.PatientResponse, error) {
	patients, err := s.db.FindAllPatients()
	if err != nil {
		return nil, err
	}
	response := make([]dto.PatientResponse, 0)
	for _, p := range patients {
		response = append(response, *p.ToPatientResponse())
	}
	return response, nil
}

/*
func (s DefaultPacientService) GetPatient(id int64) (*dto.PatientResponse, error) {
	patient, err := s.db.FinPatientById(id)
	if err != nil {
		return nil, err
	}
	response := patient
	return &response, nil
}
*/
