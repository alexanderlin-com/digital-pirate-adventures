import java.util.Random;
import java.util.Scanner;

public class Storm {
    public static Scanner scan = new Scanner(System.in);
    public static Random roll = new Random();

    public static void intro(){
        Main.clearScreen();
        System.out.println("""
                Ye be sailin' onward, but a fearsome Bitstorm be upon us!\s

                We'll need to muster all our courage and wit to navigate this tempest,\s

                but I've no doubt we'll emerge triumphant on the other side, matey!
                """);
        Main.enterToContinue();
    }

    public static void lightning(Ship player){
        int score;
        System.out.println("""
                The skies 'round ye turn as dark as a coder's soul.\s

                A taste of metal fills yer mouth, and the scent o' ozone stings yer nose.\s

                The ship's deck quivers as metal objects begin to vibrate.
                """);
        Main.enterToContinue();
        System.out.println("\nA Data Surge Lightning Strike be on the horizon! What be yer course, captain? (1/2/3):\n");

        System.out.println("""
                1. Ready the VPN defenses! We'll shield the ship and face this lightning head-on!
                2. Quick as a flash, execute a maneuver to dodge the strike and keep us safe.
                3. Brave the storm and sail straight, though it risks harm to our ship, we'll not be swayed!
                """);

        String input = scan.nextLine();
        switch (input)
        {
            case "1":
                score = Main.d20();
                if (player.sailType == 1)
                {
                    score = score + 5;
                }
                if(score > 11)
                {
                    System.out.println("Our defenses held strong, and the lightning be naught but a spark! Onward we sail, unharmed and undeterred!");
                }
                else
                {
                    player.takeDamage(roll.nextInt(25));
                    System.out.println("Alas! The lightning strikes true, and we've taken damage, but our spirit be unbreakable! We press on, aye!");
                }
                break;
            case "2":
                score = Main.d20() + (player.dodge/2);
                if(score > 11)
                {
                    System.out.println("Aye, our ship be nimble as the tide and avoids the lightning! Onward we sail, unharmed and undeterred!");
                }
                else
                {
                    player.takeDamage(roll.nextInt(25));
                    System.out.println("Alas! The lightning strikes true, and we've taken damage, but our spirit be unbreakable! We press on, aye!");
                }
                break;
            case "3":
                    double damage = roll.nextDouble();
                    player.takeDamage((int) (damage - damage*(player.armor*2)));
                    System.out.println("Alas! The lightning strikes true, and we've taken damage! Fortunately, we've braced for our ship for the impact and most of it be negated!");

                break;
            default:
                player.takeDamage(roll.nextInt(25));
                System.out.println("""
                        Ye hesitated too long, and the storm's fury be takin' its toll.\s
                        Lightning strikes the ship, causin' damage!\s
                        Act swiftly, lest the tempest consumes us!
                        """);
                break;
        }
        Main.enterToContinue();
    }


