public class Designer extends Employee {    
    private String tool;

    public Designer(String name, double salary, String tool){
        super(name, salary);
        this.tool = tool;
    }

    public String getTool() {
        return tool;
    }
    public void setTool(String tool) {
        if (tool != null && !tool.isBlank()) {
            this.tool = tool;
        }
    }

    @Override
    public void info(){
        super.info();
        System.out.println("Role: Designer");
        System.out.println("Tool: " + getTool());
    }
    @Override
    public void work(){
        System.out.println("Designer " + getName() + " creates a design in " + getTool());
    }

    @Override
    public void bonus(){
        System.out.println("Designer " + getName() + "received a bonus for good design");
    }
    
}