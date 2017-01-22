/**
 * @Author Rafal Golarz
 * http://rafalgolarz.com/blog/2017/01/15/running_golang_on_android/
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

/*
 * Termux API: termux-battery-status
 */

var Battery struct {
	Health      string
	Percentage  int
	Status      string
	Temperature float64
}

func main() {
	cmd := exec.Command("termux-battery-status")

	//StdoutPipe returns a pipe that will be connected to
	//the command's standard output when the command starts
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(stdout).Decode(&Battery); err != nil {
		log.Fatal(err)
	}

	//Wait waits for the command to exit
	//It must have been started by Start
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The battery is %s and is charged in %d percentage\n",
		Battery.Health, Battery.Percentage)
}
