package utils

import "fmt"

// Check if supplied argument *in* is a chatline, if so, return: <true>, <the player that said it>, <what did he say>
// Game specific for dystopia
func GetChatSayDystopia(players []PlayerInfoCache, in string) (bool, string, string) {

	for _, player := range players {
		// check if we found a player saying that in our playerlist
		if len(in) > len(player.Name)+5 && in[1:len(player.Name)+1] == player.Name && in[len(player.Name)+1:len(player.Name)+2] == ":" {
			fmt.Printf("CHAT: [%s] %s\n", player.Name, in[len(player.Name)+3:])
			return true, TrimCommon(player.Name), TrimCommon(in[len(player.Name)+3:])
		}
	}

	return false, "", ""
}
