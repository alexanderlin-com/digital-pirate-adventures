package pirate;

import org.junit.jupiter.api.RepeatedTest;

import static org.junit.jupiter.api.Assertions.assertTrue;

class MainTest {

    @RepeatedTest(50)
    void d20StaysWithinDieRange() {
        int roll = Main.d20();
        assertTrue(roll >= 1 && roll <= 20);
    }
}
