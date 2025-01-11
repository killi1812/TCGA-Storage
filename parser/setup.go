package parser

import "sync"

type PatientParser struct {
}

type GeneParser struct {
}

var tsvParserInstance *PatientParser = nil
var txtParserInstance *GeneParser = nil

func GetPatientParser() *PatientParser {
	sync.OnceFunc(func() {
		tsvParserInstance = &PatientParser{}
	})
	return tsvParserInstance
}

func GetGeneParser() *GeneParser {
	sync.OnceFunc(func() {
		txtParserInstance = &GeneParser{}
	})
	return txtParserInstance
}
