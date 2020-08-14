package content_test

import (
	"fmt"
	"github.com/SpencerBrown/content"
	"log"
	"os"
	"path/filepath"
)

func main_example() {
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		// Handle generating static assets
		log.Println("Generating static assets as Go source code...")
		err := generateExample()
		if err != nil {
			log.Fatalf("Error generating static assets: %v\n", err)
		}
		log.Println("Done generating Go source code")
		return
	}
	// normal invocation of the binary starts here
}

// create the static.go file from the static content files in the "static-content" directory via "go:generate"
//go:generate example generate $GOPACKAGE
func generateExample() error {

	const staticOutputName = "static.go"
	const staticContentDir = "static-content"

	staticFilesDir, err := filepath.Abs(staticContentDir)
	if err != nil {
		return fmt.Errorf("unable to locate static files: %v", err)
	}

	staticOutDir := filepath.Dir(staticFilesDir) // output .go file is one level up from static content files
	var packageName = os.Args[2]

	err = content.GenerateContent(staticFilesDir, staticOutDir, staticOutputName, packageName)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
