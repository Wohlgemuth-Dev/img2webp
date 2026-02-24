package converter

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"sync"

	"github.com/chai2010/webp"

	"github.com/Wohlgemuth-Dev/img2webp/internal/ui"
)

func Worker(jobs <-chan string, results chan<- error, quality float64, overwrite bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for imgPath := range jobs {
		results <- ConvertToWebP(imgPath, float32(quality), overwrite)
	}
}

func ConvertToWebP(imgPath string, quality float32, overwrite bool) error {
	// Open the image file
	file, err := os.Open(imgPath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", imgPath, err)
	}
	defer file.Close()

	// Determine output WebP path FIRST so we can skip before decoding
	dir := filepath.Dir(imgPath)
	baseName := filepath.Base(imgPath)
	webpPath := filepath.Join(dir, baseName+".webp")

	// Check if webp file already exists, if so skip it (unless overwrite is true)
	if !overwrite {
		if _, err := os.Stat(webpPath); err == nil {
			// File exists, skip
			fmt.Printf("  %s %s\n", ui.DimStyle.Render("→ Skipped:"), ui.DimStyle.Render(baseName))
			return nil
		}
	}

	// Decode based on magic bytes (ignoring file extension for raster images)
	var img image.Image
	var decodeErr error
	img, _, decodeErr = image.Decode(file)
	if decodeErr != nil {
		return fmt.Errorf("failed to decode image %s: %w", imgPath, decodeErr)
	}

	// Create the WebP file
	out, err := os.Create(webpPath)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", webpPath, err)
	}
	defer out.Close()

	// Encode to WebP (Lossy compression with user-defined quality)
	options := &webp.Options{Lossless: false, Quality: quality}

	err = webp.Encode(out, img, options)
	if err != nil {
		return fmt.Errorf("failed to encode %s to WebP: %w", webpPath, err)
	}

	fmt.Printf("  %s %s\n", ui.SuccessStyle.Render("✓ Converted:"), ui.Highlight.Render(baseName))
	return nil
}
