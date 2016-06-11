package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// Handle a file entry
func HandlePath(filepath string) {
	file, err := os.Open(filepath) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// handle a directory
	if info.IsDir() {
		dirInfos, err := file.Readdir(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, dirInfo := range dirInfos {
			joinPath := path.Join(filepath, dirInfo.Name())
			HandlePath(joinPath)
		}
	} else {
		// regular file

		fmt.Println(filepath)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		HandlePath(arg)
	}
}
