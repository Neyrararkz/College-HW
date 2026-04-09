

public class CardPayment implements PaymentMethod{
    private String cardNumber;
    private String holderName;

    public CardPayment(String cardNumber, String holderName){
        this.cardNumber = cardNumber;
        this.holderName = holderName;
    }

    public String getCardNumber() {
        return cardNumber;
    }
    public String getHolderName() {
        return holderName;
    }

    @Override 
    public void pay(double amount)  {
        System.out.println("Payment by card: " + amount);
    }

}
