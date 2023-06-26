public class Main {
    public static void main (String[] args) throws Exception {
        System.out.println("hello");
        Passenger passenger = new Passenger(44, 0);
        passenger.addToBankAccount(400);
        passenger.bookATicket(passenger,200);
        passenger.bookATicket(passenger,123);
        System.out.println(passenger.getBankAccount());
        Passenger sil = new SilverPassenger(98,0);
        sil.addToBankAccount(400);
        passenger.bookATicket(sil,200);
        Passenger goldPassenger = new GoldPassenger(1, 500);
        goldPassenger.bookATicket(goldPassenger,1);

    }
}
