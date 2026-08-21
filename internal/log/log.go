package log

import (
	"log"
	"os"
)

var (
	Debug = log.New(os.Stdout, "[debug] ", log.Lshortfile)
	Warn  = log.New(os.Stderr, "[warn] ", log.Lshortfile)
)
