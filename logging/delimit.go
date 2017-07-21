// Another cool trick I saw in a video.
package main

import (
	"log"
	"strings"
)

func main() {
	log.Printf(strings.Repeat("-", 20))
	defer log.Printf(strings.Repeat("-", 20))
	log.Printf("This is my message.")
}
