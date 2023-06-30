package main

import (
	"container/list"
	"fmt"
)

var id_list = list.New()

func isValid(id int) bool {
	for i := id_list.Front(); i != nil; i = i.Next() {

		if i.Value == id {
			return false
		}
	}
	return true

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

func (p *Passenger) newPassenger(id int, bankAccount int) {
	if isValid(id) {
		id_list.PushBack(id)
		p.id = id
		p.bankAccount = bankAccount
	} else {
		fmt.Printf("ID %d already exists\n", id)
	}

}

func (g *GoldPassenger) newPassenger(id int, bankAccount int) {
	g.fac = 1.5
	g.pass.newPassenger(id, bankAccount)

}
func (s *SilverPassenger) newPassenger(id int, bankAccount int) {
	s.fac = 1.25
	s.pass.newPassenger(id, bankAccount)

}

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

func bookATicket(anything interface{}, price int) {
	switch v := anything.(type) {
	case *GoldPassenger:
		if v.pass.bankAccount-price >= 0 {
			v.pass.bankAccount -= price
			v.pass.miles += v.fac * float64(price)
			fmt.Printf("Gold passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.pass.id, price, v.pass.miles, v.pass.bankAccount)
		} else {
			errordisplay()
		}

	case *SilverPassenger:
		if v.pass.bankAccount-price >= 0 {
			v.pass.bankAccount -= price
			v.pass.miles += v.fac * float64(price)
			fmt.Printf("Silver passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.pass.id, price, v.pass.miles, v.pass.bankAccount)
		} else {
			errordisplay()
		}
	case *Passenger:
		if v.bankAccount-price >= 0 {
			v.bankAccount -= price
			v.miles += float64(price)
			fmt.Printf("Passenger #%d booked a ticket.\nPrice: %d\nMiles %f\nCurrent in account: $%d\n\n", v.id, price, v.miles, v.bankAccount)
		} else {
			errordisplay()
		}

	}
}

func errordisplay() {
	fmt.Println("Booking has failed.\nMoney was not taken from your account.\nPlease try again later.\n")
}

func main() {

	pass := Passenger{}
	pass.newPassenger(11, 0)

	addToBankAccount(&pass, 98)
	bookATicket(&pass, 40)

	pass2 := Passenger{}
	pass2.newPassenger(11, 0)

	addToBankAccount(&pass2, 98)
	bookATicket(&pass2, 40)

	gp := GoldPassenger{}
	gp.newPassenger(31, 0)
	addToBankAccount(&gp, 60)
	bookATicket(&gp, 30)

	gp2 := GoldPassenger{}
	gp2.newPassenger(32, 0)
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
}
