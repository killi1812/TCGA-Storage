package parser

import "sync"

type TsvParser struct {
}

var parserInstance *TsvParser = nil

func GetParser() *TsvParser {
	sync.OnceFunc(func() {
		parserInstance = &TsvParser{}
	})
	return parserInstance
}
