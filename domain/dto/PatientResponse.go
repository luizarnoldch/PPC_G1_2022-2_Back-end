package dto

type PatientResponse struct {
	PatientId       int64  `json:"patient_id"`
	PatientName     string `json:"patient_name"`
	PatientLastName string `json:"patient_last_name"`
	PatientAge      int64  `json:"patient_age"`
}
