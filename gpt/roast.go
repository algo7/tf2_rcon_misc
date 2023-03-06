package gpt

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"tf2-rcon/network"
	"tf2-rcon/utils"
	"time"
)

// GetInsult returns an insult for the given target
func GetInsult(target string) {
	fmt.Println("Getting insult for " + target)
	// Set up query parameters
	query := url.Values{}
	query.Set("plural", "true")
	query.Set("template", target+" is <article target=adj1> <adjective min=3 max=5 id=adj1> <amount> like <article target=adj2> <adjective min=1 max=3 id=adj2> <adverb> <animal>")

	// Send GET request to API
	resp, err := http.Get("https://insult.mattbas.org/api/insult.json?" + query.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse response JSON
	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		utils.ErrorHandler(err)
	}

	// Extract insult from response data
	insult, ok := data["insult"].(string)
	if !ok {
		utils.ErrorHandler(errors.New("Could not parse insult from response data"))
	}

	time.Sleep(1000 * time.Millisecond)

	network.RconExecute("say \"" + insult + "\"")
	fmt.Println("Insult: " + insult)
}
