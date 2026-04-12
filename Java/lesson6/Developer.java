public class Developer extends Employee {    
    private String language;

    public Developer(String name, double salary, String language){
        super(name, salary);
        this.language = language;
    }

    public String getLanguage() {
        return language;
    }
    public void setLanguage(String language) {
        if (language != null && !language.isBlank()) {
            this.language = language;
        }
    }

    @Override
    public void info(){
        super.info();
        System.out.println("Role: Developer");
        System.out.println("Language: " + getLanguage());
    }
    @Override
    public void work(){
        System.out.println("Developer " + getName() + " writes code on " + getLanguage());
    }

    @Override
    public void bonus(){
        System.out.println("Developer " + getName() + " received a bonus for a good project");
    }
    
}
