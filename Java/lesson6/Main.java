public class Main {
      public static void main(String[] args){

        Developer dev = new Developer("Artem", 700000, "Python");
        Designer des = new Designer("Ramina", 500000, "Adobe Photoshop");
        Manager man = new Manager("Dinara", 1000000, 30);

        System.out.println("-----Employees-----");
        dev.info();
        System.out.println("-----");
        des.info();        
        System.out.println("-----");
        man.info();        

        System.out.println("----------");
        dev.work();
        System.out.println("-----");
        des.work();        
        System.out.println("-----");
        man.work();        

        System.out.println("----------");
        des.salaryIncrease(10000);        
        System.out.println("-----");
        des.info();
        
        System.out.println("----------");
        dev.bonus();
        System.out.println("-----");
        des.bonus();
        System.out.println("-----");
        man.bonus();

        
    }
}
