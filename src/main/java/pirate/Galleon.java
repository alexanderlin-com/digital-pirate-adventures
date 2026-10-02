public class Galleon extends Ship{

    public Galleon(int health, int bandwidth, int speed, int dodge, double armor,  int accuracy, double damage)
    {
        super(500,200,4,1,0.30,6,1.25);
    }

    public int broadsideType = 0;
    public int swivelType = 0;
    public int crewType = 0;
    public int sailType = 0;
    public int riggingType = 0;

    @Override
    public void customize()
    {
        System.out.println("Ah, a galleon! A mighty sailin' vessel for the toughest of fights, a fine choice, captain!");
        System.out.println("Arr, now let's take a look at yer ship...");
        setBroadside(broadsideType);
        setSwivel(swivelType);
        setCrew(crewType);
        setSails(sailType, dodge, speed);
        setRigging(riggingType, armor, speed);

    }

    @Override
    public void turn()
    {
        System.out.println("turn");
    }




}
