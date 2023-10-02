package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/algo7/tf2_rcon_misc/commands"
	"github.com/algo7/tf2_rcon_misc/db"
	"github.com/algo7/tf2_rcon_misc/network"
	"github.com/algo7/tf2_rcon_misc/utils"
)

// Const console message that informs you about forceful autobalance.
const teamSwitchMessage = "You have switched to team BLU and will receive 500 experience points at the end of the round for changing teams."

// playersInGame is a slice of player info cache struct that holds the player info
var playersInGame []*utils.PlayerInfo

func main() {

	// Init the grok patterns
	utils.GrokInit()

	// Connect to the rcon server
	network.Connect()

	if network.RCONConnection == nil {
		log.Println("Connection to RCON failed")
	}

	// Get the current player name
	res := network.RconExecute("name")
	currentPlayer, err := utils.GrokParsePlayerName(res)

	if err != nil {
		log.Fatalf("%v Please restart the program", err)
	}

	log.Printf("Current plyaer is %s", currentPlayer)

	// Get log path
	tf2LogPath := utils.LogPathDection()

	// Empty the log file
	err = utils.EmptyLog(tf2LogPath)

	if err != nil {
		log.Fatalf("Unable to empty the log file: %v", err)
	}

	// Tail the log
	log.Println("Tailing Logfile at:", tf2LogPath)
	t, err := utils.TailLog(tf2LogPath)
	if err != nil {
		log.Fatalf("Unable to tail the log file: %v", err)
	}

	// Loop through the text of each received line
	for line := range t.Lines {

		// Parse the line for player info
		playerInfo, err := utils.GrokParse(line.Text)
		if err != nil {

		}

		// Parse the line for chat info
		chat, err := utils.GrokParseChat(line.Text)
		if err != nil {

		}

		// Refresh player list logic
		// Dont assume status headlines as player connects
		if strings.Contains(line.Text, "Lobby updated") || (strings.Contains(line.Text, "connected") && !strings.Contains(line.Text, "uniqueid")) {
			log.Printf("Executing *status* command after line: %s", line.Text)

			// Clear the player list
			playersInGame = []*utils.PlayerInfo{}
			fmt.Println("playersInGame", playersInGame)
			// Run the status command when the lobby is updated or a player connects
			network.RconExecute("status")
		}

		// Save to DB logic
		if playerInfo != nil {

			log.Printf("%+v\n", *playerInfo)

			// Append the player to the player list
			playersInGame = append(playersInGame, playerInfo)

			// Create a player document for inserting into MongoDB
			player := db.Player{
				SteamID:   playerInfo.SteamID,
				Name:      playerInfo.Name,
				UpdatedAt: time.Now().UnixNano(),
			}

			// Add the player to the DB
			db.AddPlayer(player)
		}

		if chat != nil {
			log.Printf("Chat: %+v\n", *chat)

			// Parse the chat message for commands
			command, args, err := utils.GrokParseCommand(chat.Message)

			// If chat message is a command
			if err == nil {
				commands.CommandExecuted(command, args, chat.PlayerName, currentPlayer)
			} else {
				// Get the player's steamID64 from the playersInGame
				steamID, err := utils.GetSteamIDFromPlayerName(chat.PlayerName, playersInGame)

				if err != nil {
					// Create a chat document for inserting into MongoDB
					chatInfo := db.Chat{
						SteamID:   steamID,
						Name:      chat.PlayerName,
						Message:   chat.Message,
						UpdatedAt: time.Now().UnixNano(),
					}

					db.AddChat(chatInfo)
				}

			}

		}
	}
}
