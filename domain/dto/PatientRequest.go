package dto

type PatientRequest struct {
	PatientId          int64  `json:"ID_Paciente" form:"ID_Paciente"`
	PatientName        string `json:"Nombre_Paciente" form:"Nombre_Paciente"`
	PatientLastName    string `json:"Apellido_Paciente" form:"Apellido_Paciente"`
	PatientNick        string `json:"Nick_Paciente" form:"Nick_Paciente"`
	PatientKey         string `json:"Clave_Paciente" form:"Clave_Paciente"`
	PatientPhoto       string `json:"Foto_Paciente" form:"Foto_Paciente"`
	PatientNationality string `json:"Nacionalidad_Paciente" form:"Nacionalidad_Paciente"`
	PatientBirth       string `json:"DATE_Nac_Paciente" form:"DATE_Nac_Paciente"`
	PatientEmail       string `json:"Email_Paciente" form:"Email_Paciente"`
	PatientTelf        string `json:"Telefono_Paciente" form:"Telefono_Paciente"`
	PatientPhone       string `json:"Celular_Paciente" form:"Celular_Paciente"`
	PatientCard        string `json:"Cedula_Paciente" form:"Cedula_Paciente"`
	PatientDiscount    string `json:"Desc_Paciente" form:"Desc_Paciente"`
	PatientFile        string `json:"Archivo_Paciente" form:"Archivo_Paciente"`
	PatientState       string `json:"Estado_Paciente" form:"Estado_Paciente"`
	PatientTag         string `json:"Tag_Paciente" form:"Tag_Paciente"`
}

/*
func (p PatientRequest) IsUnderAge() bool {
	return p.PatientAge < 18
}
*/
