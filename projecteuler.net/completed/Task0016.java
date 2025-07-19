package am.makeev;

import java.math.BigInteger;
import java.util.ArrayDeque;
import java.util.Deque;

public class Task0016
{
    private static long power_digit_sum(int n)
    {
        BigInteger bi = new BigInteger("2");
        bi = bi.pow(n);
        long sum = 0;
        String s = bi.toString();
        for (char ch : s.toCharArray()) {
            sum += (int)(ch - '0');
        }
        //return bi.toString().length();
        return sum;

    }

    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(power_digit_sum(15));
        System.out.println(power_digit_sum(1000));
        //System.out.println(power_digit_sum(20));
        //System.out.println(power_digit_sum(20+1));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}