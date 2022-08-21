package jvdata

import (
	"log"
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	files, err := filepath.Glob("/data/RA*.dat")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			_, err := raRead(file)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
