class A {}
class B extends A {}
class C extends B {}

public class Test {
    public static void main(String[] args) {
        System.out.println(new A() instanceof B);
        System.out.println(new B() instanceof A);
        System.out.println(new C() instanceof A);
        System.out.println(new B() instanceof B);
    }
}