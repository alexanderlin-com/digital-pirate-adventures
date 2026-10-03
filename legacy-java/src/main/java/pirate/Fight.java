package pirate;

import java.util.Random;
import java.util.Scanner;

public class Fight {
    private static final Random roll = new Random();
    private static final Scanner scan = Main.scan;

    public static boolean initiative(Ship player, Enemy enemy)
    {
        int playerInitiative = player.speed + Main.d20();
        int enemyInitiative = enemy.speed + Main.d20();
        return playerInitiative>=enemyInitiative;
    }


    public static void intro()
    {
        Main.clearScreen();
        System.out.println("With the Bitstorm finally behind us, the air clears, and we find ourselves bathed in the calm after the digital tempest.\n"
                +"The ship sails steadily as we breathe a sigh of relief, believin' the worst to be behind us.\n"
                +"Little did we know, me hearties, that the high binary seas still held a twist of fate in store!\n");

        Main.enterToContinue();

        System.out.println("As we sail these digital seas, a shadow looms - an ominous privateer, relentless in their pursuit of buccaneers like us.\n"
                +"Defeat could mean a life behind bars for our digital transgressions - it be a fate worse than death!\n"
                +"Stand ready, me hearties, for the clash of digital swords and bytes draws nigh!");

        Main.enterToContinue();
    }

    public static void run(Ship player, Enemy enemy)
    {
        System.out.println("\nThe privateer closes in! Fight back to weaken their pursuit, or look for yer chance to flee!");

        while (true)
        {
            boolean playerFirst = initiative(player, enemy);

            if (playerFirst)
            {
                if (playerTurn(player, enemy))
                {
                    break;
                }
                enemyTurn(player, enemy);
            }
            else
            {
                enemyTurn(player, enemy);
                player.healthCheck();
                if (playerTurn(player, enemy))
                {
                    break;
                }
            }

            player.healthCheck();
            Main.enterToContinue();
        }

        System.out.println("\nThe privateer fades into the digital fog behind ye. Ye've made yer escape!");
        Main.enterToContinue();
    }

    private static boolean playerTurn(Ship player, Enemy enemy)
    {
        System.out.println("\nWhat be yer move, captain? (1/" + (player.swivel ? "2/" : "") + (player.crewType != 0 ? "3/" : "") + "9)\n");
        System.out.println("1. Broadside Cannon - let loose with the big guns.");
        if (player.swivel)
        {
            System.out.println("2. Swivel Gun - a quicker, more precise shot.");
        }
        if (player.crewType != 0)
        {
            System.out.println("3. Crew Ability - call on the crew for somethin' special.");
        }
        System.out.println("9. Flee - try to lose 'em and slip away.");

        String input = scan.nextLine();
        switch (input)
        {
            case "1":
                playerBroadside(player, enemy);
                return false;
            case "2":
                if (player.swivel)
                {
                    playerSwivel(player, enemy);
                    return false;
                }
                System.out.println("Arr, this ship's got no swivel gun, matey.");
                return playerTurn(player, enemy);
            case "3":
                if (player.crewType != 0)
                {
                    playerCrew(player, enemy);
                    return false;
                }
                System.out.println("Arr, this ship's got no crew ability set, matey.");
                return playerTurn(player, enemy);
            case "9":
                return attemptFlee(player, enemy);
            default:
                System.out.println("Arr, I didn't quite catch that, matey.");
                return playerTurn(player, enemy);
        }
    }

    private static boolean resolveHit(int attackerAccuracy, int weaponMod, int defenderDodge)
    {
        return Main.d20() + attackerAccuracy + weaponMod > 10 + defenderDodge;
    }

    private static void playerBroadside(Ship player, Enemy enemy)
    {
        int baseDamage;
        int accuracyMod;
        switch (player.broadsideType)
        {
            case 1:
                baseDamage = roll.nextInt(16) + 15;
                accuracyMod = -3;
                break;
            case 2:
                baseDamage = roll.nextInt(11) + 10;
                accuracyMod = 0;
                break;
            case 3:
                baseDamage = roll.nextInt(11) + 5;
                accuracyMod = 5;
                break;
            default:
                baseDamage = roll.nextInt(11) + 10;
                accuracyMod = 0;
                break;
        }

        player.useBandwidth(10);
        player.bandwidthCheck();

        if (resolveHit(player.accuracy, accuracyMod, enemy.dodge))
        {
            int damage = (int) (baseDamage * player.damage);
            enemy.takeDamage(damage);
            System.out.println("Yer broadside cannons roar! Ye deal " + damage + " damage to the privateer!");
        }
        else
        {
            System.out.println("Yer broadside cannons miss their mark!");
        }
    }

