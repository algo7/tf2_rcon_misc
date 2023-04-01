package utils

import "fmt"

// Check if supplied argument *in* is a chatline, if so, return: <true>, <the player that said it>, <what did he say>
// Game specific for dystopia
func GetChatSayDystopia(players []string, in string) (bool, string, string) {

	for _, player := range players {

		// check if we found a player saying that in our playerlist
		if len(in) > len(player)+5 && in[1:len(player)+1] == player && in[len(player)+1:len(player)+2] == ":" {
			fmt.Printf("CHAT: [%s] %s\n", player, in[len(player)+3:])
			return true, TrimCommon(player), TrimCommon(in[len(player)+3:])
		}
	}

	return false, "", ""
}
