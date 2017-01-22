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
 * Termux API: termux-telephony-deviceinfo
 */

var Deviceinfo struct {
	Data_activity           string
	Data_state              string
	Device_id               string
	Device_software_version string
	Phone_count             int
	Phone_type              string
	Network_operator        string
	Network_operator_name   string
	Network_country_iso     string
	Network_type            string
	Network_roaming         bool
	Sim_country_iso         string
	Sim_operator            string
	Sim_operator_name       string
	Sim_serial_number       string
	Sim_state               string
}

func main() {
	cmd := exec.Command("termux-telephony-deviceinfo")

	//StdoutPipe returns a pipe that will be connected to
	//the command's standard output when the command starts
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(stdout).Decode(&Deviceinfo); err != nil {
		log.Fatal(err)
	}

	//Wait waits for the command to exit
	//It must have been started by Start
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your network type: %s, network operator name: %s\n",
		Deviceinfo.Network_type,
		Deviceinfo.Network_operator_name)
}