    public static void whirlwind(Ship player){
        Main.clearScreen();
        System.out.println("""
                The lightning strike be but a memory, and they say lightning don't strike the same spot twice.\s
                Ye be prayin' that be true as ye forge ahead into the heart of the Bitstorm.
                """);
        Main.enterToContinue();
        System.out.println("All of a sudden, the winds be howlin', and a data debris whirlwind be upon us! What be our course, captain? (1/2/3)\n");

        System.out.println("""
                1. Ready the swivel gun! We'll give that whirlwind a taste o' digital lead to clear our path.
                2. Steady as she goes, we'll navigate with caution through the whirlwind, takin' it slow but riskin' delays.
                3. Put our brains to work and analyze the data patterns within the whirlwind for a potential shortcut.
                """);

        int score;
        String input = scan.nextLine();
        switch (input)
        {
            case "1":
                score = Main.d20();
                if(player.swivel)
                {
                    score += 5;
                }
                if(score > 12)
                {
                    if(player.swivel)
                    {
                        System.out.println("With a thunderous roar, we blasted the whirlwind with our trusty swivel gun, clearin' a path through the stormy chaos!");
                    }
                    else
                    {
                        System.out.println("Arrr, we ain't got a swivel gun, but by the grace of the sea gods, we dodged the whirlwind! Lucky we be!");
                    }
                }
                else {
                    player.takeDamage(roll.nextInt(25));
                    if(player.swivel)
                        System.out.println("Our swivel gun roared defiantly at the whirlwind, but alas, it weren't enough! The whirlwind struck back, dealin' damage to our brave vessel!");
                    else
                        System.out.println("Blast it all, we lack a swivel gun, and the whirlwind smites us, dealin' damage to our ship!");
                }
                break;
            case "2":
                score = Main.d20();
                if(score >=5)
                {
                    System.out.println("We skirted the whirlwind by takin' it slow and steady, but it cost us some precious bandwidth in the process.");
                    player.useBandwidth(roll.nextInt(25));
                }
                else
                {
                    System.out.println("We be takin' it slow, yet the whirlwind still found us, inflictin' damage upon our noble vessel!");
                    player.useBandwidth(roll.nextInt(25));
                    player.takeDamage(roll.nextInt(25));
                }
                break;
            case "3":
                score = Main.d20() + ((player.dodge+player.speed)/3);
                if(score >= 10)
                {
                    System.out.println("Avast! We spied a shortcut through the whirlwind, savvy navigatin' saved us!");
                }
                else
                {
                    System.out.println("We tried to cut through the tempestuous whirlwind, but still, it struck us, causin' damage to our ship! Arrr, cursed luck!");
                    player.takeDamage(roll.nextInt(25));
                }
                break;
            default:
                System.out.println("""
                        Yer moment o' hesitation be costly, matey!\s
                        The whirlwind's fury damages the ship!\s
                        Act swiftly, lest the tempest consumes us!
                        """);
                player.takeDamage(roll.nextInt(25));
                break;
        }
        Main.enterToContinue();
    }
    public static void eruption(Ship player){
        Main.clearScreen();
        System.out.println("With the whirlwind dissipated, a sense o' relief washes over ye, but ye know rest be not yet within reach! \n");
        Main.enterToContinue();
        System.out.println("From the depths of Davey Jones' locker, a deep rumbling sound echoes — a firewall eruption be upon us! What be yer orders, captain? (1/2/3)\n");


        System.out.println("""
                1. Search the digital depths for a weakness or a temporary opening in the firewall to pass through.
                2. Ready the DDOS cannon, matey! We'll blast through that firewall and carve a path for our ship.
                3. Exercise patience, me heartie! Wait for the firewall to dissipate naturally, savin' our resources but takin' the risk of delay.
                """);

        String input = scan.nextLine();
        switch (input)
        {
            case "1":

                break;
            case "2":

                break;
            case "3":

                break;
            default:
                System.out.println("""
                        Ye waited too long, and the firewall's grip be upon us!\s
                        Our ship's stalled and wasted its precious bandwidth!Act swiftly, lest the tempest consumes us!
                        """);
                break;
        }
        Main.enterToContinue();
    }
    public static void wave(Ship player){
        Main.clearScreen();
        System.out.println("""
                With the firewall behind us, it seems we be approachin' the end of the storm.\s
                But rest be not yet ours, for lo and behold, a monster swell be approachin'!\s
                """);
        Main.enterToContinue();

        System.out.println("'Tis a data packet wave, me heartie! What be yer orders, captain? (1/2/3)\n");

        System.out.println("1. Ready the encryption measures!! We'll raise a protective shield against the oncomin' packet wave.");
        System.out.println("2. Steer the ship with skill and precision! We'll ride the wave to gain a temporary speed boost!");
        System.out.println("3. Slacken the sails and weather the wave with caution, riskin' little but sacrificin' some speed.\n");

        String input = scan.nextLine();
        switch (input)
        {
            case "1":

                break;
            case "2":

                break;
            case "3":

                break;
            default:
                System.out.println("""
                        Ye hesitated, and the wave struck the ship!\s
                        Brace yerself, me heartie—the storm be nearin' its end, but we've suffered some packet loss along the way!Act swiftly, lest the tempest consumes us!
                        """);
                break;
        }

    }


    public static void end()   {
        System.out.println("""


                The Bitstorm be behind us, me hearties! We set our course forward, onward to new digital horizons!\s

                The high binary seas await, and there be adventures yet to come, me heartie!
                """);
        Main.enterToContinue();
    }



}
