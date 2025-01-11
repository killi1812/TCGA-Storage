package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

func (this *TsvParser) Parse(reader io.Reader) ([]PatientData, error) {
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

func (this *TxtParser) Parse(reader io.Reader) ([]PatientGenesExpressions, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '\t'

	// Read the header
	header, err := csvReader.Read()
	patients := header[1:]
	data := make([]PatientGenesExpressions, len(patients))

	for i, p := range patients {
		data[i].BCRPatientBarcode = p
		data[i].Genes = make([]GeneExpressionPair, 20530)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}
	row := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		for i, val := range record[1:] {
			data[i].Genes[row] = GeneExpressionPair{
				Gene:       record[0],
				Expression: stof(val),
			}
		}
		row++
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
