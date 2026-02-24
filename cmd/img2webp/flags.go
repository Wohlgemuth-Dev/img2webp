package main

import (
	"flag"
	"fmt"
	"os"

	"img2webp/internal/ui"
)

func parseFlags() (float64, bool) {
	var quality float64
	flag.Float64Var(&quality, "quality", 80.0, "Quality of the WebP output (0-100)")
	flag.Float64Var(&quality, "q", 80.0, "Shorthand for quality")

	var overwrite bool
	flag.BoolVar(&overwrite, "overwrite", false, "Overwrite existing WebP files")
	flag.BoolVar(&overwrite, "o", false, "Shorthand for overwrite")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Recursively finds all .jpg/.jpeg/.png files in the current directory (and subdirectories) and converts them to WebP.\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if quality < 0 || quality > 100 {
		fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("✖ Error: Quality must be between 0 and 100. Provided: %f", quality)))
		os.Exit(1)
	}

	return quality, overwrite
}
