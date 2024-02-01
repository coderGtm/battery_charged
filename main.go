package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isCharging() (bool, error) {
	dat, err := os.ReadFile("/sys/class/power_supply/BAT1/status")
	if (err != nil) {
		return false, err
	}
	return string(dat) == "Charging\n", nil
}

func getBatteryLevel() (int, error) {
	dat, err := os.ReadFile("/sys/class/power_supply/BAT1/capacity")
	if (err != nil) {
		return -1, err
	}
	dat = dat[:len(dat)-1]	// remove \n
	level, err := strconv.Atoi(string(dat))
	if (err != nil) {
		return -1, err
	}
	return level, nil
}

func notify(message string) {
	cmd := "notify-send"
	title := "Battery Notifier"
	err := exec.Command(cmd, title, message).Run()
	check(err)
}

func main() {
	intervalWhileCharging := 60
	intervalWhileDischarging := 180
	exitIfDischarging := false
	highChargeLevel := 90

	for {
		charging, err := isCharging()
		check(err)
		if (charging) {
			fmt.Println(":: Battery is charging. Checking level...")
			batteryLevel, err := getBatteryLevel()
			check(err)
			fmt.Println("Battery level is " + strconv.Itoa(batteryLevel) + "%")
			if batteryLevel >= highChargeLevel {
				fmt.Println("=> Battery level is high. Notifying...")
				notify("High Battery Level, Plese stop charging! [" + strconv.Itoa(batteryLevel) + "%]")
			}
			fmt.Println(":: Checking again in " + strconv.Itoa(intervalWhileCharging) + " seconds...")
			time.Sleep(time.Duration(intervalWhileCharging) * time.Second)
		} else if (!exitIfDischarging) {
			fmt.Println(":: Battery is discharging. Checking again in " + strconv.Itoa(intervalWhileDischarging) + " seconds...")
			time.Sleep(time.Duration(intervalWhileDischarging) * time.Second)
		} else {
			fmt.Println(":: Battery is discharging. Exiting...")
			break
		}
	}
}