package main

import (
	"github.com/VyacheslavKuzharov/blackJack/interaction"
	"github.com/VyacheslavKuzharov/blackJack/models"
)

var currentRound = 0
var playerName string
var player *models.Player
var dialer *models.Player
var deck *models.Deck
var playerHand *models.Hand
var dealerHand *models.Hand
var bank = map[string]int32{
	"player": 0,
	"dialer": 0,
}

func main() {
	startGame()

	winner := "" // "Player" || "Dealer" || ""

	for winner == "" {
		startRound()
	}

	endGame()

}

func startGame() {
	playerName = interaction.GetPlayerName()
	interaction.PrintGreeting(playerName)

	player = models.NewPlayer(playerName)
	dialer = models.NewPlayer("Dealer")
	deck = models.NewDeck()
}

func startRound() {
	currentRound++

	dealCards()
	bankAmount := playersBet()
	interaction.ShowRoundSummary(bankAmount, currentRound)
	interaction.ShowDealerInfo(dialer)
	//interaction.ShowPlayersInfo(player)
}

func endGame() {}

func dealCards() {
	playerDraw := deck.Draw(2)
	dealerDraw := deck.Draw(2)

	playerHand.Create(playerDraw)
	dealerHand.Create(dealerDraw)
}

func playersBet() int32 {
	bank["player"] = player.Bet(10)
	bank["dialer"] = dialer.Bet(10)

	return bank["player"] + bank["dialer"]
}
