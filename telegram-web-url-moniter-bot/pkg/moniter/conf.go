package moniter

import (
)

var dBFilePath string

func init() {
	dBFilePath = "test.db"
}

func SetDBFilePath(filepath string) {
	dBFilePath = filepath
}