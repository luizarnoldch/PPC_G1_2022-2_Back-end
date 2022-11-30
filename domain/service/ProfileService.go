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

type ProfileService interface {
	GetAllProfiles() ([]dto.ProfileResponse, error)
	GetProfile(id int64) (*dto.ProfileResponse, error)
	PostProfile(request dto.ProfileRequest) (*dto.ProfileMessage, error)
	UpdateProfile(id int64, request dto.ProfileRequest) (*dto.ProfileMessage, error)
	DeleteProfile(id int64) (*dto.ProfileMessage, error)
}

type DefaultProfileService struct {
	db repositories.ProfileRepository
}

func NewProfileService(db repositories.ProfileRepository) DefaultProfileService {
	return DefaultProfileService{db}
}

func (s DefaultProfileService) GetAllProfiles() ([]dto.ProfileResponse, error) {
	// WaitGroup + Chanel
	start := time.Now()
	var wg sync.WaitGroup

	profiles, err := s.db.FindAllProfiles()
	c := make(chan dto.ProfileResponse, len(profiles))
	if err != nil {
		return nil, err
	}
	fmt.Printf("profiles: %d \n", len(profiles))
	response := make([]dto.ProfileResponse, 0)

	for _, profile := range profiles {
		wg.Add(1)
		go func(area entities.Profile) {
			c <- *profile.ToProfileResponse()
			for i := 0; i < 50000; i++ {

			}
		}(profile)
		wg.Done()
	}
	wg.Wait()
	for range profiles {
		response = append(response, <-c)
	}
	close(c)
	elapsed := time.Since(start)
	fmt.Printf("process took %s \n", elapsed)
	fmt.Println(len(response))
	return response, nil
}

func (s DefaultProfileService) GetProfile(id int64) (*dto.ProfileResponse, error) {
	profile, err := s.db.FinProfileById(id)
	if err != nil {
		return nil, err
	}
	response := *profile.ToProfileResponse()
	return &response, nil
}

func (s DefaultProfileService) PostProfile(request dto.ProfileRequest) (*dto.ProfileMessage, error) {
	profileSave, err := s.db.SaveProfile(request)
	if err != nil {
		return nil, errors.New("can't save area")
	}
	response := *profileSave.ToProfileMessage(profileSave.ID_Perfil, "prifile saved")
	return &response, nil
}

func (s DefaultProfileService) UpdateProfile(id int64, request dto.ProfileRequest) (*dto.ProfileMessage, error) {
	profileUpated, err := s.db.UpdateProfile(id, request)
	if err != nil {
		return nil, errors.New("can't updated area")
	}
	response := *profileUpated.ToProfileMessage(id, "profile updated")
	return &response, nil
}

func (s DefaultProfileService) DeleteProfile(id int64) (*dto.ProfileMessage, error) {
	profile, err := s.db.DeleteProfile(id)
	if err != nil {
		return nil, err
	}
	response := *profile.ToProfileMessage(id, "profile deleted")
	return &response, nil
}
