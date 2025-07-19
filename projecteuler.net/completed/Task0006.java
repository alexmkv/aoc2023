package am.makeev;

public class Task0006
{
    public static long sum_squares_diff(long n)
    {
        long res = 0;
        for (long i = 1; i < n; ++i) {
            for (long j = i + 1; j <= n; ++j) {
                res += 2*i*j;
            }
        }
        return res;
    }
    public static void main( String[] args )
    {
        System.out.println(sum_squares_diff(10));
        System.out.println(sum_squares_diff(100));
    }
}