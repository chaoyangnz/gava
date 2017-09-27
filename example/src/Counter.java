public class Counter {

    Integer counter = 0;

    public static void main (String[] args) throws InterruptedException {
        Counter counter = new Counter();
        Task task1 = counter.new Task();
        Thread thread1 = new Thread(task1);

        Task task2 = counter.new Task();
        Thread thread2 = new Thread(task2);

        thread1.start();
        thread2.start();
    }

//    synchronized
    void performTask () {
        int temp = counter;
        counter++;
        System.out.println(Thread.currentThread()
                .getName() + " - before: "+temp+" after:" + counter);

        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    private class Task implements Runnable {

        @Override
        public void run () {
            for (int i = 0; i < 5; i++) {
                performTask();
            }
        }
    }
}

