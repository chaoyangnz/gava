package ds_util;

public class BinaryUtil {
    /**
     * convert a long to binary representation and pad to the specified length with leading zeros
     *
     * @param l
     * @param paddingTo
     * @return
     */
    public static String toBinary(long l, int paddingTo) {
        StringBuilder sb = new StringBuilder(Long.SIZE);
        int nonZeroBits = Long.SIZE - Long.numberOfLeadingZeros(l);
        if(paddingTo < nonZeroBits) {
            throw new IllegalArgumentException(String.format("%d exceed padding length %d", l, paddingTo));
        }
        int needPadding = paddingTo - nonZeroBits;
        for(int i = 0; i < needPadding; i++) {
            sb.append('0');
        }
        if(nonZeroBits > 0) sb.append(Long.toBinaryString(l));
        return sb.toString();
    }

    public static String toBinary(long l) {
        return toBinary(l, Long.SIZE);
    }

    private static String prettify(String str) {
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < str.length(); ++i) {
            if(i != 0 && i % Byte.SIZE == 0) {
                sb.append(" ");
            }
            sb.append(str.charAt(i));
        }
        return sb.toString();
    }

    public static String toPrettyBinary(long l) {
        return prettify(toBinary(l));
    }

    /**
     * convert an integer to binary representation and pad to the specified length with leading zeros
     *
     * @param i
     * @param paddingTo
     * @return
     */
    public static String toBinary(int i, int paddingTo) {
        StringBuilder sb = new StringBuilder(Integer.SIZE);
        int nonZeroBits = Integer.SIZE - Integer.numberOfLeadingZeros(i);
        if(paddingTo < nonZeroBits) {
            throw new IllegalArgumentException(String.format("%d exceed padding length %d", i, paddingTo));
        }
        int needPadding = paddingTo - nonZeroBits;
        for(int j = 0; j < needPadding; j++) {
            sb.append('0');
        }
        if(nonZeroBits > 0) sb.append(Integer.toBinaryString(i));
        return sb.toString();
    }

    public static String toBinary(int i) {
        return toBinary(i, Integer.SIZE);
    }

    public static String toPrettyBinary(int i) {
        return prettify(toBinary(i));
    }

    /**
     * Set the specified bit
     *
     * @param i
     * @param position
     * @return
     */
    public static int setBit(int i, int position) {
        return i | (1 << position);
    }

    /**
     * Clear the specified bit<br>
     *
     * -1: 11111111 11111111 111111111 111111111
     *
     * @param i
     * @param position
     * @return
     */
    public static int clearBit(int i, int position) {
        return i & (-1 ^ (1 << position));
    }

    public static void main(String[] args) {
        int a = 0b1101;
        System.out.println(toPrettyBinary(a));

        int b = clearBit(a, 2);
        System.out.println(toPrettyBinary(b));

        int c = setBit(a, 1);
        System.out.println(toPrettyBinary(c));
    }
}
