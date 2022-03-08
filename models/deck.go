package models

import (
	"math/rand"
	"time"
)

var suits = [4]string{"spades", "hearts", "diamonds", "clubs"}
var newDeck []Card
var cardsPlayed []Card

type Deck struct {
	Cards       []Card
	CardsPlayed []Card
}

func (d *Deck) Draw(size int) []Card {
	var sampleCards []Card

	for i := 0; i < size; i++ {
		randomCard, cardIndex := pickRandomCard(d.Cards)
		reducedDeck := removeCardFromDeck(d.Cards, cardIndex)

		cardsPlayed = append(cardsPlayed, randomCard)
		d.Cards = reducedDeck
		d.CardsPlayed = cardsPlayed

		sampleCards = append(sampleCards, randomCard)
	}

	return sampleCards
}

func NewDeck() *Deck {
	for index := range CardNum {
		if index == 0 {
			continue
		}

		for _, suite := range suits {
			newCard := Card{index, suite}
			newDeck = append(newDeck, newCard)
		}
	}

	return &Deck{Cards: newDeck, CardsPlayed: cardsPlayed}
}

func pickRandomCard(cards []Card) (card Card, index int) {
	randCardIndex := genRandomIndex(cards)
	card = cards[randCardIndex]

	return card, randCardIndex
}

func genRandomIndex(cards []Card) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	return randGenerator.Intn(len(cards))
}

func removeCardFromDeck(cards []Card, i int) []Card {
	cards[i] = cards[len(cards)-1]
	return cards[:len(cards)-1]
}
