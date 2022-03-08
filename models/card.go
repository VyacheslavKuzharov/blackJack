package models

var CardNum = [14]string{"", "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var CardSuite = map[string]string{"spades": "♠", "hearts": "♥", "diamonds": "♦", "clubs": "♣"}

type Card struct {
	Num int
	Suite string
}