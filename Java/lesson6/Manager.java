public class Manager extends Employee {    
    private int teamSize;

    public Manager(String name, double salary, int teamSize){
        super(name, salary);
        this.teamSize = teamSize;
    }

    public int getTeamSize() {
        return teamSize;
    }
    public void setTeamSize(int teamSize) {
        if (teamSize > 0) {
            this.teamSize = teamSize;
        }
    }

    @Override
    public void info(){
        super.info();
        System.out.println("Role: Manager");
        System.out.println("Team : " + getTeamSize() + " people");
    }
    @Override
    public void work(){
        System.out.println("Manager " + getName() + " manages a team of " + getTeamSize() + " people");
    }

    @Override
    public void bonus(){
        System.out.println("Manager " + getName() + " received a bonus for good team management");
    }
    
    
}