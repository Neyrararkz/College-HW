public class KaspiTransferPayment implements PaymentMethod {
    private String receiverPhone;
    private String receiverName;

    public KaspiTransferPayment(String receiverPhone, String receiverName) {
        this.receiverPhone = receiverPhone;
        this.receiverName = receiverName;
    }

    public String getRecieverPhone() {
        return receiverPhone;
    }
    public String getRecieverName() {
        return receiverName;
    }

    @Override
    public void pay(double amount) {
        System.out.println("Kaspi transfer payment: " + amount + " | " + receiverName + " (" + receiverPhone + ")");
    }
}