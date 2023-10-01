package commands

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/algo7/tf2_rcon_misc/network"
)

// getCompliment returns a compliment for the given target
func getCompliment(target string) {
	log.Println("Getting compliment for " + target)

	// Send GET request to API
	resp, err := http.Get("https://complimentr.com/api")
	if err != nil {
		log.Panicf("Error while calling the Compliment API: %v", err)
	}
	defer resp.Body.Close()

	// Parse response JSON
	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("Error while decoding the Compliment API response: %v", err)
	}

	// Extract insult from response data
	compliment, ok := data["compliment"].(string)
	if !ok {
		log.Println("Error while parsing the Compliment API response")
	}

	time.Sleep(1000 * time.Millisecond)

	network.RconExecute("say \"" + target + " " + compliment + "\"")
	log.Println("Compliment: " + compliment)
}
