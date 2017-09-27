public class EnclosingClasses {

    int a;

    public static void main(String[] args) {
        EnclosingClasses enclosingClasses = new EnclosingClasses();
        enclosingClasses.a = 10;

        TheInnerClass theInnerClass = enclosingClasses.new TheInnerClass();
        theInnerClass.b = 20;
        theInnerClass.accessOuter();

        TheNestedClass theNestedClass = new TheNestedClass();
        theNestedClass.c = 30;
        theNestedClass.print();
    }

    public class TheInnerClass {
        int b;

        void accessOuter() {
            System.out.println("Enclosing a: " + EnclosingClasses.this.a);
            System.out.println("TheInnerClass b: " + b);
        }
    }

    public static class TheNestedClass {
        int c;

        void print() {
            System.out.println("TheNestedClass c: " + c);
        }
    }
}
