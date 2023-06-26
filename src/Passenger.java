public class Passenger implements FlightTicket {
    private int id;

    private int bankAccount;
    private double miles;

    public Passenger(int id, int bankAccount) {
        this.id = id;
        this.bankAccount = bankAccount;
        this.miles = 0;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getBankAccount() {
        return bankAccount;
    }

    public double getMiles() {
        return miles;
    }

    public void setMiles(double miles) {
        this.miles = miles;
    }

    public void setBankAccount(int bankAccount) {
        this.bankAccount = bankAccount;
    }

    public void addToBankAccount(int amount) {
        setBankAccount(getBankAccount() + amount);
    }


    @Override
    public void bookATicket(Passenger passenger, int price) {
        setBankAccount(getBankAccount() - price);
        if (passenger instanceof GoldPassenger) {
            setMiles(getMiles() + 1.5 * price);
        } else if (passenger instanceof SilverPassenger) {
            setMiles(getMiles() + 1.25*price);
        } else {
            setMiles(getMiles() + price);
        }
        System.out.println("Current Miles:" + getMiles());

    }


}
