package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hacdias/fileutils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments provided")
		os.Exit(-1)
	}
	copyDirectory("test", "testing")

	exampleTitle := os.Args[1]

	// creating directory
	if err := os.Mkdir(fmt.Sprintf("examples/%s", exampleTitle), 0755); err != nil {
		fmt.Printf("Could not create directory titled 'examples/%s'\n", exampleTitle)
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// writing source file
	contents, err := ioutil.ReadFile("main.cpp")
	if err != nil {
		fmt.Println("Could not open main.cpp for reading - verify this is being run in Sandbox/src")
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	newContents := fmt.Sprintf("#include \"%s.h\"\n", exampleTitle)
	newContents += string(bytes.Replace(contents, []byte("main"), []byte(exampleTitle), -1))

	if err := ioutil.WriteFile(fmt.Sprintf("examples/%s/%s.cpp", exampleTitle, exampleTitle),
		[]byte(newContents), 0666); err != nil {
		fmt.Printf("Could not write file 'examples/%s/%s.cpp'", exampleTitle, exampleTitle)
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// writing header file
	headerContents := fmt.Sprintf(`#pragma once

void %s();
	`, exampleTitle)

	if err := ioutil.WriteFile(fmt.Sprintf("examples/%s/%s.h", exampleTitle, exampleTitle),
		[]byte(headerContents), 0666); err != nil {
		fmt.Printf("Could not write file 'examples/%s/%s.h'", exampleTitle, exampleTitle)
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// copying shaders
	if err := copyDirectory(exampleTitle, "shaders"); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// copying resources
	if err := copyDirectory(exampleTitle, "resources"); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

}

func copyDirectory(parentDir string, name string) error {

	destPath := filepath.Join("examples", parentDir, name)

	if err := os.Mkdir(destPath, os.FileMode(0777)); err != nil {
		return fmt.Errorf(fmt.Sprintf("Could not create directory titled 'examples/%s/%s'",
			parentDir, name))
	}
	if err := fileutils.CopyDir(name, destPath); err != nil {
		return fmt.Errorf(fmt.Sprintf("Could not copy directory titled '%s'", name))
	}
	return nil
}
