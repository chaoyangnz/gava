import java.util.HashMap;
import java.util.Map;

/**
 * Can be used in shorten url.
 * Store urls in DB and get an id, then encode this id to a [a-zA-Z0-9] string
 */
public class Base62 {
    Map<Character, Integer> char2int = new HashMap<>();
    Map<Integer, Character> int2char = new HashMap<>();

    public Base62() {
        int i = 0;
        for(char ch = 'a'; ch <= 'z'; ++ch) {
            char2int.put(ch, ++i);
            int2char.put(i, ch);
        }
        for(char ch = 'A'; ch <= 'Z'; ++ch) {
            char2int.put(ch, ++i);
            int2char.put(i, ch);
        }
        for(char ch = '0'; ch <= '9'; ++ch) {
            char2int.put(ch, ++i);
            int2char.put(i, ch);
        }
        char2int.put('?', ++i);
        int2char.put(i, '?');
        char2int.put('/', ++i);
        int2char.put(i, '/');
        char2int.put('*', ++i);
        int2char.put(i, '*');
        char2int.put('%', ++i);
        int2char.put(i, '%');
    }

    public static void main(String[] args) {
        Base62 base62 = new Base62();

        String encoded = base62.encode(1930943033930040L);

        System.out.printf("%s %n", encoded);
        System.out.printf("%d %n", base62.decode(encoded));

        long n = base62.decode("b?ot/8%bah?*muQFOmdT");


        System.out.println(Long.toOctalString(n));
    }

    public String encode(long id) {
        StringBuilder encoded = new StringBuilder();
        long n = id;
        while (n > 0) {
            encoded.append(int2char.get((int)(n % 62)));
            n /= 62;
        }

        return encoded.reverse().toString();
    }

    public long decode(String url) {
        long id = 0;
        for(char ch : url.toCharArray()) {
            id = id * 62 + char2int.get(ch);
        }
        return id;
    }
}