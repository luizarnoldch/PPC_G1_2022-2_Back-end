package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Area struct {
	AreaID       int64  `db:"ID_Area"`
	AreaName     string `db:"Nombre_Area"`
	AreaAddress  string `db:"Ubicacion_Area"`
	AreaLocation string `db:"Localidades_Area"`
	AreaCity     string `db:"Ciudad_Area"`
	AreaCountry  string `db:"Pais_Area"`
	AreaCapacity int64  `db:"Capacidad_Area"`
	AreaTag      string `db:"Tag_Area"`
	AreaDiscount string `db:"Desc_Area"`
}

func (a Area) ToAreaResponse() *dto.AreaResponse {
	return &dto.AreaResponse{
		a.AreaID,
		a.AreaName,
		a.AreaAddress,
		a.AreaLocation,
		a.AreaCity,
		a.AreaCountry,
		a.AreaCapacity,
		a.AreaTag,
		a.AreaDiscount}
}

func (a Area) ToAreaMessage(id int64, message string) *dto.AreaMessage {
	return &dto.AreaMessage{id, message}
}

func NewArea(AreaID int64,
	AreaName string,
	AreaAddress string,
	AreaLocation string,
	AreaCity string,
	AreaCountry string,
	AreaCapacity int64,
	AreaTag string,
	AreaDiscount string) Area {
	return Area{
		AreaID,
		AreaName,
		AreaAddress,
		AreaLocation,
		AreaCity,
		AreaCountry,
		AreaCapacity,
		AreaTag,
		AreaDiscount}
}
