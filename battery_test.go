package main

import "testing"

func TestIsCharging(t *testing.T) {
	_, err := isCharging()
	if err != nil {
		t.Errorf("Error while getting Charging Status: %v", err)
	}
}

func TestBatteryLevel(t *testing.T) {
	level, err := getBatteryLevel()
	if err != nil {
		t.Errorf("Error while getting Battery Level: %v", err)
	}
	if level < 0 || level > 100 {
		t.Errorf("Invalid battery level: %d", level)
	}
}