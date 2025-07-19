package am.makeev;


import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.text.Collator;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.TreeSet;

/**
 * A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
 * For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
 *
 * A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
 *
 * As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of
 * two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be
 * written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis
 * even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers
 * is less than this limit.
 *
 * Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
 */
public class Task0023
{

    private static boolean isAbundant(int v) {
        int dvSum = 0;
        int max = (v/2)+1;
        for (int i = 1; i < max; ++i) {
            if (v % i == 0)
            {
                dvSum += i;
                if (dvSum > v) return true;
            }
        }
        return false;
    }

    private static int sum_of_abundant_numbers()
    {
        final int MAX_AB_INT = 28123;

        final TreeSet<Integer> ab = new TreeSet<>();

        int sum = 1;
        for (int i = 2; i <= MAX_AB_INT; ++i) {
            boolean can = false;
            int half = i/2;
            for (Integer i0 : ab) {
                if (i0 > half) break;
                int s = i - i0;
                if (s > 0 && ab.contains(s)) {
                    can = true;
                    break;
                }
            }
            if (!can) {
                sum += i;
            }
            if (isAbundant(i))
            {
                ab.add(i);
            }
        }
        return sum;
    }

    public static void main( String[] args ) throws IOException
    {
        long startTime = System.currentTimeMillis();
        System.out.println(sum_of_abundant_numbers());
        System.out.println(System.currentTimeMillis()-startTime);
    }

}