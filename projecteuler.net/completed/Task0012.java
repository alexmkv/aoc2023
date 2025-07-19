package am.makeev;


public class Task0012
{
    public static long triangle_with_n_divisors(long n)
    {
        long num = 1;
        long i = 2;
        long max_div = 0;
        while (true) {
            long res = count_divisors(num);
            if (max_div < res) {
                System.out.format("%d: %d\n", num, res);
                max_div = res;
            }
            if (res > n) break;
            num += i;
            ++i;
        }
        return num;
    }
    public static long count_divisors(long v0)
    {
        long v = v0;
        long c = 1;
        for (long i = 2; i <= v; ++i) {
            long cv = 0;
            while (v % i == 0)
            {
                v = v / i;
                cv++;
            }
            if (cv > 0) c *= (cv+1);
        }
        return c;
    }
    public static long count_divisors_simple(long v)
    {
        long c = 0;
        for (long i = 1; i <= v; ++i) {
            if (v % i == 0) c++;
        }
        return c;
    }
    public static void main( String[] args )
    {
        System.out.println(triangle_with_n_divisors(5));
        System.out.println(triangle_with_n_divisors(500));
    }
}