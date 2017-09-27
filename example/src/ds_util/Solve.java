package ds_util;

import java.util.Arrays;

public class Solve {
    public int solve(int[] arr) {
        Arrays.sort(arr);
        int i = arr.length-1;
        int j = i - 1;
        int k = j - 1;

        while (i >=0 && j >= 0 && k >= 0) {
            if(arr[j] + arr[k] > arr[i]) {
                return arr[i] + arr[j] + arr[k];
            }
            i--;
            j = i - 1;
            k = j - 1;
        }

        return 0;
    }

    public static void main(String[] args) {
//        int[] arr = {2, 3, 4, 5, 10};
//        int[] arr = {4, 5, 10, 20};
        int[] arr = {4, 7, 8, 3, 12, 10};
        System.out.println(new Solve().solve(arr));
    }
}
