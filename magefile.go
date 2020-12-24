// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	exeFileName = "mygen.exe"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Run

// Clean clean up after yourself
func Clean() {
	fmt.Println("Clean...")
	os.Remove(exeFileName)
}

// Build build
func Build() error {
	mg.Deps(Clean)
	fmt.Println("Build...")
	return sh.RunV("go", "build", "-o", exeFileName, ".")
}

// Run execute app
func Run() error {
	mg.Deps(Build)
	fmt.Println("Run...")
	return sh.RunV("./"+exeFileName, "-id", "13", "-base", "1")
}

// Test execute test
func Test() error {
	fmt.Println("Test...")
	return sh.RunV("go", "test")
}

// Generate generate code
func Generate() error {
	fmt.Println("Generate...")
	folders := []string{
		"layer",
		"mode",
		"palette",
		"part",
		"sprite",
		"status",
	}
	for _, folder := range folders {
		path := fmt.Sprintf("./%s", folder)
		err := sh.RunV("go", "generate", path)
		if err != nil {
			return err
		}
	}
	return nil
}

// Statik
func Statik() error {
	fmt.Println("Statik...")
	return sh.RunV("statik", "-f", "-src", "_assets", "-dest", "assets", "-tags", "release")
}

// Release
func Release() error {
	mg.Deps(Test)
	mg.Deps(Statik)
	fmt.Println("Release Build...")
	err := sh.RunV("go", "build", "-tags", "release", "-o", exeFileName, ".")
	if err != nil {
		return err
	}
	fmt.Println("Release Run...")
	return sh.RunV("./" + exeFileName)
}
