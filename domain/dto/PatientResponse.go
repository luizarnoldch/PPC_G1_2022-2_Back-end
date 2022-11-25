package dto

type PatientResponse struct {
	PatientId          int64  `json:"ID_Paciente"`
	PatientName        string `json:"Nombre_Paciente"`
	PatientLastName    string `json:"Apellido_Paciente"`
	PatientNick        string `json:"Nick_Paciente"`
	PatientKey         string `json:"Clave_Paciente"`
	PatientPhoto       string `json:"Foto_Paciente"`
	PatientNationality string `json:"Nacionalidad_Paciente"`
	PatientBirth       string `json:"DATE_Nac_Paciente"`
	PatientEmail       string `json:"Email_Paciente"`
	PatientTelf        string `json:"Telefono_Paciente"`
	PatientPhone       string `json:"Celular_Paciente"`
	PatientCard        string `json:"Cedula_Paciente"`
	PatientDiscount    string `json:"Desc_Paciente"`
	PatientFile        string `json:"Archivo_Paciente"`
	PatientState       string `json:"Estado_Paciente"`
	PatientTag         string `json:"Tag_Paciente"`
}
