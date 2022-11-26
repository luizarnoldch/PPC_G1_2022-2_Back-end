package dto

type AreaRequest struct {
	AreaID       int64  `json:"ID_Area" form:"ID_Area"`
	AreaName     string `json:"Nombre_Area" form:"Nombre_Area"`
	AreaAddress  string `json:"Ubicacion_Area" form:"Ubicacion_Area"`
	AreaLocation string `json:"Localidades_Area" form:"Localidades_Area"`
	AreaCity     string `json:"Ciudad_Area" form:"Ciudad_Area"`
	AreaCountry  string `json:"Pais_Area" form:"Pais_Area"`
	AreaCapacity int64  `json:"Capacidad_Area" form:"Capacidad_Area"`
	AreaTag      string `json:"Tag_Area" form:"Tag_Area"`
	AreaDiscount string `json:"Desc_Area" form:"Desc_Area"`
}

/*
func (p PatientRequest) IsUnderAge() bool {
	return p.PatientAge < 18
}
*/
