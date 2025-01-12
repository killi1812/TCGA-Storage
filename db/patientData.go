package db

type PatientData struct {
	BCRPatientBarcode string `json:"bcr_patient_barcode" bson:"bcr_patient_barcode,omitempty"`
	DSS               bool   `json:"dss" bson:"dss,omitempty"`
	OS                bool   `json:"os" bson:"os,omitempty"`
	ClinicalStage     string `json:"clinical_stage" bson:"clinical_stage,omitempty"`
}
