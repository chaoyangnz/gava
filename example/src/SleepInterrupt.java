public class SleepInterrupt {

    public static void main(String[] args) throws InterruptedException {
        Thread t= new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    Thread.sleep(10000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        });

        t.start();
        Thread.sleep(2000);
        t.interrupt();
    }
}
