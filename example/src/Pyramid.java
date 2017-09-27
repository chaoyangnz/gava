public class Pyramid {
    public static void main(String args[]) {
        System.err.println("Pyramid pattern of star in Java : ");
        drawPyramidPattern();
        System.out.println("Pyramid of numbers in Java : ");
        drawPyramidOfNumbers();

//        System.out.println(Pyramid.class.getClassLoader());
    }

    /**
     * This method draws a pyramid pattern using asterisk character. You can * replace the asterisk with any other character to draw a pyramid of that.
     */
    public static void drawPyramidPattern() {
        for (int i = 0; i < 5; i++) {
            for (int j = 0; j < 5 - i; j++) {
                System.out.print(" ");
            }
            for (int k = 0; k <= i; k++) {
                System.out.print("* ");
            }
            System.out.println();
        }
    }

    /**
     * This method draws a pyramid of numbers.
     */
    public static void drawPyramidOfNumbers() {
        for (int i = 0; i < 5; i++) {
            for (int j = 0; j < 5 - i; j++) {
                System.out.print(" ");
            }
            for (int k = 0; k <= i; k++) {
                System.out.print(k + " ");
            }
            System.out.println();
        }
    }
}
