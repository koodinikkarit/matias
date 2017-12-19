package program

import (
	"log"
	"os"
	"path/filepath"
)

func getExecutablePath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("getExecutablePath error: %v", err)
	}
	return dir, err
}

func fileExits(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
