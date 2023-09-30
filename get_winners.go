package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func GetTeamMap() map[string]int {
	// Return a new copy of the map every time
	return map[string]int{
		"falcons":    1,
		"bills":      2,
		"bears":      3,
		"bengals":    4,
		"browns":     5,
		"cowboys":    6,
		"broncos":    7,
		"lions":      8,
		"packers":    9,
		"titans":     10,
		"colts":      11,
		"chiefs":     12,
		"raiders":    13,
		"rams":       14,
		"dolphins":   15,
		"vikings":    16,
		"patriots":   17,
		"saints":     18,
		"giants":     19,
		"jets":       20,
		"eagles":     21,
		"cardinals":  22,
		"steelers":   23,
		"chargers":   24,
		"49ers":      25,
		"seahawks":   26,
		"buccaneers": 27,
		"commanders": 28,
		"panthers":   29,
		"jaguars":    30,
		"ravnes":     33,
		"texans":     34,
	}
}

func GetResultsJson(week, year string) []byte {
	url := fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?dates=%s&week=%s", year, week)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make the HTTP request: %v", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	return bodyBytes
}

func CreateResultMap(payload map[string]any) map[int]bool {

	winnerMap := make(map[int]bool)
	eventSlice := payload["events"].([]interface{})
	for _, event := range eventSlice {
		competitionsSlice := event.(map[string]interface{})["competitions"].([]interface{})
		for _, competition := range competitionsSlice {
			competitorsSlice := competition.(map[string]interface{})["competitors"].([]interface{})
			for _, competitor := range competitorsSlice {
				intVar, _ := strconv.Atoi(competitor.(map[string]interface{})["id"].(string))
				winnerMap[intVar] = competitor.(map[string]interface{})["winner"].(bool)
			}
		}
	}
	return winnerMap
}

func teamWon(name string, payload map[string]any) bool {
	resultMap := CreateResultMap(payload)
	lowerCaseName := strings.ToLower(name)
	teamMap := GetTeamMap()
	return resultMap[teamMap[lowerCaseName]]
}

func GetCurrentFootballSeasonYear() string {
	var year = time.Now().Year()
	if time.Now().Month() < 4 {
		year = year + 1
	}
	return strconv.Itoa(year)
}

func GetWeek(s string) string {
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindString(s)

	if match == "" {
		fmt.Println("Make sure the week number is in the third cell of the first row")
		os.Exit(1)
	}

	return match
}

func main() {
	// Load the Excel file
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please call this will an excel file")
		fmt.Println("Example: go run get_winnners.go /home/austin/Downloads/test.xlsx")
		os.Exit(1)
	}
	excelFileName := args[1]
	newExcelFileName := "calculatedWins.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	// Assuming the data is in the first sheet
	sheet := xlFile.Sheets[0]
	style := xlsx.NewStyle()
	style.Fill = *xlsx.NewFill("solid", "FF0000", "FFFFFFFF")
	style.ApplyFill = true

	currentWeek := GetWeek(sheet.Rows[0].Cells[2].String())
	currentYear := GetCurrentFootballSeasonYear()
	if len(os.Args) > 2 {
		currentYear = os.Args[2]
	}
	fmt.Printf("Current week %s and football year is %s\n", currentWeek, currentYear)
	content := GetResultsJson(currentWeek, currentYear)
	var payload map[string]any
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	for _, row := range sheet.Rows[1:] {

		var cells []string
		for _, cell := range row.Cells {
			text := cell.String()
			cells = append(cells, text)
		}
		if !teamWon(cells[2], payload) {
			cell := row.Cells[2]
			cell.SetStyle(style)
		}
	}

	if err := xlFile.Save(newExcelFileName); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}
}
