package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Patient struct {
	PatientId          int64  `db:"ID_Paciente"`
	PatientName        string `db:"Nombre_Paciente"`
	PatientLastName    string `db:"Apellido_Paciente"`
	PatientNick        string `db:"Nick_Paciente"`
	PatientKey         string `db:"Clave_Paciente"`
	PatientPhoto       string `db:"Foto_Paciente"`
	PatientNationality string `db:"Nacionalidad_Paciente"`
	PatientBirth       string `db:"DATE_Nac_Paciente"`
	PatientEmail       string `db:"Email_Paciente"`
	PatientTelf        string `db:"Telefono_Paciente"`
	PatientPhone       string `db:"Celular_Paciente"`
	PatientCard        string `db:"Cedula_Paciente"`
	PatientDiscount    string `db:"Desc_Paciente"`
	PatientFile        string `db:"Archivo_Paciente"`
	PatientState       string `db:"Estado_Paciente"`
	PatientTag         string `db:"Tag_Paciente"`
}

func (p Patient) ToPatientResponse() *dto.PatientResponse {
	return &dto.PatientResponse{
		p.PatientId,
		p.PatientName,
		p.PatientLastName,
		p.PatientNick,
		p.PatientPhoto,
		p.PatientKey,
		p.PatientNationality,
		p.PatientBirth,
		p.PatientEmail,
		p.PatientTelf,
		p.PatientPhone,
		p.PatientCard,
		p.PatientDiscount,
		p.PatientFile,
		p.PatientState,
		p.PatientTag,
	}
}

func (p Patient) ToPatientMessage(id int64, message string) *dto.PatientMessage {
	return &dto.PatientMessage{PatientId: id, PatientMessage: message}
}

func NewPatient(
	PatientId int64,
	PatientName string,
	PatientLastName string,
	PatientNick string,
	PatientKey string,
	PatientPhoto string,
	PatientNationality string,
	PatientBirth string,
	PatientEmail string,
	PatientTelf string,
	PatientPhone string,
	PatientCard string,
	PatientDiscount string,
	PatientFile string,
	PatientState string,
	PatientTag string,
) Patient {
	return Patient{
		PatientId,
		PatientName,
		PatientLastName,
		PatientNick,
		PatientKey,
		PatientPhoto,
		PatientNationality,
		PatientBirth,
		PatientEmail,
		PatientTelf,
		PatientPhone,
		PatientCard,
		PatientDiscount,
		PatientFile,
		PatientState,
		PatientTag,
	}
}
