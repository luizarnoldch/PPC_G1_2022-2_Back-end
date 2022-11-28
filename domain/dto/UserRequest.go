package dto

type UserRequest struct {
	ID_Usuario           int64  `form:"ID_Usuario" json:"ID_Usuario"`
	ID_Area              int64  `form:"ID_Area" json:"ID_Area"`
	ID_Perfil            int64  `form:"ID_Perfil" json:"ID_Perfil"`
	Nombre_Usuario       string `form:"Nombre_Usuario" json:"Nombre_Usuario"`
	Apellido_Usuairo     string `form:"Apellido_Usuairo" json:"Apellido_Usuairo"`
	Nick_Usuario         string `form:"Nick_Usuario" json:"Nick_Usuario"`
	Clave_Usuario        string `form:"Clave_Usuario" json:"Clave_Usuario"`
	Mail_Usuario         string `form:"Mail_Usuario" json:"Mail_Usuario"`
	Especialidad_Usuario string `form:"Especialidad_Usuario" json:"Especialidad_Usuario"`
	Estado_Usuario       int64  `form:"Estado_Usuario" json:"Estado_Usuario"`
}
