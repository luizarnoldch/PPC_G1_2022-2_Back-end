package dto

type ProfileRequest struct {
	ID_Perfil        int64  `json:"ID_Perfil" form:"ID_Perfil"`
	Nombre_Perfil    string `json:"Nombre_Perfil" form:"Nombre_Perfil"`
	Estado_Perfil    int64  `json:"Estado_Perfil" form:"Estado_Perfil"`
	Atributos_Perfil string `json:"Atributos_Perfil" form:"Atributos_Perfil"`
}
