package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"jpg2webp/internal/converter"
	"jpg2webp/internal/sysutil"
	"jpg2webp/internal/ui"
)

func main() {
	// Parse command line flags
	var quality float64
	flag.Float64Var(&quality, "quality", 80.0, "Quality of the WebP output (0-100)")
	flag.Float64Var(&quality, "q", 80.0, "Shorthand for quality")

	var overwrite bool
	flag.BoolVar(&overwrite, "overwrite", false, "Overwrite existing WebP files")
	flag.BoolVar(&overwrite, "o", false, "Shorthand for overwrite")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Recursively finds all .jpg/.jpeg/.png files in the current directory (and subdirectories) and converts them to WebP.\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Ensure quality is within valid bounds
	if quality < 0 || quality > 100 {
		fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("✖ Error: Quality must be between 0 and 100. Provided: %f", quality)))
		os.Exit(1)
	}

	fmt.Printf("\n%s %s\n\n",
		ui.TitleStyle.Render("Image → WebP Converter"),
		ui.DimStyle.Render("(run with -h for help)"))

	// Find the directory where the executable is located, or just use the current working directory
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("✕ Failed to get executable path: %v", err)))
		os.Exit(1)
	}
	baseDir := filepath.Dir(exePath)

	// For development (go run), os.Executable() might return a high-entropy temp directory
	if strings.Contains(strings.ToLower(baseDir), "temp") || strings.Contains(strings.ToLower(baseDir), "tmp") {
		baseDir, _ = os.Getwd()
	}

	fmt.Printf("%s %s\n", ui.InfoStyle.Render("• Searching in:"), ui.Highlight.Render(baseDir))
	fmt.Printf("%s %s\n", ui.InfoStyle.Render("• Quality:"), ui.Highlight.Render(fmt.Sprintf("%.2f%%", quality)))
	fmt.Println()

	// Find all supported image files
	var imgFiles []string

	err = filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
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

	if len(imgFiles) == 0 {
		fmt.Println(ui.WarnStyle.Render("⚠ No supported images (JPEG/PNG) found in the directory."))
		return
	}

	fmt.Printf("%s\n\n", ui.InfoStyle.Render(fmt.Sprintf("→ Found %d images. Starting conversion...", len(imgFiles))))

	startTime := time.Now()

	// Concurrency setup
	numWorkers := runtime.NumCPU()
	jobs := make(chan string, len(imgFiles))
	results := make(chan error, len(imgFiles))
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go converter.Worker(jobs, results, quality, overwrite, &wg)
	}

	// Send jobs
	for _, file := range imgFiles {
		jobs <- file
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	var errCount int
	for err := range results {
		if err != nil {
			fmt.Printf("%s %v\n", ui.ErrorStyle.Render("✖"), err)
			errCount++
		}
	}

	elapsed := time.Since(startTime)
	successCount := len(imgFiles) - errCount

	fmt.Println()
	if errCount > 0 {
		fmt.Printf("%s Converted %s images in %s (%s errors)\n",
			ui.SuccessStyle.Render("✓ Done:"),
			ui.Highlight.Render(fmt.Sprint(successCount)),
			ui.Highlight.Render(elapsed.String()),
			ui.ErrorStyle.Render(fmt.Sprint(errCount)))
	} else {
		fmt.Printf("%s Converted all %s images in %s\n",
			ui.SuccessStyle.Render("✓ Done:"),
			ui.Highlight.Render(fmt.Sprint(successCount)),
			ui.Highlight.Render(elapsed.String()))
	}
	fmt.Println()

	if sysutil.IsLaunchedFromExplorer() {
		fmt.Printf("%s\n", ui.DimStyle.Render("Press Enter to exit..."))
		fmt.Scanln()
	}
}
