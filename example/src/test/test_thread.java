package test;

public class test_thread {

    public static void main(String[] args) {
        Thread thread = Thread.currentThread();
        System.out.println("group:" + thread.getThreadGroup().getName());
        System.out.println("id:" + thread.getId());

        new Thread(new Runnable() {
            @Override
            public void run() {
                Thread thread = Thread.currentThread();
                System.out.println("thread group:" + thread.getThreadGroup().getName());
                System.out.println("thread id:" + thread.getId());
                System.out.println("thread name:" + thread.getName());
                System.out.println("thread priority:" + thread.getPriority());
            }
        }).start();

        System.out.println("name:" + thread.getName());
        System.out.println("priority:" + thread.getPriority());


    }
}
