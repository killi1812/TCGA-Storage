package parser

type PatientDataWhole struct {
	BCRPatientBarcode           string `json:"bcr_patient_barcode"`
	Type                        string `json:"type"`
	AgeAtDiagnosis              int    `json:"age_at_initial_pathologic_diagnosis"`
	Gender                      string `json:"gender"`
	Race                        string `json:"race"`
	Stage                       string `json:"ajcc_pathologic_tumor_stage"`
	ClinicalStage               string `json:"clinical_stage"`
	HistologicalType            string `json:"histological_type"`
	HistologicalGrade           string `json:"histological_grade"`
	InitialDiagnosisYear        int    `json:"initial_pathologic_dx_year"`
	MenopauseStatus             string `json:"menopause_status"`
	BirthDaysTo                 int    `json:"birth_days_to"`
	VitalStatus                 string `json:"vital_status"`
	TumorStatus                 string `json:"tumor_status"`
	LastContactDaysTo           int    `json:"last_contact_days_to"`
	DeathDaysTo                 int    `json:"death_days_to"`
	CauseOfDeath                string `json:"cause_of_death"`
	NewTumorEventType           string `json:"new_tumor_event_type"`
	NewTumorEventSite           string `json:"new_tumor_event_site"`
	NewTumorEventSiteOther      string `json:"new_tumor_event_site_other"`
	NewTumorEventDxDaysTo       int    `json:"new_tumor_event_dx_days_to"`
	TreatmentOutcomeFirstCourse string `json:"treatment_outcome_first_course"`
	MarginStatus                string `json:"margin_status"`
	ResidualTumor               string `json:"residual_tumor"`
	OS                          int    `json:"os"`
	OSTime                      int    `json:"os_time"`
	DSS                         int    `json:"dss"`
	DSSTime                     int    `json:"dss_time"`
	DFI                         int    `json:"dfi"`
	DFITime                     int    `json:"dfi_time"`
	PFI                         int    `json:"pfi"`
	PFITime                     int    `json:"pfi_time"`
	Redaction                   string `json:"redaction"`
}

type PatientData struct {
	BCRPatientBarcode string `json:"bcr_patient_barcode"`
	DSS               bool   `json:"dss"`
	OS                bool   `json:"os"`
	ClinicalStage     string `json:"clinical_stage"`
}
