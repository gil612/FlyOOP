package main

import (
	"container/list"
	"fmt"
)

// global variable of ID lists
var id_list = list.New()

// Miles factor for gold members is set to 0.5
var goldfac = 0.5

// Miles factor for silver members is set to 0.25
var silverfac = 0.25

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

type flightTicket interface {
	bookATicket(int)
	getLoungeVoucher()
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

func newTravel(f flightTicket, price int) {
	f.bookATicket(price)
	f.getLoungeVoucher()
}

func (p *Passenger) bookATicket(price int) {
	if p.bankAccount-price < 0 {
		errordisplay(p.id, p.bankAccount, price)
		return
	}
	p.bankAccount -= price
	p.miles += float64(price)
	fmt.Printf("Passenger #%d booked a ticket.\nPrice: $%d\nMiles: %.2f\nCurrent in account (after transaction): $%d\n\n", p.id, price, p.miles, p.bankAccount)
}

func (g *GoldPassenger) bookATicket(price int) {
	if g.pass.bankAccount-price < 0 {
		errordisplay(g.pass.id, g.pass.bankAccount, price)
		return
	}
	g.pass.miles += g.fac * float64(price)
	g.pass.bookATicket(price)
}

func (s *SilverPassenger) bookATicket(price int) {
	if s.pass.bankAccount-price < 0 {
		errordisplay(s.pass.id, s.pass.bankAccount, price)
		return
	}
	s.pass.miles += s.fac * float64(price)
	s.pass.bookATicket(price)
}

func (p *Passenger) getLoungeVoucher() {}
func (g *GoldPassenger) getLoungeVoucher() {
	fmt.Printf("Passenger #%d gets a free lounge access\n", g.pass.id)
}
func (s *SilverPassenger) getLoungeVoucher() {
	if s.pass.bankAccount-20 >= 0 {
		s.pass.bankAccount -= 20
		fmt.Printf("For entering the lounge passenger #%d paid $20. Current in account: $%d\n\n", s.pass.id, s.pass.bankAccount)
	}
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

// Cunstructor for the GoldPassenger type
// @fac is set to 0.5 (assigned from the global variable @goldfac)
// receiver of a GoldPassenger reference
func (g *GoldPassenger) newPassenger(id int, bankAccount int) {
	g.fac = goldfac
	g.pass.newPassenger(id, bankAccount)
}

// Cunstructor for the SilverPassenger type
// @fac is set to 0.25 (assigned from the global variable @silverfac)
// receiver of a SilverPassenger reference
func (s *SilverPassenger) newPassenger(id int, bankAccount int) {
	s.fac = silverfac
	s.pass.newPassenger(id, bankAccount)

}

// In this method we add cash to the bank account.
// Type of a passenger must be identified
func addToBankAccount(anything interface{}, amount int) {

	t1 := 0
	t2 := 0
	var id int

	switch v := anything.(type) {
	case *GoldPassenger:
		t1 = v.pass.bankAccount
		v.pass.bankAccount += amount
		t2 = v.pass.bankAccount
		id = v.pass.id
	case *SilverPassenger:
		t1 = v.pass.bankAccount
		v.pass.bankAccount += amount
		t2 = v.pass.bankAccount
		id = v.pass.id
	case *Passenger:
		t1 = v.bankAccount
		v.bankAccount += amount
		t2 = v.bankAccount
		id = v.id
	}

	fmt.Printf("Deposit of passenger #%d: %d + %d = $%d\n\n", id, t1, amount, t2)
}

// The function displays a message when there is insufficient amount to purchase a ticket
func errordisplay(id int, baccount int, price int) {
	fmt.Printf("Booking for Passenger #%d has failed.\nTicket Price: $%d \nCurrent: $%d\n\n", id, price, baccount)
}

// Demo
// Id of a gold passenger starts with 1
// Id of a silver Passenger starts with 2
// Id of a normal Passenger starts with 3

func main() {

	// normal passenger
	pass := Passenger{}
	pass.newPassenger(31, 0)
	addToBankAccount(&pass, 98)
	// bank account after purchase: &58
	newTravel(&pass, 40)
	// booking is not possible
	newTravel(&pass, 60)

	//  gold member registration
	gp := GoldPassenger{}
	gp.newPassenger(11, 98)
	// bank account after purchase: &58
	// gets a free lounge access
	newTravel(&gp, 40)

	// silver member registration
	sil := SilverPassenger{}
	sil.newPassenger(21, 0)
	addToBankAccount(&sil, 70)
	// bank account after purchase + lounge access's fee: $30
	newTravel(&sil, 20)
	// bank account after purchase: $10
	// No lounge access
	newTravel(&sil, 20)
	// bank account: $35
	addToBankAccount(&sil, 25)
	// booking is not possible, but still charged for the lounge
	newTravel(&sil, 40)

	gp2 := GoldPassenger{}
	// new gold passenger with $20 at registration
	gp2.newPassenger(12, 20)
	// booking has failed, but gets a message with a free access to the lounge
	newTravel(&gp2, 50)

	pass2 := Passenger{}
	// #31 already exists
	pass2.newPassenger(31, 0)
	// pass2 behave like an independent object
	addToBankAccount(&pass2, 98)
	newTravel(&pass2, 40)
}
