package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments provided")
		os.Exit(-1)
	}

	exampleTitle := os.Args[1]

	// writing main file
	contents, err := ioutil.ReadFile("main.cpp")
	if err != nil {
		fmt.Println("Could not open main.cpp for reading - verify this is being run in Sandbox/src")
		os.Exit(-1)
	}

	newContents := bytes.Replace(contents, []byte("main"), []byte(exampleTitle), -1)

	if err := os.Mkdir(fmt.Sprintf("examples/%s", exampleTitle), 0755); err != nil {
		fmt.Printf("Could not create directory titled '%s'\n", exampleTitle)
		os.Exit(-1)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("examples/%s/%s.cpp", exampleTitle, exampleTitle), newContents, 0666); err != nil {
		fmt.Printf("Could not write file 'examples/%s/%s.cpp'", exampleTitle, exampleTitle)
		os.Exit(-1)
	}

}
