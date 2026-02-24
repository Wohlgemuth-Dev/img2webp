package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/Wohlgemuth-Dev/img2webp/internal/converter"
	"github.com/Wohlgemuth-Dev/img2webp/internal/sysutil"
	"github.com/Wohlgemuth-Dev/img2webp/internal/ui"
)

func processImages(imgFiles []string, quality float64, overwrite bool) int {
	numWorkers := runtime.NumCPU()
	jobs := make(chan string, len(imgFiles))
	results := make(chan error, len(imgFiles))
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go converter.Worker(jobs, results, quality, overwrite, &wg)
	}

	for _, file := range imgFiles {
		jobs <- file
	}
	close(jobs)

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

	return errCount
}

func printResults(totalCount, errCount int, elapsed time.Duration) {
	successCount := totalCount - errCount

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
