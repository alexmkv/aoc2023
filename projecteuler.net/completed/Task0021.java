package am.makeev;


import java.math.BigInteger;
import java.util.HashMap;
import java.util.Map;

public class Task0021
{

    private static long amicable_numbers(int n)
    {
        Map<Integer, Integer> mp = new HashMap<>(n);
        int[] m = new int[n];
        for (int i = 1; i < n; ++i) {
            int s = 0;
            for (int j = 1; j < i; ++j) {
                if (i % j ==0) s+=j;
            }
            m[i] = s;
        }
        int r = 0;
        for (int i = 1; i < n; ++i) {
            if (m[i] > i && m[i] < n && m[m[i]] == i) {
                r += i + m[i];
            }
        }
        return r;
    }


    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        //System.out.println(amicable_numbers(10));
        // err: 32026982
        System.out.println(amicable_numbers(10000));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}