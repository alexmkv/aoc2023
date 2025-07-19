package am.makeev;


import java.math.BigInteger;

public class Task0020
{

    private static long factorial_digit_sum(int n)
    {
        BigInteger res = new BigInteger("1");
        for (int i = 2; i <= n; ++i) {
            res = res.multiply(BigInteger.valueOf(i));
        }
        return res.toString().chars().mapToLong(v -> v - '0').sum();
    }


    public static void main( String[] args )
    {
        long startTime = System.currentTimeMillis();
        System.out.println(factorial_digit_sum(10));
        System.out.println(factorial_digit_sum(100));
        System.out.println(System.currentTimeMillis()-startTime);
    }

}