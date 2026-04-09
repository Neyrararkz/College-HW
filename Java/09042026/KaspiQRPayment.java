public class KaspiQRPayment implements PaymentMethod {
    private String QRcode;
    private String shopName;

    public KaspiQRPayment(String QRcode, String shopName) {
        this.QRcode = QRcode;
        this.shopName = shopName;
    }

    public String getQRcode() {
        return QRcode;
    }
    public String getShopName() {
        return shopName;
    }

    @Override
    public void pay(double amount) {
        System.out.println("Kaspi QR payment: " + " | " + shopName + " (" + QRcode + ")");
    }
}