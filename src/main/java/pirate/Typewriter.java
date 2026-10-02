package pirate;

public final class Typewriter {

    private static final int DEFAULT_DELAY_MS = 20;

    private Typewriter() {}

    public static void print(String text) {
        print(text, DEFAULT_DELAY_MS);
    }

    public static void print(String text, int delayMs) {
        for (char c : text.toCharArray()) {
            System.out.print(c);
            System.out.flush();
            try {
                Thread.sleep(delayMs);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                return;
            }
        }
    }

    public static void println(String text) {
        println(text, DEFAULT_DELAY_MS);
    }

    public static void println(String text, int delayMs) {
        print(text, delayMs);
        System.out.println();
    }
}