    private static void playerSwivel(Ship player, Enemy enemy)
    {
        player.useBandwidth(5);
        player.bandwidthCheck();

        switch (player.swivelType)
        {
            case 1:
                if (resolveHit(player.accuracy, 0, enemy.dodge))
                {
                    int damage = (int) ((roll.nextInt(10) + 1) * player.damage);
                    enemy.takeDamage(damage);
                    enemy.accuracy -= 2;
                    System.out.println("Yer zip bomb swivel hits for " + damage + " damage and scrambles their targeting systems!");
                }
                else
                {
                    System.out.println("Yer zip bomb swivel shot goes wide!");
                }
                break;
            case 2:
                if (resolveHit(player.accuracy, 0, enemy.dodge))
                {
                    int damage = (int) ((roll.nextInt(10) + 1) * player.damage);
                    enemy.takeDamage(damage);
                    player.accuracy += 2;
                    System.out.println("Yer IP-trace swivel hits for " + damage + " damage and sharpens yer own aim!");
                }
                else
                {
                    System.out.println("Yer IP-trace swivel shot goes wide!");
                }
                break;
            case 3:
                int damage = (int) (roll.nextInt(26) * player.damage);
                enemy.takeDamage(damage);
                System.out.println("Yer logic bomb swivel unleashes pure chaos, dealin' " + damage + " damage!");
                break;
            default:
                System.out.println("Ye haven't settled on a swivel gun load, matey!");
                break;
        }
    }

    private static void playerCrew(Ship player, Enemy enemy)
    {
        player.useBandwidth(15);
        player.bandwidthCheck();

        switch (player.crewType)
        {
            case 1:
                if (resolveHit(player.accuracy, 0, enemy.dodge))
                {
                    int damage = (int) ((roll.nextInt(11) + 8) * player.damage);
                    enemy.takeDamage(damage);
                    enemy.speed -= 2;
                    System.out.println("Yer crew's musket volley hits for " + damage + " damage and slows their pursuit!");
                }
                else
                {
                    System.out.println("Yer crew's musket volley misses!");
                }
                break;
            case 2:
                boolean hit = resolveHit(player.accuracy, 0, enemy.dodge);
                player.armor *= 0.5;
                if (hit)
                {
                    int damage = (int) ((roll.nextInt(11) + 10) * player.damage * 2);
                    enemy.takeDamage(damage);
                    System.out.println("Yer crew overclocks the powder kegs! The blast hits for " + damage + " damage, but yer defenses take a hit from the strain!");
                }
                else
                {
                    System.out.println("Yer crew overclocks the powder kegs, but the shot goes wide! Yer defenses still took a hit from the strain!");
                }
                break;
            default:
                System.out.println("Ye haven't settled on a crew ability, matey!");
                break;
        }
    }

    private static boolean attemptFlee(Ship player, Enemy enemy)
    {
        if (enemy.health <= 0)
        {
            System.out.println("\nTheir systems be fried from yer assault! Ye slip away unopposed, matey!");
            return true;
        }
        if (initiative(player, enemy))
        {
            System.out.println("\nWith a burst of speed, ye leave the privateer in yer wake! Ye've made yer escape!");
            return true;
        }
        System.out.println("\nYe try to break away, but the privateer keeps pace! No luck this time, matey!");
        return false;
    }

    private static void enemyTurn(Ship player, Enemy enemy)
    {
        int attack = roll.nextInt(6) + 1;
        int baseDamage;
        String name;
        switch (attack)
        {
            case 1:
                name = "IP Tracker Swivel Gun";
                baseDamage = roll.nextInt(11) + 5;
                break;
            case 2:
                name = "Firewall Broadside";
                baseDamage = roll.nextInt(16) + 10;
                break;
            case 3:
                name = "Data Snare Chainshot";
                baseDamage = roll.nextInt(11) + 5;
                break;
            case 4:
                name = "Proxy Buster Musket Volley";
                baseDamage = roll.nextInt(11) + 5;
                break;
            case 5:
                name = "Encryption Jammer Grapeshot";
                baseDamage = roll.nextInt(11) + 8;
                break;
            default:
                name = "Codebreaker Broadside";
                baseDamage = roll.nextInt(16) + 10;
                break;
        }

        System.out.println("\nThe privateer unleashes a " + name + "!");

        if (!resolveHit(enemy.accuracy, 0, player.dodge))
        {
            System.out.println("Ye manage to avoid the attack!");
            return;
        }

        int damage = (int) (baseDamage * (1 - player.armor));
        player.takeDamage(damage);
        System.out.println("It strikes true, dealin' " + damage + " damage!");

        switch (attack)
        {
            case 1:
            case 2:
                player.nerfSpeed(2);
                System.out.println("Yer ship's speed be hamperin' from the hit!");
                break;
            case 3:
                player.nerfDodge(2);
                System.out.println("Yer ship's maneuverability be hamperin' from the hit!");
                break;
            case 4:
            case 5:
                player.nerfArmor(0.05);
                System.out.println("Yer ship's defenses be weakenin' from the hit!");
                break;
            case 6:
                player.nerfAccuracy(2);
                System.out.println("Yer ship's targeting systems be jammed from the hit!");
                break;
        }
    }

}
