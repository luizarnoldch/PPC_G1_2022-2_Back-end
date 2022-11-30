package service

import (
	"errors"
	"fmt"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
	"github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/repositories"
	// "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/entities"
	"sync"
	"time"
)

type CitationService interface {
	GetAllCitations() ([]dto.CitationResponse, error)
	GetCitation(id int64) (*dto.CitationResponse, error)
	PostCitation(request dto.CitationRequest) (*dto.CitationMessage, error)
	UpdateCitation(id int64, request dto.CitationRequest) (*dto.CitationMessage, error)
	DeleteCitation(id int64) (*dto.CitationMessage, error)
}

type DefaultCitationService struct {
	db repositories.CitationRepository
}

func NewCitationService(db repositories.CitationRepository) DefaultCitationService {
	return DefaultCitationService{db}
}

func (s DefaultCitationService) GetAllCitations() ([]dto.CitationResponse, error) {
	// WaitGroup + Chanel
	start := time.Now()
	var wg sync.WaitGroup

	citations, err := s.db.FindAllCitation()
	c := make(chan dto.CitationResponse, len(citations))
	if err != nil {
		return nil, err
	}
	fmt.Printf("citations: %d \n", len(citations))
	response := make([]dto.CitationResponse, 0)

	for _, citation := range citations {
		wg.Add(1)
		go func(citation entities.Citation) {
			c <- *citation.ToCitationResponse()
			for i := 0; i < 50000; i++ {

			}
		}(citation)
		wg.Done()
	}
	wg.Wait()
	for range citations {
		response = append(response, <-c)
	}
	close(c)
	elapsed := time.Since(start)
	fmt.Printf("process took %s \n", elapsed)
	fmt.Println(len(response))
	return response, nil
}

func (s DefaultCitationService) GetCitation(id int64) (*dto.CitationResponse, error) {
	citation, err := s.db.FinCitationById(id)
	if err != nil {
		return nil, err
	}
	response := *citation.ToCitationResponse()
	return &response, nil
}

func (s DefaultCitationService) PostCitation(request dto.CitationRequest) (*dto.CitationMessage, error) {
	citationSave, err := s.db.SaveCitation(request)
	if err != nil {
		return nil, errors.New("can't save area")
	}
	response := *citationSave.ToCitationMessage(citationSave.ID_Cita, "citation saved")
	return &response, nil
}

func (s DefaultCitationService) UpdateCitation(id int64, request dto.CitationRequest) (*dto.CitationMessage, error) {
	citationUpated, err := s.db.UpdateCitation(id, request)
	if err != nil {
		return nil, errors.New("can't updated area")
	}
	response := *citationUpated.ToCitationMessage(id, "citation updated")
	return &response, nil
}

func (s DefaultCitationService) DeleteCitation(id int64) (*dto.CitationMessage, error) {
	citation, err := s.db.DeleteCitation(id)
	if err != nil {
		return nil, err
	}
	response := *citation.ToCitationMessage(id, "citation deleted")
	return &response, nil
}
