import java.io.IOException;

public class WaitInterrupt {



    public static void main(String[] args) throws InterruptedException, IOException {
        Thread t = new Thread(new Runnable() {
            Object lock = new Object();
            @Override
            public void run() {
                try {
                    synchronized (lock) {
                        lock.wait();
                    }
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        });

        t.start();
        Thread.sleep(3000);

        t.interrupt();

//        System.in.read();
    }


}

