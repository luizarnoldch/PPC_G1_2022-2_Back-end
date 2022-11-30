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

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, error)
	GetUser(id int64) (*dto.UserResponse, error)
	PostUser(request dto.UserRequest) (*dto.UserMessage, error)
	UpdateUser(id int64, request dto.UserRequest) (*dto.UserMessage, error)
	DeleteUser(id int64) (*dto.UserMessage, error)
}

type DefaultUserService struct {
	db repositories.UserRepository
}

func NewUserService(db repositories.UserRepository) DefaultUserService {
	return DefaultUserService{db}
}

func (s DefaultUserService) GetAllUsers() ([]dto.UserResponse, error) {
	// WaitGroup + Chanel
	start := time.Now()
	var wg sync.WaitGroup

	users, err := s.db.FindAllUser()
	c := make(chan dto.UserResponse, len(users))
	if err != nil {
		return nil, err
	}
	fmt.Printf("users: %d \n", len(users))
	response := make([]dto.UserResponse, 0)

	for _, user := range users {
		wg.Add(1)
		go func(user entities.User) {
			c <- *user.ToUserResponse()
			for i := 0; i < 50000; i++ {
			}
		}(user)
		wg.Done()
	}
	wg.Wait()
	for range users {
		response = append(response, <-c)
	}
	close(c)
	elapsed := time.Since(start)
	fmt.Printf("process took %s \n", elapsed)
	fmt.Println(len(response))
	return response, nil
}

func (s DefaultUserService) GetUser(id int64) (*dto.UserResponse, error) {
	user, err := s.db.FinUserById(id)
	if err != nil {
		return nil, err
	}
	response := *user.ToUserResponse()
	return &response, nil
}

func (s DefaultUserService) PostUser(request dto.UserRequest) (*dto.UserMessage, error) {
	userSave, err := s.db.SaveUser(request)
	if err != nil {
		return nil, errors.New("can't save user")
	}
	response := *userSave.ToUserMessage(userSave.ID_Usuario, "user saved")
	return &response, nil
}

func (s DefaultUserService) UpdateUser(id int64, request dto.UserRequest) (*dto.UserMessage, error) {
	userUpated, err := s.db.UpdateUser(id, request)
	if err != nil {
		return nil, errors.New("can't updated user")
	}
	response := *userUpated.ToUserMessage(id, "user updated")
	return &response, nil
}

func (s DefaultUserService) DeleteUser(id int64) (*dto.UserMessage, error) {
	user, err := s.db.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	response := *user.ToUserMessage(id, "area deleted")
	return &response, nil
}
