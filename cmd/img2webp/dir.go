package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Wohlgemuth-Dev/img2webp/internal/ui"
)

func getBaseDirectory() string {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("✕ Failed to get executable path: %v", err)))
		os.Exit(1)
	}
	baseDir := filepath.Dir(exePath)

	if strings.Contains(strings.ToLower(baseDir), "temp") || strings.Contains(strings.ToLower(baseDir), "tmp") {
		baseDir, _ = os.Getwd()
	}

	return baseDir
}

func findImages(baseDir string) []string {
	var imgFiles []string

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				imgFiles = append(imgFiles, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("✖ Error walking directory: %v", err)))
		os.Exit(1)
	}

	return imgFiles
}
