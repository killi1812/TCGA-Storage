package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strconv"
)

func (this *PatientParser) Parse(reader io.Reader) ([]PatientData, error) {
	var result []PatientData

	csvReader := csv.NewReader(reader)
	csvReader.Comma = '\t'

	// Read the header
	_, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	// Read rows and parse into struct
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		// Parse the fields into the struct
		data := PatientData{
			BCRPatientBarcode: record[1],
			ClinicalStage:     record[7],
			OS:                stob(record[25]),
			DSS:               stob(record[27]),
		}
		result = append(result, data)
	}

	return result, nil
}

var genesFilter = []string{"C6orf150", "CCL5", "CXCL10", "TMEM173", "CXCL9", "CXCL11", "NFKB1", "IKBKE", "IRF3", "TREX1", "ATM", "IL6", "IL8"}

func (this *GeneParser) Parse(reader io.ReadCloser, patientCode string) (PatientGenesExpressions, error) {
	defer reader.Close()
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '\t'

	// Read the header
	header, err := csvReader.Read()
	patients := header[1:]
	index := slices.Index(patients, patientCode)
	if index == -1 {
		return PatientGenesExpressions{}, PatientNotFound
	}

	data := PatientGenesExpressions{
		BCRPatientBarcode: patientCode,
		Genes:             make([]GeneExpressionPair, len(genesFilter)),
	}

	if err != nil {
		return PatientGenesExpressions{}, fmt.Errorf("failed to read header: %w", err)
	}
	i := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return PatientGenesExpressions{}, fmt.Errorf("failed to read row: %w", err)
		}

		if slices.Contains(genesFilter, record[0]) {
			data.Genes[i] = GeneExpressionPair{
				Gene:       record[0],
				Expression: stof(record[index+1])}
			i++
		}
	}
	return data, nil
}

func stof(str string) float64 {
	br, _ := strconv.ParseFloat(str, 64)
	return br
}

func stob(str string) bool {
	br, _ := strconv.Atoi(str)
	return br == 1
}
