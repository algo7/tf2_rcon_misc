package commands

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/algo7/tf2_rcon_misc/network"
)

// getInsult returns an insult for the given target
func getInsult(target string) {
	log.Println("Getting insult for " + target)
	// Set up query parameters
	query := url.Values{}
	query.Set("plural", "true")
	query.Set("template", target+" is <article target=adj1> <adjective min=3 max=5 id=adj1> <amount> like <article target=adj2> <adjective min=1 max=3 id=adj2> <adverb> <animal>")

	// Send GET request to API
	resp, err := http.Get("https://insult.mattbas.org/api/insult.json?" + query.Encode())
	if err != nil {
		log.Printf("Error while calling the Insult API: %v", err)
	}
	defer resp.Body.Close()

	// Parse response JSON
	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("Error while decoding the Insult API response: %v", err)
	}

	// Extract insult from response data
	insult, ok := data["insult"].(string)
	if !ok {
		log.Println("Error while parsing the Insult API response")
	}

	time.Sleep(1000 * time.Millisecond)

	network.RconExecute("say \"" + insult + "\"")
	log.Println("Insult: " + insult)
}
