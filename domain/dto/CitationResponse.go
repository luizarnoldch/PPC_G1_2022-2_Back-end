package dto

type CitationResponse struct {
	ID_Cita                    int64  `json:"ID_Cita" `
	ID_Usuario                 int64  `json:"ID_Usuario"`
	ID_Paciente                int64  `json:"ID_Paciente"`
	Titulo_Cita                string `json:"Titulo_Cita"`
	DATE_Inicio_Cita           string `json:"DATE_Inicio_Cita"`
	DATE_Fin_Cita              string `json:"DATE_Fin_Cita"`
	Hora_Inicio_Cita           string `json:"Hora_Inicio_Cita"`
	Hora_Fin_Cita              string `json:"Hora_Fin_Cita"`
	Descripcion_Cita           string `json:"Descripcion_Cita"`
	Identificador_Cita         string `json:"Identificador_Cita"`
	Hora_Llegada_Paciente_Cita string `json:"Hora_Llegada_Paciente_Cita"`
	Hora_Salida_Paciente_cita  string `json:"Hora_Salida_Paciente_Cita"`
	Motivo_Anulacion_Cita      string `json:"Motivo_Anulacion_Cita"`
	Estado_Cita                string `json:"Estado_Cita"`
}
