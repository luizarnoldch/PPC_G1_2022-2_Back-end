package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type User struct {
	ID_Usuario           int64  `db:"ID_Usuario"`
	ID_Area              int64  `db:"ID_Area"`
	ID_Perfil            int64  `db:"ID_Perfil"`
	Nombre_Usuario       string `db:"Nombre_Usuario"`
	Apellido_Usuairo     string `db:"Apellido_Usuairo"`
	Nick_Usuario         string `db:"Nick_Usuario"`
	Clave_Usuario        string `db:"Clave_Usuario"`
	Mail_Usuario         string `db:"Mail_Usuario"`
	Especialidad_Usuario string `db:"Especialidad_Usuario"`
	Estado_Usuario       int64  `db:"Estado_Usuario"`
}

func (u User) ToUserResponse() *dto.UserResponse {
	return &dto.UserResponse{
		u.ID_Usuario,
		u.ID_Area,
		u.ID_Perfil,
		u.Nombre_Usuario,
		u.Apellido_Usuairo,
		u.Nick_Usuario,
		u.Clave_Usuario,
		u.Mail_Usuario,
		u.Especialidad_Usuario,
		u.Estado_Usuario}
}

func (u User) ToUserMessage(id int64, message string) *dto.UserMessage {
	return &dto.UserMessage{id, message}
}

func NewUser(ID_Usuario int64,
	ID_Area int64,
	ID_Perfil int64,
	Nombre_Usuario string,
	Apellido_Usuairo string,
	Nick_Usuario string,
	Clave_Usuario string,
	Mail_Usuario string,
	Especialidad_Usuario string,
	Estado_Usuario int64) User {
	return User{ID_Usuario,
		ID_Area,
		ID_Perfil,
		Nombre_Usuario,
		Apellido_Usuairo,
		Nick_Usuario,
		Clave_Usuario,
		Mail_Usuario,
		Especialidad_Usuario,
		Estado_Usuario}
}
