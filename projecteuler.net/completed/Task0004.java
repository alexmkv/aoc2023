package am.makeev;

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */
public class Task0004
{
    public static boolean is_palindrome(long i)
    {
        String s = Long.toString(i);
        if (s.length() % 2 != 0) return false;
        for (int j = 0; j < s.length()/2; ++j) {
            if (s.charAt(j) != s.charAt(s.length() - j - 1)) return false;
        }
        return true;
    }
    public static long largest_palindrome()
    {
        long maxpoly = 0;
        for (int i = 999; i > 0 ; --i) {
            for (int j = 999; j >= 0; --j) {
                long v = i * j;
                if (v < maxpoly) break;
                if (is_palindrome(v))
                    if (v > maxpoly) maxpoly = v;
            }
        }
        return maxpoly;
    }
    public static void main( String[] args )
    {
        System.out.println(largest_palindrome());
    }
}