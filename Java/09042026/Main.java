public class Main {
    public static void main(String[] args) {
        PaymentService service = new PaymentService();
        
        PaymentMethod[] payments = {
            new CardPayment("1234 5678 9012 3456", "Ali"),
            new CashPayment("KZT"),
            new MobilePayment("+77001234567"),
            new KaspiTransferPayment("+77005556677", "Diana"),
            new KaspiQRPayment("QR123456", "Coffee Shop"),
            new InstallmentPayment(12, "Kaspi")
        };

        double amount = 10000;

        for (PaymentMethod payment : payments) {
            service.processPayment(payment, amount);
            System.out.println("---------");
        }
    }
}
