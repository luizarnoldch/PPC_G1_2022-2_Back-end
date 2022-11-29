package dto

type ProfileResponse struct {
	ID_Perfil        int64  `json:"ID_Perfil"`
	Nombre_Perfil    string `json:"Nombre_Perfil"`
	Estado_Perfil    int64  `json:"Estado_Perfil"`
	Atributos_Perfil string `json:"Atributos_Perfil"`
}
