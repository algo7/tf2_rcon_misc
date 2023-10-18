package ui

import (
	"strconv"

	"github.com/algo7/tf2_rcon_misc/utils"
	"github.com/rivo/tview"
)

var app *tview.Application
var playerTable *tview.Table

// InitUI initializes the UI
func InitUI() {
	app = tview.NewApplication()
	playerTable = tview.NewTable().SetBorders(true)
}

func App() *tview.Application {
	return app
}

func PlayerTable() *tview.Table {
	return playerTable
}

// UpdateView updates the UI
func UpdateView(playersInGame []*utils.PlayerInfo) {
	app.QueueUpdateDraw(func() {
		// Clear the table first
		playerTable.Clear()

		// Adding headers
		playerTable.SetCell(0, 0, tview.NewTableCell("SteamID"))
		playerTable.SetCell(0, 1, tview.NewTableCell("Name"))
		playerTable.SetCell(0, 2, tview.NewTableCell("UserID"))
		playerTable.SetCell(0, 3, tview.NewTableCell("SteamAccType"))
		playerTable.SetCell(0, 4, tview.NewTableCell("SteamUniverse"))

		// Adding data
		for i, player := range playersInGame {

			playerTable.SetCell(i+1, 0, tview.NewTableCell(strconv.FormatInt(player.SteamID, 10)))
			playerTable.SetCell(i+1, 1, tview.NewTableCell(player.Name))
			playerTable.SetCell(i+1, 2, tview.NewTableCell(strconv.Itoa(player.UserID)))
			playerTable.SetCell(i+1, 3, tview.NewTableCell(player.SteamAccType))
			playerTable.SetCell(i+1, 4, tview.NewTableCell(strconv.Itoa(player.SteamUniverse)))
		}
	})
}
