public class Frigate extends Ship{

    public Frigate(int health, int bandwidth, int speed, int dodge, double armor, int accuracy, double damage)
    {
        super(350,250,7,3,0.20,3,1);
    }

    public int broadsideType = 0;
    public int swivelType = 0;
    public int sailType = 0;
    public int riggingType = 0;




    @Override
    public void customize()
    {
        System.out.println("A frigate! Lookin' to balance out all aspects o' yer ship, a fine choice, captain!");
        System.out.println("Arr, now let's take a look at yer ship...");
        setBroadside(broadsideType);
        setSwivel(swivelType);
        setSails(sailType, dodge, speed);
        setRigging(riggingType, armor, speed);

    }

    @Override
    public void turn()
    {
        System.out.println("turn");
    }



}
