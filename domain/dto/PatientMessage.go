package dto

type PatientMessage struct {
	PatientId      int64  `json:"patient_id"`
	PatientMessage string `json:"patient_message"`
}
