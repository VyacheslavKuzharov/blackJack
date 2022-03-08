package interaction

import (
	"fmt"
	"github.com/VyacheslavKuzharov/blackJack/models"
)

func PrintGreeting(playerName string)  {
	fmt.Println("BLACK JACK")
	fmt.Println("Starting a new game...")
	fmt.Printf("Good luck, %v!", playerName)
}

func ShowRoundSummary(bankAmount int32, round int)  {
	addMoreSpace(2)

	fmt.Printf("----------Round № %v ----------", round)
	addMoreSpace(1)
	fmt.Printf("----------Bank: %v$ ----------", bankAmount)
	addMoreSpace(1)
}

func addMoreSpace(n int)  {
	for i := 0; i < n; i++ {
		fmt.Println("")
	}
}

func ShowDealerInfo(dialer *models.Player)  {
	addMoreSpace(2)

	fmt.Println("Dialer")
	fmt.Printf("Amount: %v", dialer.Amount)
	addMoreSpace(1)
	fmt.Printf("Cards: %v", dialer.Amount)
	addMoreSpace(1)
	fmt.Printf("Points: %v", dialer.Amount)
	addMoreSpace(1)
}