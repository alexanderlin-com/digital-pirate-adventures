import java.util.Scanner;
public class Ship {
    public static Scanner scan = new Scanner(System.in);

    public static String playerName;
    public static String shipName;

    public int health;
    public int bandwidth;
    public int speed;
    public int dodge;
    public double armor;
    public int accuracy;
    public double damage;
    public int broadsideType;
    public int swivelType;
    public int crewType;
    public int sailType;
    public int riggingType;
    public boolean swivel = true;


    public Ship(int health, int bandwidth, int speed, int dodge, double armor, int accuracy, double damage) {
        this.health = health;
        this.bandwidth = bandwidth;
        this.speed = speed;
        this.dodge = dodge;
        this.armor = armor;
        this.accuracy = accuracy;
        this.damage = damage;
    }

    public void setBroadside(int broadsideType)
    {
        System.out.println("Yer broadside cannons be yer main attack. What type of cannon be yer choice? (1/2/3)\n");
        System.out.println("""
                1. DDOS Cannon - Devastatin' damage, but accuracy be lackin.
                2. DNS Leak Launchers - Medium damage n' accuracy; a safe choice
                3. Zero Day Exploit Guns - They be strikin' true with high accuracy, but lower damage.
                """);
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                this.broadsideType = 1;
                System.out.println("They may not be the most accurate, but our DDOS cannons will make quick work o' the enemy.\n");
                break;
            case "2":
                this.broadsideType = 2;
                System.out.println("Keepin' everything balanced, I see. The best o' both worlds.\n");
                break;
            case "3":
                this.broadsideType = 3;
                System.out.println("No room for surprises, eh? Sometimes reliability be a the best choice.\n");
                break;
            default:
                System.out.println("Arr, I didn't quite catch that, matey. Input (1/2/3/) to choose yer cannon type.\n");
                setBroadside(broadsideType);
        }
    }

    public void setSwivel(int swivelType)
    {
        System.out.println("Yer swivel guns be a fine secondary gun. Smaller, but more precise. What be yer choice o' swivel? (1/2/3)\n");
        System.out.println("""
                1. Zip Bomb Swivels - Load yer swivel guns with zip bombs to confuse yer enemy and lower their accuracy./n
                2. IP-Trace Swivels - Load em' with shots that trace yer enemies' digital footprints, boostin' yer accuracy.
                3. Logic Bomb Swivels - Throw yer caution to the sea, unleashing pure RNG chaos and causin' unpredictable damage with every shot!
                """);
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                this.swivelType = 1;
                System.out.println("Any enemy we be facin' will have a hard time hittin' us once they get a taste o' these guns!\n");
                break;
            case "2":
                this.swivelType = 2;
                System.out.println("We'll maximize our ship's accuracy with these trusty guns!\n");
                break;
            case "3":
                this.swivelType = 3;
                System.out.println("Hehehe. Leavin' it all to chance, I see. If it's chaos ye wanted, it's chaos ye got.\n");
                break;
            default:
                System.out.println("Arr, I didn't quite catch that, matey. Input (1/2/3) to choose yer swivel gun ammo.\n");
                setSwivel(swivelType);
        }
    }

    public void setCrew(int crewType)
    {
        System.out.println("Our ship be havin' room for extra crew abilities. What be yer choice for em', matey? (1/2)\n");
        System.out.println("""
                1. Latency Musket Volley - Rally yer crew to fire a musket volley at yer enemy, dealin' damage and disruptin' their systems.
                2. Overclock Powder Kegs - Push yer crew to the limit, increasin' yer damage while sacrificin' some o' yer armor.
                """);
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                System.out.println("A solid choice, matey! we'll be sowin' chaos among the enemy with this crew maneuver!\n");
                this.crewType = 1;
                break;
            case "2":
                System.out.println("We won't be needin' defense if we blow our enemies out o' the water before they hit us!\n");
                this.crewType = 2;
                break;
            default:
                System.out.println("Arr, I didn't quite catch that, matey. Input (1/2) to choose yer crew ability.\n");
                setCrew(crewType);
        }
    }

    public void setSails(int sailType, int dodge, int speed)
    {
        System.out.println("Our ship's sails be the lifeblood o' our vessel. What kind o' sails be on yer ship? (1/2)\n");
        System.out.println("""
                1. VPN Sails - Gives our ship extra maneuverability while sacrificin' some speed, navigatin' the waves with stealth.
                2. Fiber Optic Sails - Harness the speed of fiber optics, lettin' yer ship race through the waters
                """);
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                System.out.println("Our ship might be a little slower, but it be more nimble on the waves!\n");
                this.sailType = 1;
                this.dodge = dodge + 2;
                this.speed = speed - 2;
                break;
            case "2":
                System.out.println("Speed be our priority! We be wastin' no time on our adventure!\n");
                this.sailType = 2;
                this.speed = speed + 2;
                break;
            default:
                System.out.println("Arr, I didn't quite catch that, matey. Input (1/2) to choose yer ship's sails.\n");
                setSails(sailType, dodge, speed);
        }

    }

    public void setRigging(int riggingType, double armor, int speed)
    {
        System.out.println("Our ship not be complete without proper rigging. What kind o' rigging be on yer ship? (1/2)\n");
        System.out.println("""
                1. Anti-virus Rigging - Fortify yer defenses, makin yer ship tougher while sacrificin' some speed
                2. Tune-up Utility Rigging - Keep yer ship in top condition, allowing it to perform at its maximum speed.
                """);
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                this.riggingType = 1;
                this.armor += 0.1;
                this.speed = speed - 2;
                System.out.println("Our defenses be top o' the line with this anti-virus!\n");
                break;
            case "2":
                this.riggingType = 2;
                this.speed = speed + 2;
                System.out.println("We be keepin' our ship running as fast as possible!\n");
                break;
            default:
                System.out.println("Arr, I didn't quite catch that, matey. Input (1/2) to choose yer ship's riggin'\n");
                setRigging(riggingType, armor, speed);
        }
    }
    public void customize()
    {

    }

    public void takeDamage(int damage)
    {
        this.health -= damage;
    }
    public void useBandwidth(int burn)
    {
        this.bandwidth -= burn;
    }
    public void nerfAccuracy(int nerf)
    {
        this.accuracy -= nerf;
    }
    public void nerfDodge(int nerf)
    {
        this.dodge -= nerf;
    }
    public void nerfArmor(double nerf)
    {
        this.armor -= nerf;
    }


    public void turn()
    {

    }

    public void healthCheck() {
        if (health == 0) {
            System.out.println("\n\nYou are dead.\n\n*RE4 Leon death sound*");
            System.exit(0);

        }
    }

    public void bandwidthCheck()    {
        if(bandwidth == 0)
        {
            System.out.println("We've stretched our bandwidth thin! Our ship be takin' damage!");
            takeDamage(Main.d20());
        }
    }

}
