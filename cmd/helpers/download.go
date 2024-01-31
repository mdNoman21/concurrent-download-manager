package helpers

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/mdNoman21/concurrent-download-manager/models"
)

var validate = validator.New()

func DownloadFromURL(url string) error {
	// Create a Download instance with the provided URL and other settings
	d := models.Download{
		Url:           url, // Assuming you have a URL field in your Download struct
		TargetPath:    "final.mp4",
		TotalSections: 10,
	}

	// Validate the Download struct
	if err := validate.Struct(d); err != nil {
		// Handle validation errors
		return err
	}

	// Start the download
	if err := d.Do(); err != nil {
		// Handle download errors
		log.Printf("An error occurred while downloading the file: %s\n", err)
		return err
	}

	return nil
}
