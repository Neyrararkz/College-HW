public class Shop {
    private Product[] products;
    private int size;

    public Shop(int capacity){
        if(capacity < 1) capacity =1;
        products = new Product[capacity];
        size = 0;
    }

    public boolean addProduct(String name, int price){
        if(size >= products.length){
            System.out.println("Shop is full! Cannot add more.");
            return false;
        }
        if (name == null || name.isEmpty()){
            System.out.println("Invalid name!");
            return false;
        }
        Product p = new Product(name,price);
        products[size] = p;
        size++;
        System.out.println("Added");
        p.print();
        return true;

    }
    public void printAll(){
        if (size == 0){
            System.out.println("Catalog is empty!");
            return;
        }
        System.out.println("=== CATALOG ===");

        for(int i = 0; i < size; i++){
            products[i].print();
        }
    }
    public Product findById(int id){
        for (int i = 0; i < size; i++){
            if(products[i].getId() == id){
                return products[i];
            }
        }
        return null;
    }
    public boolean removeById(int id){
        for(int i = 0; i < size; i++){
            if(products[i].getId() == id){
                for(int j = i; j < size - 1; j++){
                    products[j] = products[j + 1];
                }
                products[size - 1] = null;
                size--;
                return true;
            }
        }
        return false;
    }

    
    public boolean updatePriceById(int id, int newPrice){
        Product p = findById(id);
        if(p == null) return false;

        p.setPrice(newPrice);
        return true;
    }
    public int countProducts(){
        return size;
    }
    public void clear(){
        for(int i = 0; i < size; i++){
            products[i] = null;
        }
        size = 0;
    }
 
}