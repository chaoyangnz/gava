package test;

public class test_athrow {

    public static void main(String[] args) {
        String a = "2";
        try {
            go();
            System.out.println("Do something after going");
        } catch (Exception1 ex) { // only catch Exception1
            System.out.println(a);
            System.out.println(ex.getMessage());
        }

    }

    public static void go() {
        boolean ok = false;
        System.out.println("prepare to go");
        onRoad();
        ok = true;
    }

    public static void onRoad() {
        rest();
        throw new Exception1("this is exception 1");
    }

    public static int rest() {
        int a = 0;
        for(; a < 10; a++) {
            if(a == 7) {
                throw new Exception2("this is exception 2");
            }
        }
        return 1;
    }
}

class Exception1 extends RuntimeException {

    public Exception1(String xxxx) {
        super(xxxx);
    }
}

class Exception2 extends RuntimeException {

    public Exception2(String yyy) {
        super(yyy);
    }
}
