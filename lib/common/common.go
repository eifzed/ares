package common

import (
	"io"
	"os"

	"github.com/labstack/gommon/log"
)

func SafelyCloseFile(f io.Closer) {
	if err := f.Close(); err != nil {
		log.Warnf("Failed to close file: %s\n", err)
	}
}

func IsDevelopment() bool {
	isLocal := os.Getenv("ISLOCAL")
	return isLocal == "1"
}

const (
	TimeYYYMMDD_Dash       = "2006-01-02"
	TimeYYYMMDDHHMMSS_Dash = "2006-01-02 15:04:05"
)
