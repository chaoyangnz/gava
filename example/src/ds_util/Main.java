package ds_util;

public class Main {
    public static void main(String[] args) {
        final String[] strs = {"hello"};
        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    Thread.sleep(5000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                int len = strs.length;
                System.out.println(strs[0]);
            }
        }).start();
        strs[0] = "world";
    }

    public void foo() {

    }
}
