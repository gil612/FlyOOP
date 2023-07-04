# Flight Booking Process
This code simulates a flight booking through frequent flyer program. There are 2 programs: Gold and Silver Passenger. They benefit from a higher miles factor when they book a ticket. A traveler must first register, either as Gold or Silver or just as "normal" passenger. The data of every gold and silver passenger is kept as a Passenger type, which contains an id and a current bank account. Current miles are set to 0.
In every purchase the price of the ticket is redeemed from the traveler's bank account.
This code contains method translations from the Java code, and aims to implement OOP concepts like polymorphism, inheritance and substitutability as much as possible.
The process: A traveler must first register, either as Gold or Silver or just as "normal" passenger, with a certain n amount of money, or he can add later with a function "addToBankAccount"
To start a new booking process, we must use the function "newTravel". This function implements call functions of the interface flightticket, which are bookATicket and getLoungeVoucher.


### Booking policy
When booking a ticket, every passenger’s bank account is being checked if there is enough money to cover the ticket price. Then the amount is redeemed from the bank account.

### Miles earnings policy
The miles that are added to the passenger's data is the price of the ticket.
Gold and silver members benefit a factor 0.5 and o.25 respectively. These are set with the global variables goldfac and silverfac.


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
passenger #31 booked a ticket. 
Price: $40 
Miles: 40.00 
Current in account: $58 
``` 

Then we try to book an expensive ticket for the passenger
``` 
newTravel(&pass, 60)
``` 

Output:
``` 
Booking for Passenger #31 has failed. Money was not taken from your account. Please try again later.
``` 
Every registration of a passenger is done in the (p *Passenger)newPassenger(int,int) function. Gold and Silver Registrations are being directed to this function.
moving on to gold passenger #11. This time when registering the bank account is set to $98. And he purchases a $40 ticket. He also gets free lounge access.
``` 
gp := GoldPassenger{}
gp.newPassenger(11, 98)
``` 
Miles earning for gold and silver members is done within two steps, first as a regular passenger and then according to their policy in their function.
So here 40 * 1.5 is 60. That means the gold passenger collected 60 miles.
Gold passengers also get a message they have free access to the lounge.

``` 
newTravel(&gp, 40)
``` 

Output:
``` 
passenger #11 booked a ticket.
Price: $40
Miles: 60.00
Current in account: $58

You get free lounge access
``` 


Let’s try to create a silver passenger #21
``` 
sil := SilverPassenger{}
	sil.newPassenger(21, 0)
	addToBankAccount(&sil, 70)
``` 
A silver passenger gets access to the lounge with a $20 fee. 
``` 
 passenger #21 booked a ticket.
Price: $20
Miles: 25.00
Current in account: $50
``` 

For entering the lounge you paid $20. Current in account: $30

When we book another ticket, this passenger doesn't get any notice about a lounge
Output
``` 
passenger #21 booked a ticket.
Price: $20
Miles: 50.00
Current in account: $10
``` 
Then we add $25 to the bank account and try to purchase a $40 ticket

``` 
addToBankAccount(&sil, 25)
newTravel(&sil, 40)
``` 
Output:
``` 
Deposit: 10 + 25 = $35

Booking for Passenger #21 has failed. Money was not taken from your account. Please try again later.

For entering the lounge you paid $20. Current in account: $15
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
Booking for Passenger #12 has failed. Money was not taken from your account. Please try again later.

You get free lounge access
``` 
Insights about the code:
Implementing the Liskov Principle was done almost successfully. There is a call to get lounge voucher for normal passengers. But since according to the policy he's not allowed to enter the lounge, he doesn't get any message.
Another drawback is what happens when we register a passenger with an ID that already exists. The ID given is 0. But the object referenced to this passenger is different from the other passenger’s object with the same ID. There should be an error handling of this case. In GO there aren't any exeptions.


```
pass2 := Passenger{}
pass2.newPassenger(31, 0) // #31 already exists
addToBankAccount(&pass2, 98)
newTravel(&pass2, 40)
```
Output:

```  
ID 31 already exists
Deposit: 0 + 98 = $98

passenger #0 booked a ticket.
Price: $40
Miles: 40.00
Current in account: $58

```

In the case of adding cash to bank account with the function addToBankAccount, because this a secure and simple operation for all passengers, I thought it would be efficient to write all in one function without any referencing like in other methods.
A multiple receiver for all different passengers is not possible. Therefore I used the empty interface{}.
It was too much repetitiveness like this case of gold and silver passengers:

```
case *GoldPassenger:
    t1 = v.pass.bankAccount
    v.pass.bankAccount += amount
    t2 = v.pass.bankAccount
case *SilverPassenger:
    t1 = v.pass.bankAccount
    v.pass.bankAccount += amount
    t2 = v.pass.bankAccount
```
Overall, it was possible to implement OO principles with GO.
