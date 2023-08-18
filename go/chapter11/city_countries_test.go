package main

import "testing"

func Test_cityCountry(t *testing.T) {
	value := cityCountry("Madrid", "Spain")

	if value != "Madrid, Spain" {
		t.Errorf("Test FAILED: Expected 'Madrid, Spain' got %s", value)
	}
}
