package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mdNoman21/concurrent-download-manager/helpers"
)

func main() {
	startTime := time.Now()

	// Prompt the user to enter a URL
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter URL:")
	if scanner.Scan() {
		url := scanner.Text()
		fmt.Println("URL entered:", url)

		err := helpers.DownloadFromURL(url)
		if err != nil {
			log.Print(err.Error())
		}

		fmt.Printf("Download finished in %v seconds\n", time.Since(startTime).Seconds())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
