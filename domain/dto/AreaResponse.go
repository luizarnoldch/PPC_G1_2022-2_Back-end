package dto

type AreaResponse struct {
	AreaID       int64  `json:"ID_Area"`
	AreaName     string `json:"Nombre_Area"`
	AreaAddress  string `json:"Ubicacion_Area"`
	AreaLocation string `json:"Localidades_Area"`
	AreaCity     string `json:"Ciudad_Area"`
	AreaCountry  string `json:"Pais_Area"`
	AreaCapacity int64  `json:"Capacidad_Area"`
	AreaTag      string `json:"Tag_Area"`
	AreaDiscount string `json:"Desc_Area"`
}
