package jago;

public class Jago {

    public  static void initSystem() {
        System.setProperty("java.lang.Integer.IntegerCache.high", "127");
    }

    private static native void print(String x);

    public static native void print(double x);

    public static void print(Object o) {
        print(String.valueOf(o));
    }

    public static void printf(String format, Object...args) {
        print(String.format(format, args));
    }

    public static void println(Object o) {
        print(String.valueOf(o) + "\n");
    }

    public static void println() {
        print("\n");
    }
}

