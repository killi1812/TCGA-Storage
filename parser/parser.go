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
			Type:                        record[1],
			AgeAtDiagnosis:              intConv(record[2]),
			Gender:                      record[3],
			Race:                        record[4],
			Stage:                       record[5],
			ClinicalStage:               record[6],
			HistologicalType:            record[7],
			HistologicalGrade:           record[8],
			InitialDiagnosisYear:        intConv(record[9]),
			MenopauseStatus:             record[10],
			BirthDaysTo:                 intConv(record[11]),
			VitalStatus:                 record[12],
			TumorStatus:                 record[13],
			LastContactDaysTo:           intConv(record[14]),
			DeathDaysTo:                 intConv(record[15]),
			CauseOfDeath:                record[16],
			NewTumorEventType:           record[17],
			NewTumorEventSite:           record[18],
			NewTumorEventSiteOther:      record[19],
			NewTumorEventDxDaysTo:       intConv(record[20]),
			TreatmentOutcomeFirstCourse: record[21],
			MarginStatus:                record[22],
			ResidualTumor:               record[23],
			OS:                          intConv(record[24]),
			OSTime:                      intConv(record[25]),
			DSS:                         intConv(record[26]),
			DSSTime:                     intConv(record[27]),
			DFI:                         intConv(record[28]),
			DFITime:                     intConv(record[29]),
			PFI:                         intConv(record[30]),
			PFITime:                     intConv(record[31]),
			Redaction:                   record[32],
		}
		result = append(result, data)
	}

	return result, nil
}

func intConv(str string) int {
	br, _ := strconv.Atoi(str)
	return br
}
