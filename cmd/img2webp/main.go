package main

import (
	"fmt"
	"time"

	"img2webp/internal/ui"
)

func main() {
	quality, overwrite := parseFlags()

	fmt.Printf("\n%s %s\n\n",
		ui.TitleStyle.Render("Image → WebP Converter"),
		ui.DimStyle.Render("(run with -h for help)"))

	baseDir := getBaseDirectory()

	fmt.Printf("%s %s\n", ui.InfoStyle.Render("• Searching in:"), ui.Highlight.Render(baseDir))
	fmt.Printf("%s %s\n", ui.InfoStyle.Render("• Quality:"), ui.Highlight.Render(fmt.Sprintf("%.2f%%", quality)))
	fmt.Println()

	imgFiles := findImages(baseDir)

	if len(imgFiles) == 0 {
		fmt.Println(ui.WarnStyle.Render("⚠ No supported images (JPEG/PNG) found in the directory."))
		return
	}

	fmt.Printf("%s\n\n", ui.InfoStyle.Render(fmt.Sprintf("→ Found %d images. Starting conversion...", len(imgFiles))))

	startTime := time.Now()
	errCount := processImages(imgFiles, quality, overwrite)
	elapsed := time.Since(startTime)

	printResults(len(imgFiles), errCount, elapsed)
}
