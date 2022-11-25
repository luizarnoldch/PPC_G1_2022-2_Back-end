package service

import (
	"errors"
	"fmt"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/repositories"
	"sync"
	"time"
	// "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
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
	// Secuencial
	/*
		start := time.Now()
		patients, err := s.db.FindAllPatients()
		if err != nil {
			return nil, err
		}
		fmt.Printf("Pascientes: %d \n", len(patients))
		response := make([]dto.PatientResponse, 0)
		for _, p := range patients {
			response = append(response, *p.ToPatientResponse())
			for i := 0; i < 50000; i++ {
			}
		}
		elapsed := time.Since(start)
		fmt.Printf("Processes took %s \n", elapsed)
		fmt.Println(len(response))
	*/

	// 1 Canal
	/*
		patients, err := s.db.FindAllPatients()
		if err != nil {
			return nil, err
		}

		fmt.Printf("Pascientes: %d \n", len(patients))

		c := make(chan dto.PatientResponse, len(patients))
		response := make([]dto.PatientResponse, 0)

		PatientToResponse := func(c chan<- dto.PatientResponse, patient entities.Patient) {
			c <- *patient.ToPatientResponse()

		}

		for _, p := range patients {
			go PatientToResponse(c, p)
		}

		for range patients {
			response = append(response, <-c)
		}
		close(c)
	*/

	// 2 canales
	/*
		patients, err := s.db.FindAllPatients()
		if err != nil {
			return nil, err
		}
		rows := len(patients)

		fmt.Printf("Pascientes: %d \n", rows)

		jobsChan := make(chan entities.Patient, rows)
		completedJobsChan := make(chan dto.PatientResponse, rows)

		response := make([]dto.PatientResponse, 0)

		worker := func(jobsChan <-chan entities.Patient, completedJobsChan chan<- dto.PatientResponse) {
			for j := range jobsChan {
				completedJobsChan <- *j.ToPatientResponse()
			}
		}

		for w := 1; w <= 144; w++ {
			go worker(jobsChan, completedJobsChan)
		}

		for i := 0; i < rows; i++ {
			jobsChan <- patients[i]
		}
		close(jobsChan)

		for i := 0; i < rows; i++ {
			response = append(response, <-completedJobsChan)
		}

	*/

	// WaitGroup 1 : No se completan todos los envios
	/*
		start := time.Now()
		var wg sync.WaitGroup
		patients, err := s.db.FindAllPatients()
		if err != nil {
			return nil, err
		}
		rows := len(patients)
		response := make([]dto.PatientResponse, 0)

		wg.Add(rows)
		PatientToResponse := func(res *[]dto.PatientResponse, patient entities.Patient) {
			*res = append(*res, *patient.ToPatientResponse())
			for i := 0; i < 50000; i++ {
			}
			wg.Done()
		}

		for _, p := range patients {
			go PatientToResponse(&response, p)
		}
		wg.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Processes took %s \n", elapsed)
		fmt.Println(len(response))

	*/

	// WaitGroup + Mutex
	/*
		start := time.Now()
		var n sync.WaitGroup
		var mutex = sync.Mutex{}
		patients, err := s.db.FindAllPatients()
		if err != nil {
			return nil, err
		}
		fmt.Printf("Pascientes: %d \n", len(patients))
		response := make([]dto.PatientResponse, 0)

		for _, patient := range patients {
			n.Add(1)
			go func(patient entities.Patient) {
				mutex.Lock()
				response = append(response, *patient.ToPatientResponse())
				// Agregar Procesos Pesados
				for i := 0; i < 50000; i++ {
				}
				mutex.Unlock()
				n.Done()
			}(patient)
		}
		n.Wait()
		elapsed := time.Since(start)
		fmt.Printf("Processes took %s \n", elapsed)
		fmt.Println(len(response))
	*/

	// WaitGroup + Chanel

	start := time.Now()
	var n sync.WaitGroup
	//var mutex = sync.Mutex{}
	patients, err := s.db.FindAllPatients()
	c := make(chan dto.PatientResponse, len(patients))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Pascientes: %d \n", len(patients))
	response := make([]dto.PatientResponse, 0)

	for _, patient := range patients {
		n.Add(1)
		go func(patient entities.Patient) {
			c <- *patient.ToPatientResponse()
			// Agregar Procesos Pesados
			for i := 0; i < 50000; i++ {
			}
			n.Done()
		}(patient)
	}
	n.Wait()
	for len(c) > 0 {
		response = append(response, <-c)
	}
	close(c)
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s \n", elapsed)
	fmt.Println(len(response))

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
	/*
		if req.IsUnderAge() {
			return nil, errors.New("patient under 18 age")
		}
	*/

	patientSave, err := s.db.SavePatient(req)
	if err != nil {
		return nil, errors.New("can't Save Patient")
	}
	response := *patientSave.ToPatientMessage(patientSave.PatientId, "patient saved")
	return &response, nil
}

func (s DefaultPacientService) UpdatePatient(id int64, req dto.PatientRequest) (*dto.PatientMessage, error) {
	/*
		if req.IsUnderAge() {
			return nil, errors.New("patient under 18 age")
		}

	*/
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
