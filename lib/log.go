package lib

import (
	"log"
	"runtime"
)

const (
	LEVEL_DEBUG = iota
	LEVEL_ERROR
)

// write log
func Write(level int, message error) {
	lvl := [2]string{"debug", "error"}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		if level == LEVEL_DEBUG {
			log.Printf("%s: %s File: %s Line: %d \n", lvl[level], message.Error(), file, line)
		} else {
			log.Fatalf("%s: %s File: %s Line: %d \n", lvl[level], message.Error(), file, line)
		}
	}
}
