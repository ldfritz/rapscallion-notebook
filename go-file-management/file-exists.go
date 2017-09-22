package main

import (
	"log"
	"os"
)

func main() {
	for _, f := range []string{"file-exists.go", "does-not-exist"} {
		if _, err := os.Stat(f); err != nil {
			log.Printf("Unable to confirm %s existences: %v", f, err)
			continue
		}
		log.Printf("file found: %s", f)
	}
}
