public class Employee {
    private String name;
    private double salary;

    public Employee(String name, double salary){
        this.name = name;
        this.salary = salary;
    }

    public String getName() {
        return name;
    }    
    public double getSalary() {
        return salary;
    }

    public void setName(String name){
        if (name != null && !name.isBlank()){
            this.name = name;
        }
    }
    public void setSalary(double salary){
        if (salary >= 0){
            this.salary = salary;
        }
    }

    public void info(){
        System.out.println("Employee: " + getName());
        System.out.println("Salary: " + getSalary());
    }
    public void work(){
        System.out.println("Employee " + getName() + " works");
    }

    public void salaryIncrease(double amount) {
        if (amount > 0) {
            this.salary = salary + amount;
            System.out.println("Employee " + getName() + " received a salary increase");
            System.out.println("Salary: " + getSalary());
        }
    }
    public void bonus() {
        System.out.println("Employee " + getName() + " received a bonus for good work");
    }
}
