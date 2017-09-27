package test;

public class test_shift {

    public static void main(String[] args) {
        int a = 89400;
        System.out.println(a << 8);
        System.out.println(a >> 9);
        System.out.println(a >>> 9);
        int b = -937883;
        System.out.println(b << 8);
        System.out.println(b >> 2);
        System.out.println(b >>> 2);

        long la = 894003433;
        System.out.println(la << 35);
        System.out.println(la >> 36);
        System.out.println(la >>> 36);
        long lb = -9378833;
        System.out.println(lb << 40);
        System.out.println(lb >> 42);
        System.out.println(lb >>> 42);

        System.out.println(-a);
        System.out.println(-b);
        System.out.println(-la);
        System.out.println(-lb);

        System.out.println(a % b);
        System.out.println(la % lb);

//        Jago.print(93.39 % -39.4);
        System.out.println(93.39 % -39.4);
    }
}
