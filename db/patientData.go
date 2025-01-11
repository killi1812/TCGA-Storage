package db

type PatientData struct {
	BCRPatientBarcode string `json:"bcr_patient_barcode"`
	DSS               int    `json:"dss"`
	OS                int    `json:"os"`
	ClinicalStage     string `json:"clinical_stage"`
}
