package dto

type PatientRequest struct {
	PatientName     string `json:"patient_name" form:"patient_name"`
	PatientLastName string `json:"patient_last_name" form:"patient_last_name"`
	PatientAge      int64  `json:"patient_age" form:"patient_age"`
}

func (p PatientRequest) IsUnderAge() bool {
	return p.PatientAge < 18
}
