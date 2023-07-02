// This code simulates a flight booking through frquent flyer program.
// There are 2 programms: Gold and Silver Passenger. They benefit a higher miles factor when they book a ticket.
// A traveler must first register, either as Gold or Silver or just as "normal" passenger.
// The data of every Gold and Silver passenger is kept as a Passenger type, which contains an id and a current bank account. Current miles are set to 0.

// In every purchase the price of the ticket is redeemed from the traveler's bank account.

// This code contains method translations from the Java code, and aims to implement OOP features as much as possible

package main

import (
	"container/list"
	"fmt"
)

// global variable of ID lists
var id_list = list.New()

// Checks if the given ID already exist
// returns true if valid. otherwise returns false
func isValid(id int) bool {
	for i := id_list.Front(); i != nil; i = i.Next() {

		if i.Value == id {
			return false
		}
	}
	return true

}

type flightticket interface {
	newPassenger(int, int)
}

type Passenger struct {
	id          int
	bankAccount int
	miles       float64
}

type GoldPassenger struct {
	fac  float64
	pass Passenger
}

type SilverPassenger struct {
	fac  float64
	pass Passenger
}

// Cunstructor for the Passenger type
// receiver of a Passenegr reference
func (p *Passenger) newPassenger(id int, bankAccount int) {
	if isValid(id) {
		id_list.PushBack(id)
		p.id = id
		p.bankAccount = bankAccount
	} else {
		fmt.Printf("ID %d already exists\n", id)
	}

}

// Cunstructor for the Gold passenger type, fac is set to 1.5
// receiver of a GoldPassenger reference
func (g *GoldPassenger) newPassenger(id int, bankAccount int) {
	g.fac = 1.5
	g.pass.newPassenger(id, bankAccount)

}

// Cunstructor for the Silver passenger type, fac is set to 1.25
// receiver of a SilverPassenger reference
func (s *SilverPassenger) newPassenger(id int, bankAccount int) {
	s.fac = 1.25
	s.pass.newPassenger(id, bankAccount)

}

// In this method we add cash to the bank account.
// Type of a passenger must be identified
func addToBankAccount(anything interface{}, amount int) {

	t1 := 0
	t2 := 0

	switch v := anything.(type) {
	case *GoldPassenger:
		t1 = v.pass.bankAccount
		v.pass.bankAccount += amount
		t2 = v.pass.bankAccount
	case *SilverPassenger:
		t1 = v.pass.bankAccount
		v.pass.bankAccount += amount
		t2 = v.pass.bankAccount
	case *Passenger:
		t1 = v.bankAccount
		v.bankAccount += amount
		t2 = v.bankAccount

	}

	fmt.Printf("Deposit: %d + %d = %d\n\n", t1, amount, t2)
}

// Through interface{} the passneger's type is identified
// The price is redeemed from the bank account
// Every Passenger earns miles, equal to the price. Gold and Silver enjoy a factor of 1.5 and 1.25 respectively.
func bookATicket(anything interface{}, price int) {
	switch v := anything.(type) {
	case *GoldPassenger:
		if v.pass.bankAccount-price >= 0 {
			v.pass.bankAccount -= price
			v.pass.miles += v.fac * float64(price)
			fmt.Printf("Gold passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.pass.id, price, v.pass.miles, v.pass.bankAccount)
		} else {
			errordisplay(v.pass.id)
		}

	case *SilverPassenger:
		if v.pass.bankAccount-price >= 0 {
			v.pass.bankAccount -= price
			v.pass.miles += v.fac * float64(price)
			fmt.Printf("Silver passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.pass.id, price, v.pass.miles, v.pass.bankAccount)
		} else {
			errordisplay(v.pass.id)
		}
	case *Passenger:
		if v.bankAccount-price >= 0 {
			v.bankAccount -= price
			v.miles += float64(price)
			fmt.Printf("Passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.id, price, v.miles, v.bankAccount)
		} else {
			errordisplay(v.id)
		}

	}
}

// The function displays a message when there is insufficient amount to purchase a ticket
func errordisplay(id int) {
	fmt.Printf("Booking for Passenger #%d has failed.\nMoney was not taken from your account.\nPlease try again later.\n\n", id)
}

// Demo
// Id of a gold passenger is set to 1#
// Id of a silver Passenger is set to 2#
// Id of a normal Passenger is set to 3#

func main() {

	pass := Passenger{}
	pass.newPassenger(31, 0)

	addToBankAccount(&pass, 98)
	bookATicket(&pass, 40)

	pass2 := Passenger{}
	pass2.newPassenger(31, 0) // #31 already exists

	addToBankAccount(&pass2, 98)
	bookATicket(&pass2, 40)

	gp := GoldPassenger{}
	gp.newPassenger(11, 0)
	addToBankAccount(&gp, 60)
	bookATicket(&gp, 30)

	gp2 := GoldPassenger{}
	gp2.newPassenger(12, 0)
	addToBankAccount(&gp2, 100)
	bookATicket(&gp2, 256)
	bookATicket(&gp2, 10)

	gp3 := GoldPassenger{}
	gp3.newPassenger(33, 0)
	bookATicket(&gp3, 60)

	sil := SilverPassenger{}
	sil.newPassenger(21, 0)
	addToBankAccount(&sil, 60)
	bookATicket(&sil, 40)

	addToBankAccount(&gp3, 600)
	bookATicket(&gp3, 456)

	bookATicket(&gp2, 20)
}

// Improvement Ideas:
// Should implement an exception when trying to create passenger with an Id that is already exists.
