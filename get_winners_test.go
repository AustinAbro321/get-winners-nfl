package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDidTeamWin(t *testing.T) {
	result := readJson("file.json")
	got := teamWon("Lions",result)
	expected := true

	if got != expected {
		t.Errorf("expected '%t' but got '%t'", expected, got)
	}
}

func TestCreateResultMap(t *testing.T) {
	result := readJson("test.json")

	got := CreateResultMap(result)
	expected := make(map[int]bool)

	expected[1] = true
	expected[2] = false
	expected[3] = true
	expected[4] = false

	if !mapsEqual(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestCreateResultMapAll(t *testing.T) {	
	result := readJson("file.json")
	got := CreateResultMap(result)
	expected := 32

	if len(got) != 32 {
		t.Errorf("Expected %d but got %d", expected, len(got))
	}
}

func TestGetWeekNumber(t *testing.T) {	
	got := GetWeek("week 13")
	expected := "13"

	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func readJson(filename string) map[string]interface{} {

	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func mapsEqual(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bValue, exists := b[k]; !exists || v != bValue {
			return false
		}
	}

	return true
}
