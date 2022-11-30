package entities

import "github.com/luizarnoldch/PPC_G1_2022-2_Back-end/domain/dto"

type Citation struct {
	ID_Cita                    int64  `db:"ID_Cita"`
	ID_Usuario                 int64  `db:"ID_Usuario"`
	ID_Paciente                int64  `db:"ID_Paciente"`
	Titulo_Cita                string `db:"Titulo_Cita"`
	DATE_Inicio_Cita           string `db:"DATE_Inicio_Cita"`
	DATE_Fin_Cita              string `db:"DATE_Fin_Cita"`
	Hora_Inicio_Cita           string `db:"Hora_Inicio_Cita"`
	Hora_Fin_Cita              string `db:"Hora_Fin_Cita"`
	Descripcion_Cita           string `db:"Descripcion_Cita"`
	Identificador_Cita         string `db:"Identificador_Cita"`
	Hora_Llegada_Paciente_Cita string `db:"Hora_Llegada_Paciente_Cita"`
	Hora_Salida_Paciente_cita  string `db:"Hora_Salida_Paciente_Cita"`
	Motivo_Anulacion_Cita      string `db:"Motivo_Anulacion_Cita"`
	Estado_Cita                string `db:"Estado_Cita"`
}

func (c Citation) ToCitationResponse() *dto.CitationResponse {
	return &dto.CitationResponse{
		c.ID_Cita,
		c.ID_Usuario,
		c.ID_Paciente,
		c.Titulo_Cita,
		c.DATE_Inicio_Cita,
		c.DATE_Fin_Cita,
		c.Hora_Inicio_Cita,
		c.Hora_Fin_Cita,
		c.Descripcion_Cita,
		c.Identificador_Cita,
		c.Hora_Llegada_Paciente_Cita,
		c.Hora_Salida_Paciente_cita,
		c.Motivo_Anulacion_Cita,
		c.Estado_Cita}
}

func (c Citation) ToCitationMessage(id int64, message string) *dto.CitationMessage {
	return &dto.CitationMessage{id, message}
}

func NewCitation(ID_Cita int64,
	ID_Usuario int64,
	ID_Paciente int64,
	Titulo_Cita string,
	DATE_Inicio_Cita string,
	DATE_Fin_Cita string,
	Hora_Inicio_Cita string,
	Hora_Fin_Cita string,
	Descripcion_Cita string,
	Identificador_Cita string,
	Hora_Llegada_Paciente_Cita string,
	Hora_Salida_Paciente_cita string,
	Motivo_Anulacion_Cita string,
	Estado_Cita string) Citation {
	return Citation{
		ID_Cita,
		ID_Usuario,
		ID_Paciente,
		Titulo_Cita,
		DATE_Inicio_Cita,
		DATE_Fin_Cita,
		Hora_Inicio_Cita,
		Hora_Fin_Cita,
		Descripcion_Cita,
		Identificador_Cita,
		Hora_Llegada_Paciente_Cita,
		Hora_Salida_Paciente_cita,
		Motivo_Anulacion_Cita,
		Estado_Cita}
}
