import java.util.Scanner;
import java.util.Random;

public class Encounter {

    private static final Random roll = new Random();
    private static final Scanner scan = new Scanner(System.in);
    public static boolean textbook = false;
    public static boolean library = false;

    public static void email()//email
    {
        Main.clearScreen();
        System.out.println("First, let's cast our eyes upon our electronic treasure chest and see if we've plundered any digital messages.");
        System.out.println("\nAhoy, it seems there be a fresh scroll in the inbox!");
        System.out.println("\nEnter yer secret passphrase to unfurl the message, matey:");
        scan.nextLine();
        System.out.println("[Hey what's up, the torrent file for that new game just got released this morning." +
                "\nps if you get the chance could you send me the textbook for the class? ty.]");

        System.out.println("\nHmm, it be lookin' like we ought to make a course for The Pirate Bay." +
                "\nBut mayhaps we should pay a visit to the library if fortune smiles upon us.\n");
        libOrSea();
    }
    public static void libOrSea()
    {
        System.out.println("1. Go to the Library (1)\n" +
                "2. Sail straight to The Pirate Bay (2)");
        String ans = scan.nextLine();
        switch (ans)
        {
            case "1":
                library = true;
                System.out.println("\nAye, matey! To the Library Genesis we shall sail first! \n" +
                        "Knowledge be our treasure today!\n");
                Main.enterToContinue();
                library();
                break;
            case "2":
                library = false;
                System.out.println("There be no time to waste! Arrr, hoist the anchor! \n" +
                        "We'll be setting a course straight for The Pirate Bay, adventure awaits!\n");
                Main.enterToContinue();
                break;
            default:
                System.out.println("Yarrr! Ye didn't quite get the hang of it. \n" +
                        "Try inputting '1' to visit the Library or '2' to head to The Pirate Bay.\n");
                libOrSea();
        }

    }

    public static void library() {
        Main.clearScreen();
        System.out.println("Arrr! Ye find yerself at the legendary Library Genesis, \n" +
                "a digital treasure trove, hidden deep within the vast digital ocean.");
        System.out.println("\nStandin' at the entrance, a salty dog named Keshav ambles over, boastin' a friendly grin. \n" +
                "He introduces himself as the guardian of this trove and extends a digital textbook to ye.");
        libraryResponse();
    }
    public static void libraryResponse(){
        System.out.println("\nBe ye claimin' this tome, matey?(y/n)");
        String response = scan.nextLine();
        switch(response){
            case "y":
                textbook = true;
                System.out.println("\nWith the enthusiasm of a scallywag findin' buried treasure, ye snatch the book. \n" +
                        "Ye tuck it away for future plunder, knowin' full well it'll be dispatched to one of yer shipmates. \n" +
                        "With booty in hand, ye press on with yer grand adventure.");
                Main.enterToContinue();
                break;
            case "n":
                textbook = false;
                System.out.println("\nWith a polite but firm refusal, ye steer clear of the offer, \n" +
                        "continuin' yer escapade, a hearty 'no thanks' ringin' from yer lips.");
                Main.enterToContinue();
                break;
            default:
                System.out.println("Arrr! Ye seem to be confusin' yer compass. \n" +
                        "Try inputtin' 'y' if ye be takin' the treasure or 'n' if ye be passin' it up, matey. \n" +
                        "No other route be navigatin' this choice.");
                libraryResponse();
        }

    }



    public static void pirateBayIntro()
    {
        Main.clearScreen();
        System.out.println("""
                Af'er a harrowing journey through the turbulent seas o' the digital realm, yer vessel finally drops anchor at The Pirate Bay.\s
                This fabled sanctuary fer buccaneers, nestled deep within the binary tides, teems with activity.

                As ye make landfall, the sights an' sounds o' fellow pirates from all corners o' the digital world surround ye.\s
                Flags adorned with the symbols o' diverse factions dance in the cyber breezes, while the buzz o' anticipation fills the air.

                Settin' foot upon The Pirate Bay's digital shores, ye can't help but ponder the choices awaitin' ye.\s
                The pursuit o' that new game torrent might've led ye here, but the possibilities stretch as far as the digital horizon.

                Welcome to The Pirate Bay, where legends of the high binary seas come to life!
                """);
        Main.enterToContinue();

    }

    public static void pirateBay()
    {
        Main.clearScreen();
        if(library)
        {
            if(textbook)
            {
                System.out.println("Arr, ye've secured the coveted torrent ye sought, and ye've made a true pirate's promise to share the class tome with yer matey.\n");

                System.out.println("A fine haul and a loyal heart – the high binary seas have treated ye well, me heartie!\n");
                Main.enterToContinue();
            }
            else
            {
                System.out.println("Ye've got yer hands on the prized torrent, but ye missed the chance to honor the request of a fellow pirate.\n");

                System.out.println("Fear not, there be more chances ahead, and ye'll keep that promise in the next digital voyage.\n");
                Main.enterToContinue();
            }
        }
        else
        {
            System.out.println("Though the textbook slipped through yer grasp this time, ye've captured the digital treasure ye set out for.\n");

            System.out.println("With determination in yer heart, ye'll continue the search and ensure yer matey gets their hands on the class tome soon.\n");
            Main.enterToContinue();
        }

        System.out.println("\nWith yer ultimate goal achieved, the digital horizon stretches out before ye like an endless sea of opportunities.\n");

        System.out.println("The high binary seas be full of uncharted territory and bountiful treasures, and as ye stand on the brink of new adventures,\n");

        System.out.println("ye can't help but wonder where yer digital piracy journey will take ye next. Set sail, me heartie, for the horizon beckons!\n");
        Main.enterToContinue();

        System.out.println("What be your next plan? (1/2/3/4/5)\n");

        System.out.println("1. Explore the Pirate Bay: \n" +
                "Be on the lookout for other treasures, more torrents to plunder...");

        System.out.println("2. Engage in Digital Trade: \n" +
                "Snag the new game torrent and dedicate some time to seeding it for yer fellow buccaneers...");

        System.out.println("3. Set Sail for New Adventures: \n" +
                "Whispers tell of bountiful treasures awaitin' in the uncharted binary seas...");

        System.out.println("4. Venture to Treacherous Waters: \n" +
                "Upgradin' yer vessel with a Tor hull, sail the treacherous waters of the darkweb depths. \n" +
                "The risk be far greater than anything you've encountered, but the plunder be worth it...");

        System.out.println("5. Go Legit: \n" +
                "Ye've had yer fill of piracy and be considerin' the path of buyin' software. It may cost more, but the waters be calmer...");
        askEnding();
        Ending.endingEnd();
    }
    public static void askEnding()
    {
        String input = scan.nextLine();
        switch(input)
        {
            case "1":
                Ending.ending1();
                break;
            case "2":
                Ending.ending2();
                break;
            case "3":
                Ending.ending3();
                break;
            case "4":
                Ending.ending4();
                break;
            case "5":
                Ending.ending5();
                break;
            default:
                System.out.println("(Ye must enter 1, 2, 3, 4, or 5 to make a choice, matey!)");
                askEnding();
                break;

        }
    }




}
