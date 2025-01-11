package parser

type PatientGenesExpressions struct {
	BCRPatientBarcode string               `json:"bcr_patient_barcode"`
	Genes             []GeneExpressionPair `json:"genes"`
}

type GeneExpressionPair struct {
	Gene       string  `json:"gene"`
	Expression float64 `json:"expression"`
}
