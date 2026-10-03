package pirate;

public class Enemy {
    public int health = 100;
    public int speed = 6;
    public int accuracy = 5;
    public int dodge = 4;

    public void takeDamage(int damage) {
        this.health -= damage;
        if (this.health < 0) {
            this.health = 0;
        }
    }
}
