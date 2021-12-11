package main

import "testing"

func TestCheckIPAddressNotValid(t *testing.T) {

	ip := "0.0.0"

	expected := false

	actual := checkIPAddress(ip)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}

func TestCheckIPAddressValid(t *testing.T) {

	ip := "0.0.0.0"

	expected := true

	actual := checkIPAddress(ip)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}

func TestConvertDepthStringValid(t *testing.T) {

	depth := "5"

	expected := 5

	actual, _ := convertDepthString(depth)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}

func TestConvertDepthStringNotValid_NotInt(t *testing.T) {

	depth := "eeewrwe"

	expected := -1

	actual, _ := convertDepthString(depth)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}

func TestConvertDepthStringNotValid_ToBig(t *testing.T) {

	depth := "7"

	expected := 4

	actual, _ := convertDepthString(depth)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}

func TestConvertDepthStringNotValid_ToSmall(t *testing.T) {

	depth := "0"

	expected := 4

	actual, _ := convertDepthString(depth)

	if expected != actual {
		t.Errorf("Expected %v but actual value was %v", expected, actual)
	}

}
