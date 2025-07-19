package am.makeev;

public class Task0007
{
    public static long nth_prime(int n)
    {
        long[] primes = new long[n];
        long tst = 2;
        while (primes[n-1] == 0) {
            for (int i = 0; i < n; ++i) {
                if (primes[i] == 0)
                {
                    primes[i] = tst;
                    break;
                }
                if (tst % primes[i] == 0) break;
            }
            tst++;
        }
        return primes[n-1];
    }
    public static void main( String[] args )
    {
        System.out.println(nth_prime(6));
        System.out.println(nth_prime(10001));
    }
}