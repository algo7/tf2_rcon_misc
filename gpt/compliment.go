package gpt

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/algo7/tf2_rcon_misc/network"
	"github.com/algo7/tf2_rcon_misc/utils"
)

// GetCompliment returns a compliment for the given target
func GetCompliment(target string) {
	fmt.Println("Getting compliment for " + target)

	// Send GET request to API
	resp, err := http.Get("https://complimentr.com/api")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse response JSON
	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		utils.ErrorHandler(err, false)
	}

	// Extract insult from response data
	compliment, ok := data["compliment"].(string)
	if !ok {
		utils.ErrorHandler(errors.New("Could not parse insult from response data"), false)
	}

	time.Sleep(1000 * time.Millisecond)

	network.RconExecute("say \"" + target + " " + compliment + "\"")
	fmt.Println("Compliment: " + compliment)
}
