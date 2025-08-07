package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func createEvidenceFiles() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	documentsPath := filepath.Join(usr.HomeDir, "Documents")

	files := []string{"evidence1.txt", "evidence2.txt"}
	for _, fname := range files {
		fullPath := filepath.Join(documentsPath, fname)
		err := os.WriteFile(fullPath, []byte("You found me"), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	createEvidenceFiles()
}
