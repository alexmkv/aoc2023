package am.makeev;


import java.util.BitSet;

public class Task0010
{
    public static long primes_sum_below(int n)
    {
        BitSet ar = new BitSet(n);
        long sum = 0;
        for (int i = 2; i < n; ++i) {
            if (!ar.get(i)) {
                sum += i;
                int j = i*2;
                while (j < n) {
                    ar.set(j);
                    j += i;
                }
            }
        }
        return sum;
    }
    public static void main( String[] args )
    {
        System.out.println(primes_sum_below(10)); //60
        System.out.println(primes_sum_below(2000000));
    }
}