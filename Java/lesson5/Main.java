import java.util.Scanner;

public class Main {
      public static void main(String[] args){

        Scanner sc = new Scanner(System.in);

        Shop shop = new Shop(10);

        while (true){
                System.out.println();
                System.out.println("1) Add product");
                System.out.println("2) Show catalog");
                System.out.println("3) Find by id");
                System.out.println("4) Remove by id");
                System.out.println("0) Exit");
                System.out.print("Choose: ");

                int choice = sc.nextInt();
                sc.nextLine();

                if (choice == 0) break;

                if(choice == 1){
                    System.out.print("Name: ");
                    String name = sc.nextLine();
                    System.out.print("Price: ");
                    int price = sc.nextInt();

                    shop.addProduct(name, price);
                }
                else if(choice == 2){
                    shop.printAll(); 
                }
                else if (choice == 3){
                    System.out.print("Id: ");
                    int id = sc.nextInt();
                    sc.nextLine();

                    Product p = shop.findById(id);

                    if(p == null){
                        System.out.println("Not found :(");
                    }
                    else{
                        p.print();
                    }
                    
                }
                else if (choice == 4){
                    System.out.print("Id: ");
                    int id = sc.nextInt();
                    sc.nextLine();

                    boolean removed = shop.removeById(id);

                    if (removed){
                        System.out.println("Product removed!");
                    } else {
                        System.out.println("Product not found!");
                    }
                    
                }
                else{
                    System.out.println("Wrong option");
                }
        }
        System.out.println("Bye");
        sc.close();
    }
}
