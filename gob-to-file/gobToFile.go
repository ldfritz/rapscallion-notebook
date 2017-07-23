package main

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

type Lessons []Lesson

type Lesson struct {
	Time time.Time
	Note string
}

func save(data interface{}) error {
	f, err := os.OpenFile("data.gob", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Printf("error opening file for writing: %v", err)
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
			return
		}
	}()
	enc := gob.NewEncoder(f)
	err = enc.Encode(data)
	if err != nil {
		log.Printf("error encoding data: %v", err)
		return err
	}
	return nil
}

func main() {
	t := time.Now()
	lessons := Lessons{
		{t, "no time like the present"},
		{t.Add(30 * time.Minute), "the moment has passed"},
	}
	err := save(lessons)
	if err != nil {
		log.Printf("could not save data: %v", err)
		os.Exit(1)
	}
}
