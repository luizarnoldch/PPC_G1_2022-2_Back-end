package dto

type UserResponse struct {
	ID_Usuario           int64  `json:"ID_Usuario"`
	ID_Area              int64  `json:"ID_Area"`
	ID_Perfil            int64  `json:"ID_Perfil"`
	Nombre_Usuario       string `json:"Nombre_Usuario"`
	Apellido_Usuairo     string `json:"Apellido_Usuairo"`
	Nick_Usuario         string `json:"Nick_Usuario"`
	Clave_Usuario        string `json:"Clave_Usuario"`
	Mail_Usuario         string `json:"Mail_Usuario"`
	Especialidad_Usuario string `json:"Especialidad_Usuario"`
	Estado_Usuario       int64  `json:"Estado_Usuario"`
}
