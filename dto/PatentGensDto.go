package dto

import (
	"TCGA-storage/parser"
)

type PatientGensDto struct {
	BCRPatientBarcode string              `json:"bcr_patient_barcode"`
	DSS               bool                `json:"dss"`
	OS                bool                `json:"os"`
	ClinicalStage     string              `json:"clinical_stage"`
	Genes             []GeneExpressionDto `json:"genes"`
}

type GeneExpressionDto struct {
	Gene       string  `json:"gene"`
	Expression float64 `json:"expression"`
}

func NewPatientGensDto(patient parser.PatientData, data parser.PatientGenesExpressions) PatientGensDto {
	return PatientGensDto{
		BCRPatientBarcode: patient.BCRPatientBarcode,
		DSS:               patient.DSS,
		OS:                patient.OS,
		ClinicalStage:     patient.ClinicalStage,
		Genes:             NewGeneExpressionDto(data.Genes),
	}
}

func NewGeneExpressionDto(data []parser.GeneExpressionPair) []GeneExpressionDto {
	genes := make([]GeneExpressionDto, len(data))
	for i := 0; i < len(data); i++ {
		genes[i] = GeneExpressionDto{
			Gene:       data[i].Gene,
			Expression: data[i].Expression,
		}
	}

	return genes
}
