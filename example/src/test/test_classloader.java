package test;

import java.util.Properties;

public class test_classloader {
    public static void main(String[] args) {
        Properties properties = System.getProperties();

        properties.list(System.out);

        ClassLoader cl = test_classloader.class.getClassLoader();
        cl.getParent();
    }
}
