package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln("unable to stat input", err)
	}

	fmt.Printf("%-14s %s\n", "FileMode", info.Mode())

	for _, v := range []struct {
		Name string
		Mode os.FileMode
	}{
		{"ModeDir", os.ModeDir},
		{"ModeAppend", os.ModeAppend},
		{"ModeExclusive", os.ModeExclusive},
		{"ModeTemporary", os.ModeTemporary},
		{"ModeSymlink", os.ModeSymlink},
		{"ModeDevice", os.ModeDevice},
		{"ModeNamedPipe", os.ModeNamedPipe},
		{"ModeSocket", os.ModeSocket},
		{"ModeSetuid", os.ModeSetuid},
		{"ModeSetgid", os.ModeSetgid},
		{"ModeCharDevice", os.ModeCharDevice},
		{"ModeSticky", os.ModeSticky},
		{"ModeType", os.ModeType},
		{"ModePerm", os.ModePerm},
	} {

		fmt.Printf("%-14s %s\n", v.Name, info.Mode()&v.Mode)
	}

	fmt.Println("")

	if info.Mode()&os.ModeNamedPipe != 0 {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln("unable to read input", err)
		}
		fmt.Print(string(content))
	} else {
		fmt.Println("<no piped input>")
	}
}
