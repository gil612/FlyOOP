package main

import (
	"fmt"
	"reflect"
)

type Passenger struct {
	id          int
	bankAccount int
	miles       int
}

type GoldPassenger struct {
	fac  int
	gold Passenger
}

func (p *Passenger) newPassenger(id int, bankAccount int) {
	p.id = id
	p.bankAccount = bankAccount
}

func (p *Passenger) addToBankAccount(amount int) {
	p.bankAccount += amount
}

func (p *Passenger) bookATicket(price int) {
	p.bankAccount -= price

	if reflect.TypeOf(str1).Name() == "string" {
		fmt.Printf("true\n")
	}

	if reflect.TypeOf(*p).Name() == "Passenger" {
		fmt.Printf("ptrue\n")
	}

	if reflect.TypeOf(p).Name() == "GoldPassenger" {
		fmt.Printf("gptrue\n")
	}

}

func main() {
	gp3 := CreateGoldPassenger(2, 4, 100)
	gp := GoldPassenger{
		fac: 2,
		gold: Passenger{
			id:          4,
			bankAccount: 1000,
		},
	}

	gp2 := &GoldPassenger{
		fac: 2,
		gold: Passenger{
			id:          5,
			bankAccount: 1000,
		},
	}
	//  pass := new(Passenger)
	// pass.newPassenger(1, 51, 1)

	// pass.addToBankAccount(98)
	// fmt.Println(pass.bankAccount)
	// pass.bookATicket(40)

	// gp := &GoldPassenger{
	// 	Passenger{
	// 		id:          3,
	// 		bankAccount: 650,
	// 		fac:         5,
	// 	},
	// }

	// fmt.Println(gp.Base.bankAccount)
	// gp.Base.addToBankAccount(98)
	// gp.Base.bookATicket(500)
	gp2.bookATicket(51)
	fmt.Println(gp.fac)
	fmt.Println(gp2.fac)
}
