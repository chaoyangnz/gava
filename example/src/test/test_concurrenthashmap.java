package test;

import java.lang.reflect.Field;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class test_concurrenthashmap {

    public static void main(String[] args) throws NoSuchFieldException, IllegalAccessException {
        Map map = new ConcurrentHashMap();
        Field field = ConcurrentHashMap.class.getDeclaredField("U");
        field.setAccessible(true);


        Field f = sun.misc.Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        sun.misc.Unsafe unsafe = (sun.misc.Unsafe) f.get(null);
        System.out.println(unsafe.ADDRESS_SIZE);
    }
}
