package am.makeev;

import java.util.ArrayList;
import java.util.List;

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */
public class Task0005
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
    public static long smallest_multiply(int num)
    {
        List<Integer> divisors = new ArrayList<>();
        for (int i = 2; i <= num; ++i) {
            int d = i;
            for (Integer d0: divisors) {
                if (d % d0 == 0) d = d/d0;
            }
            if (d != 1) divisors.add(d);
        }
        return divisors.stream().reduce(1, (a, b) -> a*b);
    }
    public static void main( String[] args )
    {
        System.out.println(smallest_multiply(10));
        System.out.println(smallest_multiply(20));
    }
}