package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Profile struct {
	ID_Perfil        int64  `db:"ID_Perfil"`
	Nombre_Perfil    string `db:"Nombre_Perfil"`
	Estado_Perfil    int64  `db:"Estado_Perfil"`
	Atributos_Perfil string `db:"Atributos_Perfil"`
}

func (p Profile) ToProfileResponse() *dto.ProfileResponse {
	return &dto.ProfileResponse{
		p.ID_Perfil,
		p.Nombre_Perfil,
		p.Estado_Perfil,
		p.Atributos_Perfil}
}

func (a Profile) ToProfileMessage(id int64, message string) *dto.ProfileMessage {
	return &dto.ProfileMessage{id, message}
}

func NewProfile(
	ID_Perfil int64,
	Nombre_Perfil string,
	Estado_Perfil int64,
	Atributos_Perfil string) Profile {
	return Profile{
		ID_Perfil,
		Nombre_Perfil,
		Estado_Perfil,
		Atributos_Perfil}
}
