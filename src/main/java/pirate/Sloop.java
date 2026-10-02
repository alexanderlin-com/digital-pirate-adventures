public class Sloop extends Ship{


    public Sloop(int health, int bandwidth, int speed, int dodge, double armor, int accuracy, double damage) {
        super(200, 100, 10, 7, 0.1, 0, 0.75);
    }

    @Override
    public void customize()
    {
        System.out.println("A sloop! Prioritizin' speed and maneuverability, a fine choice, captain!");
        System.out.println("Now let's take a look at yer ship...\n");
        setBroadside(broadsideType);
        setSails(sailType,dodge,speed);
        setRigging(riggingType,armor,speed);
        this.swivel = false;

    }


    @Override
    public void turn()
    {
        System.out.println("turn");
    }




}
