public class MobilePayment implements PaymentMethod {
    private String phoneNumber;

    public MobilePayment(String phoneNumber){
        this.phoneNumber = phoneNumber;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    @Override
    public void pay(double amount) {
        System.out.println("Mobile payment: " + amount + " (" + phoneNumber + ")");
    }
}