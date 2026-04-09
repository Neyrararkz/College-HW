public class CashPayment implements PaymentMethod {
    private String currency;

    public CashPayment(String currency) {
        this.currency = currency;
    }

    public String getCurrency() {
        return currency;
    }

    @Override
    public void pay(double amount) {
        System.out.println("Cash payment: " + amount + "" + currency);
    }

}