package main

import (
	"fmt"
)

type Wallet struct{
	Cash int
}

func (w *Wallet) Pay(check int) int {
	w.Cash -= check
	return w.Cash
}

type CreditCard struct{
	Cash int
	CardId int
}

func (ccard *CreditCard) Pay(check int) int {
	fmt.Printf("Card Number: %d \n", ccard.CardId)
	ccard.Cash -= check
	return ccard.Cash
}

type Payment interface{
	Pay(int) int
}

func PayCheck(moneyHolder Payment, check int) {
	fmt.Printf("Оплата произошла через: %T \n", moneyHolder)
	balance := moneyHolder.Pay(check)
	fmt.Printf("Остаток по карте: %d \n", balance)
}

func main() {
	var wal = Wallet{2300}
	var card = &CreditCard{12000, 32354456}
	
	PayCheck(&wal, 499)
	PayCheck(card, 1200)
}