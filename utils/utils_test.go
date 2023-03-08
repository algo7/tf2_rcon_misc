package utils

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSteam3IDToSteam64(t *testing.T) {
	t.Log("Testing [U:1:1524567963]")
	Steam3IDToSteam64("[U:1:1524567963]")

	t.Log("Testing [U:1:259772]")
	Steam3IDToSteam64("[U:1:259772]")
}

func TestGetPlayerNameFromLine(t *testing.T) {
	const testString1 = "#   2006 \"atomy\"             [U:1:259772]        04:44       45    0 active"
	t.Log(testString1)
	if result := GetPlayerNameFromLine(testString1); result != "atomy" {
		t.Errorf("Expected 'atomy' as return, but got '%s'", result)
	}

	const testString2 = "#   3206 \"RussianVesper\"     [U:1:303060879]     01:33       63   75 spawning"
	t.Log(testString2)
	if result := GetPlayerNameFromLine(testString2); result != "RussianVesper" {
		t.Errorf("Expected 'RussianVesper' as return, but got '%s'", result)
	}
}

func TestGetChatSay(t *testing.T) {
	// open fixture console.log file for usage
	file, err := os.Open("../test/fixtures/console.log")
	if err != nil {
		t.Errorf("Got error while opening fixture file: '%s'", err)
	}

	var players = []string{"The.Real.Genesis", "gibb (official)"}
	var chatLines = []string{}
	var chatUsers = []string{}
	var chatText = []string{}

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		isSay, player, text := GetChatSay(players, line)

		if isSay {
			chatLines = append(chatLines, line)
			chatUsers = append(chatUsers, player)
			chatText = append(chatText, text)
			fmt.Println("dbg1:", line)
			fmt.Println("dbg2:", isSay, player, text)
		}
	}

	if 2 != len(chatLines) {
		t.Errorf("Expected 2 lines to be chat, but found '%d'", len(chatLines))
	}

	if chatText[0] != "Hey guys, thanks for tuning in for another video on x.com, I'm Ian Mccollum and I'm here today at the Rock Islan" {
		t.Errorf("Expected chatText-0 to be 'Hey guys, thanks for tuning in for another video on x.com, I'm Ian Mccollum and I'm here today at the Rock Islan'\n but found '%s'", chatText[0])
	}

	if chatUsers[0] != "The.Real.Genesis" {
		t.Errorf("Expected chatusers-0 to be 'The.Real.Genesis' but found '%s'", chatUsers[0])
	}

	if chatText[1] != "00:33 gaming" {
		t.Errorf("Expected chatText-1 to be '00:33 gaming'\n but found '%s'", chatText[1])
	}

	if chatUsers[1] != "gibb (official)" {
		t.Errorf("Expected chatusers-1 to be 'gibb (official)' but found '%s'", chatUsers[1])
	}
}
