package util

import (
	"os"
	"path/filepath"
)

type CardData struct {
	PAN        string
	Expiry     string
	Cardholder string
}

func Path() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
