public class InstallmentPayment implements PaymentMethod{
    private int months;
    private String bank;

    public InstallmentPayment(int months, String bank) {
        this.months = months;
        this.bank = bank;
    }

    public int getMonths() {
        return months;
    }
    public String getBank() {
        return bank;
    }

    @Override
    public void pay(double amount) {
        System.out.println("Installment payment: " + amount + " | Bank: " + bank + " (" + months + " months)");
    }
}
