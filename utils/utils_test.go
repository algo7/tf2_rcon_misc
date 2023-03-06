package utils

import "testing"

func TestSteam3IDToSteam64(t *testing.T) {
	t.Log("Testing [U:1:1524567963]")
	Steam3IDToSteam64("[U:1:1524567963]")

	t.Log("Testing [U:1:259772]")
	Steam3IDToSteam64("[U:1:259772]")
}