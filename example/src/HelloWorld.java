public class HelloWorld {
    public static void main(String[] args) {
        if(args.length < 1) throw new IllegalArgumentException("At least one argument is required");
        System.out.println("Hello world, " + args[0]);
    }
}
