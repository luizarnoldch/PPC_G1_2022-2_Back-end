package service

import (
	"errors"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/repositories"
)

type PacientService interface {
	GetAllPatients() ([]dto.PatientResponse, error)
	GetPatient(id int64) (*dto.PatientResponse, error)
	PostPatient(request dto.PatientRequest) (*dto.PatientMessage, error)
	UpdatePatient(id int64, request dto.PatientRequest) (*dto.PatientMessage, error)
	DeletePatient(id int64) (*dto.PatientMessage, error)
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

func (s DefaultPacientService) GetPatient(id int64) (*dto.PatientResponse, error) {
	patient, err := s.db.FinPatientById(id)
	if err != nil {
		return nil, err
	}
	response := *patient.ToPatientResponse()

	return &response, nil
}

func (s DefaultPacientService) PostPatient(req dto.PatientRequest) (*dto.PatientMessage, error) {

	if req.IsUnderAge() {
		return nil, errors.New("patient under 18 age")
	}

	patientSave, err := s.db.SavePatient(req)
	if err != nil {
		return nil, errors.New("can't Save Patient")
	}
	response := *patientSave.ToPatientMessage(patientSave.PatientId, "patient saved")
	return &response, nil
}

func (s DefaultPacientService) UpdatePatient(id int64, req dto.PatientRequest) (*dto.PatientMessage, error) {
	if req.IsUnderAge() {
		return nil, errors.New("patient under 18 age")
	}
	patientUpdate, err := s.db.UpdatePatient(id, req)
	if err != nil {
		return nil, errors.New("can't update Patient")
	}
	response := *patientUpdate.ToPatientMessage(id, "patient updated")
	return &response, nil
}

func (s DefaultPacientService) DeletePatient(id int64) (*dto.PatientMessage, error) {
	patient, err := s.db.DeletePatient(id)
	if err != nil {
		return nil, err
	}
	response := *patient.ToPatientMessage(id, "patient deleted")
	return &response, nil
}
