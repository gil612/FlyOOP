# Flight Booking Process
This code simulates flight booking through frequent flyer program. There are 2 categories: Gold and Silver Passenger.

This code contains method translations from the Java code, and aims to implement OO concepts like polymorphism, inheritance and substitutability as much as possible.
The process: A traveler must first register, either as Gold or Silver or just as "normal" passenger with an ID, with certain amount of money, or he can add later with a function _addToBankAccount_.
For the ID assignment: Gold members passnegers start with 1#, silver members passengers start with 2# and normal passengers start with 3#. Every ID is inserted into a list, after checking if this ID already exists.

```
func (p *Passenger) newPassenger(id int, bankAccount int) {
	if isValid(id) {
		id_list.PushBack(id)
		p.id = id
		p.bankAccount = bankAccount
	} else {
		fmt.Printf("ID %d already exists\n", id)
	}
}
```

To start a new booking process, we must use the function _newTravel_. This function calls the functions of the interface _flightTicket_, which are _bookATicket_ and _getLoungeVoucher_.

### Booking policy
When booking a ticket, every passenger’s bank account is being checked if there is enough money to cover the ticket price. Then the amount is redeemed from the bank account.

### Miles earnings policy
The miles that are added to the passenger's data is the price of the ticket.
Gold and silver members benefit a factor 0.5 and 0.25 respectively. These are set with the global variables goldfac and silverfac.


### Lounge access policy
Gold Passengers get free access.
Silver passengers get an access for $20, if there is a sufficient amount in the bank account.



We start with a simple example to test the code.
A passenger is registering as ID=31 and adds money to the bank account

``` 
pass := Passenger{} 
pass.newPassenger(31, 0) 
addToBankAccount(&pass, 98)
``` 



Output:
``` 
 Deposit: 0 + 98 = $98 
``` 
start a new travel:
``` 
 newTravel(&pass, 40)
``` 

Output:

``` 
Passenger #31 booked a ticket.
Price: $40
Miles: 40.00
Current in account (after transaction): $58
``` 

Then we try to book an expensive ticket for the passenger
``` 
newTravel(&pass, 60)
``` 

Output:
``` 
Booking for Passenger #31 has failed.
Ticket Price: $60 
Current: $58
``` 
Every registration of a passenger is done in the (p *Passenger)newPassenger(int,int) function. Gold and Silver Registrations are being directed to this function.
moving on to gold passenger #11. This time when registering the bank account is set to $98. Then he purchases a $40 ticket. He also gets a free lounge access.
``` 
gp := GoldPassenger{}
gp.newPassenger(11, 98)
``` 
Miles earning for gold and silver members is done within two steps, first as a regular passenger and then according to their policy in their function.
So here 40 * (1 + 0.5) is 60. That means the gold passenger collected 60 miles.
Gold passengers also get a message that they have a free access to the lounge.

``` 
newTravel(&gp, 40)
``` 

Output:
``` 
Passenger #11 booked a ticket.
Price: $40
Miles: 60.00
Current in account (after transaction): $58

Passenger #11 gets a free lounge access
``` 


Let’s try to create a silver passenger #21
``` 
sil := SilverPassenger{}
sil.newPassenger(21, 0)
addToBankAccount(&sil, 70)
newTravel(&sil, 20)
``` 
A silver passenger gets access to the lounge with a $20 fee. 
``` 
Passenger #21 booked a ticket.
Price: $20
Miles: 25.00
Current in account (after transaction): $50

For entering the lounge passenger #21 paid $20. Current in account: $30
``` 



When we book another ticket, this passenger doesn't get any notice about a lounge

``` 
newTravel(&sil, 20)
``` 
Output
``` 
Passenger #21 booked a ticket.
Price: $20
Miles: 50.00
Current in account (after transaction): $10

For entering the lounge passenger #21 paid $20. Current in account: $10
``` 
Then we add $25 to the bank account and try to purchase a $40 ticket

``` 
addToBankAccount(&sil, 25)
newTravel(&sil, 40)
``` 
Output:
``` 
Deposit of passenger #21: 10 + 25 = $35

Booking for Passenger #21 has failed.
Ticket Price: $40 
Current: $35

For entering the lounge passenger #21 paid $20. Current in account: $15
``` 

The passenger gets a message that he failed to book a ticket because of insufficient amount but still got charged for $20 to get access to the lounge.
This problem occurred again with the gold passenger, but luckily, he wasn’t charged 🤓:

``` 
gp2 := GoldPassenger{}
gp2.newPassenger(12, 20)
newTravel(&gp2, 50)
```
Output:
``` 
Booking for Passenger #12 has failed.
Ticket Price: $50 
Current: $20

Passenger #12 gets a free lounge access
``` 
### Insights about the code
Unfortunately, gold and silver passengers have an additional and unnecessary check for their ticket price and bank account. This is because they are being referenced to _func (p *Passenger) bookATicket(price int)_.

Implementing the Liskov substitution principle was done almost successfully. There is a call to get lounge voucher for normal passengers - which is our super class. But since according to the policy he's not allowed to enter the lounge, he doesn't get any message. This is done with an empty function:

```
func (p *Passenger) getLoungeVoucher() {}
```

when we register a passenger with an ID that already exists, The assigned ID is 0. But the object referenced to this passenger is different from the other passenger’s object with the same ID. There should be an error handling of this case.


```
pass2 := Passenger{}
pass2.newPassenger(31, 0) // #31 already exists
addToBankAccount(&pass2, 98)
newTravel(&pass2, 40)
```
Output:

```  
ID 31 already exists
Deposit of passenger #0: 0 + 98 = $98

Passenger #0 booked a ticket.
Price: $40
Miles: 40.00
Current in account (after transaction): $58

```

In the case of adding cash to bank account with the function addToBankAccount, because this a secure and simple operation for all passengers, I thought it would be efficient to write all in one function without any referencing like in other methods.
A multiple receiver for all different passengers in one function is not possible. Therefore I used the empty interface{}.
It was too much repetitiveness like this case of gold and silver passengers:

```
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
```



Overall, it was possible to implement OO principles with GO.
