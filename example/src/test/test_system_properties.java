package test;

import java.util.Properties;

public class test_system_properties {

    public static void main(String[] args) {
        Properties properties = System.getProperties();

        //System.out.println(properties.getProperty("line.separator").equals("\n"));
        properties.list(System.out);
//        while (.propertyNames().hasMoreElements()){
//            String name = (String)properties.propertyNames().nextElement();
//            System.out.println(String.format("%s: %s", name, properties.get(name)));
//        }
    }
}
