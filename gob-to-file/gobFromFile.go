package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"
)

type Lessons []Lesson

type Lesson struct {
	Time time.Time
	Note string
}

func load(lessons *Lessons) error {
	f, err := os.OpenFile("data.gob", os.O_RDONLY, 0664)
	if err != nil {
		log.Printf("error opening file for reading: %v", err)
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
			return
		}
	}()
	enc := gob.NewDecoder(f)
	err = enc.Decode(lessons)
	if err != nil {
		log.Printf("error decoding data: %v", err)
		return err
	}
	return nil
}

func main() {
	lessons := Lessons{}
	err := load(&lessons)
	if err != nil {
		log.Printf("could not load data: %v", err)
		os.Exit(1)
	}
	for _, lesson := range lessons {
		fmt.Printf("%v\t%s\n", lesson.Time, lesson.Note)
	}
}
