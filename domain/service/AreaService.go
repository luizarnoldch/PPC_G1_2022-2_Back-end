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

type AreaService interface {
	GetAllAreas() ([]dto.AreaResponse, error)
	GetArea(id int64) (*dto.AreaResponse, error)
	PostArea(request dto.AreaRequest) (*dto.AreaMessage, error)
	UpdateArea(id int64, request dto.AreaRequest) (*dto.AreaMessage, error)
	DeleteArea(id int64) (*dto.AreaMessage, error)
}

type DefaultAreaService struct {
	db repositories.AreaRepository
}

func NewAreaService(db repositories.AreaRepository) DefaultAreaService {
	return DefaultAreaService{db}
}

func (s DefaultAreaService) GetAllAreas() ([]dto.AreaResponse, error) {
	// WaitGroup + Chanel
	start := time.Now()
	var wg sync.WaitGroup

	areas, err := s.db.FindAllAreas()
	c := make(chan dto.AreaResponse, len(areas))
	if err != nil {
		return nil, err
	}
	fmt.Printf("areas: %d \n", len(areas))
	response := make([]dto.AreaResponse, 0)

	for _, area := range areas {
		wg.Add(1)
		go func(area entities.Area) {
			c <- *area.ToAreaResponse()
			for i := 0; i < 50000; i++ {

			}
		}(area)
		wg.Done()
	}
	wg.Wait()
	for range areas {
		response = append(response, <-c)
	}
	close(c)
	elapsed := time.Since(start)
	fmt.Printf("process took %s \n", elapsed)
	fmt.Println(len(response))
	return response, nil
}

func (s DefaultAreaService) GetArea(id int64) (*dto.AreaResponse, error) {
	area, err := s.db.FinAreaById(id)
	if err != nil {
		return nil, err
	}
	response := *area.ToAreaResponse()
	return &response, nil
}

func (s DefaultAreaService) PostArea(request dto.AreaRequest) (*dto.AreaMessage, error) {
	areaSave, err := s.db.SaveArea(request)
	if err != nil {
		return nil, errors.New("can't save area")
	}
	response := *areaSave.ToAreaMessage(areaSave.AreaID, "area saved")
	return &response, nil
}

func (s DefaultAreaService) UpdateArea(id int64, request dto.AreaRequest) (*dto.AreaMessage, error) {
	areaUpated, err := s.db.UpdateArea(id, request)
	if err != nil {
		return nil, errors.New("can't updated area")
	}
	response := *areaUpated.ToAreaMessage(id, "area updated")
	return &response, nil
}

func (s DefaultAreaService) DeleteArea(id int64) (*dto.AreaMessage, error) {
	area, err := s.db.DeleteArea(id)
	if err != nil {
		return nil, err
	}
	response := *area.ToAreaMessage(id, "area deleted")
	return &response, nil
}
