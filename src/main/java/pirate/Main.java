import java.util.Random;
import java.util.Scanner;


public class Main {
        private static final Random roll = new Random();
        public static final Scanner scan = new Scanner(System.in);

        public static void main(String[] args) {

            Enemy enemy = new Enemy();

            intro();
            introPlayerResponse();
            setup();
            Ship player = askShipType();
            player.customize();
            embark();
            Encounter.email();
            Storm.intro();
            Storm.lightning(player);
            Storm.whirlwind(player);
            Storm.eruption(player);
            Storm.wave(player);
            Storm.end();
            Fight.intro();


            Encounter.pirateBayIntro();
            Encounter.pirateBay();

        }

        public static void intro() {
            System.out.println("""
                    Ahoy there fellow digital pirate!
                    Welcome aboard yer trust ship and prepare to sail the internet seas!
                    I be yer trusty virtual quartermaster.
                    Ready to embark on this digital adventure? (y/n)""");
        }
        public static void introPlayerResponse() {
            Scanner scan = new Scanner(System.in);
            String response = scan.nextLine();
            switch (response) {
                case "y":
                    System.out.println("Aye, ye be ready for adventure!\n");
                    break;
                case "n":
                    System.out.println("Ah, a landlubber, are ye? Farewell for now!");
                    System.exit(0);
                default:
                    System.out.println("Arr, I didn't quite catch that. Try 'y' for yes or 'n' for no, matey!");
                    introPlayerResponse();
            }
        }
        public static void setup()
        {
            askPlayerName();
            askShipName();
            System.out.println("No finer names for the high seas, indeed!");
        }

        public static void askPlayerName()
        {
            System.out.println("Ahoy! Before we set sail, what be yer name, matey?");
            String input = scan.nextLine();
            if (input.isEmpty()) {
                System.out.println("Looks like ye didn't input anything or just spaces, matey. What be yer name?");
                askPlayerName();
            }
            else    {
                Ship.playerName = input;
            }
        }
        public static void askShipName()
        {
            System.out.println("And what be the name of our ship?");
            String input = scan.nextLine();
            if (input.isEmpty()) {
                System.out.println("Looks like ye didn't input anything or just spaces, matey. What ye be callin' yer ship?");
                askPlayerName();
            }
            else    {
                Ship.shipName = input;
            }
        }


        public static Ship askShipType()
        {
            System.out.println("""
                    What be the type o' yer ship, matey? (1/2/3)
                    1. Sloop - fastest speed, lowest firepower
                    2. Frigate - balanced speed n' firepower
                    3. Galleon - superior firepower, slowest speed""");
            String input = scan.nextLine();
            return switch (input) {
                case "1" -> new Sloop(200, 100, 10, 7, 0.1, 0, 0.75);
                case "2" -> new Frigate(350, 250, 7, 3, 0.20, 3, 100);
                case "3" -> new Galleon(500, 200, 4, 1, 0.30, 6, 1.25);
                default -> {
                    System.out.println("Arr, I didn't quite catch that. Input 1/2/3, to choose yer ship type, matey.");
                    yield askShipType();
                }
            };

        }
        public static void embark()
        {
            System.out.println("Take " + Ship.shipName + " ahead full, Captain " + Ship.playerName + "! \n" +
                    "The high binary seas await!");
            enterToContinue();
        }

  














        public static void clearScreen() {
            for (int i = 0; i < 50; ++i) System.out.println();
        }

        public static void enterToContinue() {
            System.out.println("\n(Enter anything to continue)");
            scan.nextLine();
        }



        public static int d20() {
            return roll.nextInt(20) + 1;
        }



}

