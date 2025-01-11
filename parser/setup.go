package parser

import "sync"

type TsvParser struct {
}

type TxtParser struct {
}

var tsvParserInstance *TsvParser = nil
var txtParserInstance *TxtParser = nil

func GetTsvParser() *TsvParser {
	sync.OnceFunc(func() {
		tsvParserInstance = &TsvParser{}
	})
	return tsvParserInstance
}

func GetTxtParser() *TxtParser {
	sync.OnceFunc(func() {
		txtParserInstance = &TxtParser{}
	})
	return txtParserInstance
}
