public class Product {
    private static int nextId = 1;

    private int id;
    private String name;
    private int price;

    public Product(String name, int price){
        this.id = nextId;
        nextId++;

        this.name = name;
        if(price<0) price = 0;
        this.price = price;
    }

    public int getId(){
        return id;
    }
    public String getName(){
        return name;
    }
    public int getPrice(){
        return price;
    }
    public void print(){
        System.out.println("[" + id + "]" + name + " -> " + price);
    }

    
    public void setPrice(int price){
        if(price < 0) price = 0;
        this.price = price;
    }

}